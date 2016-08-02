package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var mw *walk.MainWindow
	var c1 *walk.Composite
	var c2 *walk.Composite

	bmp, err := walk.NewBitmapFromFile("../../img/plus.png")
	if err != nil {
		panic(err)
	}
	defer bmp.Dispose()

	MainWindow{
		AssignTo: &mw,
		Title:    "Background Example",
		Layout:   VBox{MarginsZero: true},
		MinSize:  Size{300, 400},
		Children: []Widget{
			Composite{
				AssignTo: &c1,
				Layout:   VBox{},
				Children: []Widget{
					TextEdit{},
				},
			},
			Composite{
				AssignTo: &c2,
				Layout:   VBox{},
				Children: []Widget{
					TextEdit{},
				},
			},
			ImageView{
				Image: bmp,
			},
		},
	}.Create()

	scb, err := walk.NewSolidColorBrush(walk.RGB(255, 0, 0))
	if err != nil {
		panic(err)
	}
	defer scb.Dispose()

	c1.SetBackground(scb)

	bmb, err := walk.NewBitmapBrush(bmp)
	if err != nil {
		panic(err)
	}
	defer bmb.Dispose()

	c2.SetBackground(bmb)

	mw.Show()
	mw.Run()
}
