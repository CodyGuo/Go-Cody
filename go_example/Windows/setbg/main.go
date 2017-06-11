package main

import (
	"errors"
	"flag"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/golang/sys/windows/registry"
	"golang.org/x/image/bmp"
)

type WallpaperStyle uint

func (wps WallpaperStyle) String() string {
	return wallpaperStyles[wps]
}

const (
	Fill    WallpaperStyle = iota // 填充
	Fit                           // 适应
	Stretch                       // 拉伸
	Tile                          // 平铺
	Center                        // 居中
	Cross                         // 跨区

)

var wallpaperStyles = map[WallpaperStyle]string{
	0: "填充",
	1: "适应",
	2: "拉伸",
	3: "平铺",
	4: "居中",
	5: "跨区"}

var (
	bgFile       string
	bgStyle      int
	sFile        string
	waitTime     int
	activeScreen bool
	passwd       bool
)

var (
	regist registry.Key
)

func init() {
	flag.StringVar(&bgFile, "b", "", "set bg file path.")
	flag.IntVar(&bgStyle, "style", 2, "set desktop WallpaperStyle")

	flag.BoolVar(&activeScreen, "a", true, "set screen active.")
	flag.StringVar(&sFile, "s", "", "set screen save file path.")
	flag.IntVar(&waitTime, "t", 0, "set screen save wait time.")
	flag.BoolVar(&passwd, "p", true, "sets whether the screen saver requires the user to enter a password to display the Windows desktop. ")

}

func main() {
	flag.Parse()

	var err error
	regist, err = registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.ALL_ACCESS)
	checkErr(err)
	defer regist.Close()

	// 设置桌面背景
	if bgFile != "" {
		style := WallpaperStyle(bgStyle)
		if bgStyle < 0 || bgStyle > 5 {
			style = Stretch
		}
		setDesktopWallpaper(bgFile, style)
		log.Printf("设置桌面背景和位置 --> %s, %s\n", bgFile, style)
	}

	ok := getScreenSaver()
	log.Printf("获取屏幕保护开关 --> %t\n", ok)
	// 关闭屏幕保护
	if ok && !activeScreen {
		regist.DeleteValue("SCRNSAVE.EXE")
		log.Println("关闭屏幕保护")
		return
	}

	// 设置屏幕保护
	if sFile != "" && activeScreen {
		err = regist.SetStringValue("SCRNSAVE.EXE", sFile)
		checkErr(err)
		setScreenSaver(SPI_SETSCREENSAVEACTIVE, TRUE)
		log.Printf("设置屏幕保护 --> %s\n", sFile)
		ok = getScreenSaver()
	}

	if ok {
		// 设置屏幕保护时间
		if waitTime > 0 {
			setScreenSaver(SPI_SETSCREENSAVETIMEOUT, uint32(60*waitTime))
			log.Printf("设置屏幕保护等待时间 --> %d分钟\n", waitTime)
		}

		// 设置屏幕保护 在恢复时使用密码
		var (
			passwdSwitch string
			passwdBool   uint32
		)
		if passwd {
			passwdSwitch = "1"
			passwdBool = TRUE
		} else {
			passwdSwitch = "0"
			passwdBool = FALSE
		}
		// XP / server 2003
		setRegistString("ScreenSaverIsSecure", passwdSwitch)
		// vista or later
		if checkVersion() {
			setScreenSaver(SPI_SETSCREENSAVESECURE, passwdBool)
		}
		log.Printf("设置屏幕保护恢复时是否使用密码 --> %t\n", passwd)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// http://blog.csdn.net/kfysck/article/details/8153264
// Check that the OS is Vista or later (Vista is v6.0).
func checkVersion() bool {
	version := GetVersion()
	major := version & 0xFF
	if major < 6 {
		return false
	}
	return true
}

// jpg转换为bmp
func ConvertedWallpaper(bgfile string) string {
	file, err := os.Open(bgfile)
	checkErr(err)
	defer file.Close()

	img, err := jpeg.Decode(file) //解码
	checkErr(err)

	bmpPath := os.Getenv("USERPROFILE") + `\Local Settings\Application Data\Microsoft\Wallpaper1.bmp`
	bmpfile, err := os.Create(bmpPath)
	checkErr(err)
	defer bmpfile.Close()

	err = bmp.Encode(bmpfile, img)
	checkErr(err)
	return bmpPath
}

func setDesktopWallpaper(bgFile string, style WallpaperStyle) error {
	ext := filepath.Ext(bgFile)
	// vista 以下的系统需要转换jpg为bmp（xp、2003）
	if !checkVersion() && ext != ".bmp" {
		setRegistString("ConvertedWallpaper", bgFile)
		bgFile = ConvertedWallpaper(bgFile)
	}

	// 设置桌面背景
	setRegistString("Wallpaper", bgFile)

	/* 设置壁纸风格和展开方式
	在Control Panel\Desktop中的两个键值将被设置
	TileWallpaper
	 0: 图片不被平铺
	 1: 被平铺
	WallpaperStyle
	 0:  0表示图片居中，1表示平铺
	 2:  拉伸填充整个屏幕
	 6:  拉伸适应屏幕并保持高度比
	 10: 图片被调整大小裁剪适应屏幕保持纵横比
	 22: 跨区
	*/
	var bgTileWallpaper, bgWallpaperStyle string
	bgTileWallpaper = "0"
	switch style {
	case Fill: // (Windows 7 or later)
		bgWallpaperStyle = "10"
	case Fit: // (Windows 7 or later)
		bgWallpaperStyle = "6"
	case Stretch:
		bgWallpaperStyle = "2"
	case Tile:
		bgTileWallpaper = "1"
		bgWallpaperStyle = "0"
	case Center:
		bgWallpaperStyle = "0"
	case Cross: // win10 or later
		bgWallpaperStyle = "22"
	}

	setRegistString("WallpaperStyle", bgWallpaperStyle)
	setRegistString("TileWallpaper", bgTileWallpaper)

	ok := SystemParametersInfo(SPI_SETDESKWALLPAPER, FALSE, nil, SPIF_UPDATEINIFILE|SPIF_SENDWININICHANGE)
	if !ok {
		return errors.New("Desktop background Settings fail.")
	}
	return nil
}

func setRegistString(name, value string) {
	err := regist.SetStringValue(name, value)
	checkErr(err)
}

func setScreenSaver(uiAction, uiParam uint32) {
	ok := SystemParametersInfo(uiAction, uiParam, nil, SPIF_UPDATEINIFILE|SPIF_SENDWININICHANGE)
	if !ok {
		log.Fatal("Screen saver Settings fail.")
	}
}

func getScreenSaver() bool {
	_, _, err := regist.GetStringValue("SCRNSAVE.EXE")
	if err != nil {
		return false
	}
	return true
}
