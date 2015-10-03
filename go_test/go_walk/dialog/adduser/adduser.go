// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    "log"
    "time"
)

import (
    "github.com/lxn/walk"
)

func main() {
    if _, err := RunMyDialog(nil); err != nil {
        log.Fatal(err)
    }
}

type MyDialog struct {
    *walk.Dialog
    ui  myDialogUI
}

func (dlg *MyDialog) setState(state walk.PIState) {
    if err := dlg.ProgressIndicator().SetState(state); err != nil {
        log.Print(err)
    }
}

func RunMyDialog(owner walk.Form) (int, error) {
    dlg := new(MyDialog)
    if err := dlg.init(owner); err != nil {
        return 0, err
    }

    dlg.ui.uploadBtn.Clicked().Attach(func() {
        go func() {
            dlg.ProgressIndicator().SetTotal(100)
            var i uint32
            for i = 0; i < 100; i++ {
                fmt.Println("SetProgress", i)
                time.Sleep(100 * time.Millisecond)
                if err := dlg.ProgressIndicator().SetCompleted(i); err != nil {
                    log.Print(err)
                }
            }
        }()
    })
    dlg.SetScreenCenter(true)
    return dlg.Run(), nil
}
