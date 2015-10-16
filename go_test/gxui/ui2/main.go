package main

import (
    "io/ioutil"
    "log"
    "os"

    "github.com/google/gxui"
    "github.com/google/gxui/drivers/gl"
    "github.com/google/gxui/themes/dark"
)

var (
    FontPath = map[string]string{
        "宋体": "STZHONGS.TTF",
        "幼圆": "SIMYOU.TTF",
    }
)

func appMain(driver gxui.Driver) {
    theme := dark.CreateTheme(driver)

    window := theme.CreateWindow(800, 600, "Hi")
    window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray50))

    fontData, err := ioutil.ReadFile(sysFont("幼圆")) //this font comes from windows
    if err != nil {
        log.Fatalf("error reading font: %v", err)
    }
    font, err := driver.CreateFont(fontData, 50)
    if err != nil {
        panic(err)
    }
    label := theme.CreateLabel()
    label.SetFont(font)
    label.SetColor(gxui.Red50)
    label.SetText("支持一下中文")

    button := theme.CreateButton()
    button.SetText("点我一下")

    window.AddChild(label)
    window.AddChild(button)

    window.OnClose(driver.Terminate)

}

func main() {
    gl.StartDriver(appMain)
}

func sysFont(font string) (fontFile string) {

    sysFontPath := os.Getenv("windir") + "\\fonts\\"
    fontName, ok := FontPath[font]
    if ok {
        fontFile = sysFontPath + fontName
    }

    return
}
