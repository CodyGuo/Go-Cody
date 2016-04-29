package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Log struct {
	Index  int
	File   string
	Path   string
	Result string
	Remark string
}

type LogModel struct {
	walk.TableModelBase
	items []*Log
}

func NewLogModel() *LogModel {
	return new(LogModel)
}

func (lm *LogModel) RowCount() int {
	return len(lm.items)
}

func (lm *LogModel) Items() interface{} {
	return lm.items
}

// 结果展示
func (l *LogModel) Value(row, col int) interface{} {
	item := l.items[row]

	switch col {
	case 0:
		return item.Index
	case 1:
		return item.File
	case 2:
		return item.Path
	case 3:
		return item.Result
	case 4:
		return item.Remark
	}
	panic("unexpected col")
}

func (lm *LogModel) AddLogFile(path, file string) int {
	index := 0
	if len(lm.items) > 0 {
		index = lm.items[len(lm.items)-1].Index
	}

	lm.items = append(lm.items, &Log{
		Index:  index + 1,
		File:   file,
		Path:   path,
		Result: "未解密",
		Remark: "",
	})
	lm.PublishRowsReset()
	return index
}

func (lm *LogModel) SetResult(index int, result, remark string) {
	lm.items[index].Result = result
	lm.items[index].Remark = remark
	lm.PublishRowsReset()
}

func (lm *LogModel) GetPath(index int) string {
	return lm.items[index].Path
}

func main() {
	mw := new(MyWindow)
	mw.RunApp()
}

type MyWindow struct {
	*walk.MainWindow
	tv    *walk.TableView
	model *LogModel

	row int
}

func (mw *MyWindow) RunApp() {
	mw.model = NewLogModel()
	open := walk.NewAction()
	open.SetText("打开目录")

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "iMan高级调试日志解密工具 2.1",
		Layout:   VBox{},
		MinSize:  Size{980, 650},
		Children: []Widget{
			TableView{
				AssignTo:            &mw.tv,
				LastColumnStretched: true,
				ToolTipText:         "把日志拖放上来即可解密.",
				Columns: []TableViewColumn{
					{Title: "序号", Width: 45},
					{Title: "文件名", Width: 180},
					{Title: "文件路径", Width: 200},
					{Title: "状态", Width: 70},
					{Title: "备注", Width: 0},
				},
				Model: mw.model,
				OnCurrentIndexChanged: func() {
					mw.row = mw.tv.CurrentIndex()
				},
				ContextMenuItems: []MenuItem{
					ActionRef{&open},
				},
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}

	open.Triggered().Attach(func() {
		if len(mw.model.items) == 0 {
			exec.Command("cmd", "/c", "start", ".").Run()
		} else {
			path, _ := os.Getwd()
			exec.Command("cmd", "/c", "start", path+"\\logout\\").Run()
		}
	})

	mw.dropFiles()

	icon, _ := walk.NewIconFromResourceId(3)
	mw.SetIcon(icon)

	walk.MsgBox(mw, "提示", "把日志拖放到空白区即可解密！", walk.MsgBoxIconInformation)
	mw.Run()
}

func (mw *MyWindow) dropFiles() {
	mw.tv.DropFiles().Attach(func(files []string) {
		go func() {
			for i, file := range files {
				info, _ := os.Stat(file)
				if info.IsDir() {
					files = append(files[:i], files[i+1:]...)
					files = append(files, getFileList(file)...)
				}
			}

			var fileDir, fileName string
			index := 0
			ok := make(chan bool)
			for _, file := range files {
				fileDir = filepath.Dir(file)
				fileName = filepath.Base(file)
				fmt.Println("===========================================")
				fmt.Printf("Files: \n%v \nFileName: \n%s \nDir: \n%s\n", file, fileName, fileDir)
				index = mw.model.AddLogFile(fileDir, fileName)
				go mw.decode(ok, index, file, fileName)
				<-ok
			}
		}()
	})
}

func (mw *MyWindow) decode(ok chan bool, index int, file, fileName string) {
	filenameOnly := GetFileName(fileName)
	os.Mkdir("logout", 0755)
	outlogFile := ".\\logout\\" + filenameOnly + ".rar"

	mw.model.items[index].Result = "正在解密中..."
	nac_cmd := exec.Command("./bin/openssl", "des3", "-salt", "-d", "-k", "zaq#@!", "-in", file, "-out", outlogFile)
	err := nac_cmd.Run()
	if err != nil {
		if strings.Contains(err.Error(), "exit status 1") {
			mw.model.SetResult(index, "解密失败", "不是iMan的高级调试日志或者已经是文明日志.")
		} else {
			mw.model.SetResult(index, "未解密", "请检查bin目录下的解密程序是否完整.包含openssl.exe libeay32.dll ssleay32.dll."+err.Error())
		}
		os.RemoveAll(outlogFile)
	} else {
		mw.model.SetResult(index, "解密成功", "右键打开解密后目录.")
	}
	ok <- true
}

func GetFileName(fileName string) string {
	var filenameWithSuffix, fileSuffix, filenameOnly string
	filenameWithSuffix = path.Base(fileName)
	fileSuffix = path.Ext(filenameWithSuffix)
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)

	return filenameOnly
}

func getFileList(path string) []string {
	files := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return files
}
