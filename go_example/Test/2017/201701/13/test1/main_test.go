// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {

	mw.Run()
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
