package main

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var prevPB *walk.PushButton
	var nextPB *walk.PushButton

	count := 0
	MainWindow{
		Title:  "Sample",
		Layout: VBox{},
		Children: []Widget{
			PushButton{
				AssignTo: &prevPB,
				Text:     "Previous",
				Enabled:  false,
				OnClicked: func() {
					if count != 0 {
						count -= 1
						nextPB.SetEnabled(true)
					} else {
						prevPB.SetEnabled(false)
					}
				},
			},
			PushButton{
				AssignTo: &nextPB,
				Text:     "Next",
				OnClicked: func() {
					if count != 10 {
						count += 1
						prevPB.SetEnabled(true)
					} else {
						nextPB.SetEnabled(false)
					}
				},
			},
			PushButton{
				Text: "Check",
				OnClicked: func() {
					fmt.Println("nextButton:", nextPB.Enabled())
					fmt.Println("previousButton:", prevPB.Enabled())
					fmt.Println("Count:", count)
				},
			},
		},
	}.Run()
}
