// THIS FILE WAS GENERATED BY A TOOL, DO NOT EDIT!

package main

import (
    "github.com/lxn/walk"
)

type myDialogUI struct {
    uploadBtn *walk.PushButton
    fileLe    *walk.LineEdit
    uploadGb  *walk.GroupBox
}

func (w *MyDialog) init(owner walk.Form) (err error) {
    w.SetMinimizeBox(true)
    w.SetMaximizeBox(false)
    w.SetFixedSize(true)

    if w.Dialog, err = walk.NewDialog(owner); err != nil {
        return err
    }

    succeeded := false
    defer func() {
        if !succeeded {
            w.Dispose()
        }
    }()

    var font *walk.Font
    if font == nil {
        font = nil
    }

    w.SetName("windows")
    if err := w.SetClientSize(walk.Size{407, 178}); err != nil {
        return err
    }
    if err := w.SetTitle(`iMan-测试程序`); err != nil {
        return err
    }

    if w.ui.uploadGb, err = walk.NewGroupBox(w); err != nil {
        return err
    }

    if err := w.ui.uploadGb.SetBounds(walk.Rectangle{20, 20, 371, 141}); err != nil {
        return err
    }

    w.ui.uploadGb.SetTitle("升级包上传")

    // fileLe
    if w.ui.fileLe, err = walk.NewLineEdit(w.ui.uploadGb); err != nil {
        return err
    }

    if err := w.ui.fileLe.SetBounds(walk.Rectangle{20, 40, 231, 31}); err != nil {
        return err
    }

    // uploadFileBt
    if w.ui.uploadBtn, err = walk.NewPushButton(w.ui.uploadGb); err != nil {
        return err
    }

    w.ui.uploadBtn.SetName("uploadFileBt")
    if err := w.ui.uploadBtn.SetBounds(walk.Rectangle{270, 40, 75, 31}); err != nil {
        return err
    }

    if err := w.ui.uploadBtn.SetText(`上传`); err != nil {
        return err
    }

    succeeded = true

    return nil
}
