package main

import (
	"strings"
	"unsafe"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

func main() {
	var inTE, outTE *walk.TextEdit
	var screamPB *walk.PushButton

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{
						AssignTo: &inTE,
						OnKeyPress: func(key walk.Key) {
							if key == walk.KeyReturn {
								screamPB.SetFocus()
								// MustSynthesizeKeyPress(walk.KeySpace)
							}
						},
					},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				AssignTo: &screamPB,
				Text:     "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}

func MustSynthesizeKeyPress(key walk.Key) {
	kis := []win.KEYBD_INPUT{
		{Type: win.INPUT_KEYBOARD, Ki: win.KEYBDINPUT{WVk: uint16(key)}},
		{Type: win.INPUT_KEYBOARD, Ki: win.KEYBDINPUT{WVk: uint16(key), DwFlags: win.KEYEVENTF_KEYUP}},
	}

	if 0 == win.SendInput(2, unsafe.Pointer(&kis[0]), int32(unsafe.Sizeof(kis[0]))) {
		panic("SendInput failed")
	}
}
