package main

import (
	"log"
	"time"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	width  = 800
	height = 600
)

type MyWindow struct {
	*walk.MainWindow
}

type People struct {
	Index   int
	Name    string
	Age     int
	Date    time.Time
	checked bool
}

type PeopleModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	tems       []*People
}

func (mw *MyWindow) RunApp() {
	mw.SetMaximizeBox(false)
	mw.SetFixedSize(true)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Layout:   HBox{},
		MinSize:  Size{width, height},
		Children: []Widget{
			PushButton{
				Text: "button",
				OnClicked: func() {
					log.Println("Hello world.")
				},
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}

func main() {
	mw := new(MyWindow)
	mw.RunApp()
}
