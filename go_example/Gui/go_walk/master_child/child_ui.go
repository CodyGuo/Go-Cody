package main

import (
    "github.com/lxn/walk"
)

var (
    dlg *DlgChild
)

type ChildUI struct {
    NameLb *walk.Label
}

type DlgChild struct {
    *walk.Dialog

    ui  ChildUI
}

func (dlg *DlgChild) init(owner walk.Form, title string) (err error) {
    dlg.Dialog, _ = walk.NewDialogWithFixedSize(nil)
    succeeded := false
    defer func() {
        if !succeeded {
            dlg.Dispose()
        }
    }()

    dlg.SetClientSize(walk.Size{220, 110})

    dlg.SetTitle("子窗体 - " + title)

    dlg.ui.NameLb, _ = walk.NewLabel(dlg)
    dlg.ui.NameLb.SetText("子窗体名字: " + title)
    dlg.ui.NameLb.SetBounds(walk.Rectangle{30, 40, 180, 30})

    succeeded = true

    return nil
}
