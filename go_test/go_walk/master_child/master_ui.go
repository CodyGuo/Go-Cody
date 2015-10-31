package main

import (
    "github.com/lxn/walk"
)

type MyWindowUI struct {
    SecTitileLb *walk.Label
    SecTitileLe *walk.LineEdit

    CreateBtn *walk.PushButton
}

type MyWindow struct {
    *walk.MainWindow

    ui  MyWindowUI
}

func (mw *MyWindow) init() (err error) {
    mw.SetMinimizeBox(true)
    mw.SetMaximizeBox(false)
    mw.SetFixedSize(true)

    mw.MainWindow, _ = walk.NewMainWindow()
    succeeded := false
    defer func() {
        if !succeeded {
            mw.Dispose()
        }
    }()

    mw.SetClientSize(walk.Size{260, 160})

    mw.SetTitle("主窗体")

    mw.ui.SecTitileLb, _ = walk.NewLabel(mw)
    mw.ui.SecTitileLb.SetText("子窗体：")
    mw.ui.SecTitileLb.SetBounds(walk.Rectangle{10, 50, 50, 20})

    mw.ui.SecTitileLe, _ = walk.NewLineEdit(mw)
    mw.ui.SecTitileLe.SetBounds(walk.Rectangle{60, 50, 160, 20})
    mw.ui.SecTitileLe.SetWidth(150)

    mw.ui.CreateBtn, _ = walk.NewPushButton(mw)
    mw.ui.CreateBtn.SetText("生成子窗体")
    mw.ui.CreateBtn.SetBounds(walk.Rectangle{90, 100, 75, 25})

    succeeded = true

    return nil
}
