package main

import (
    "github.com/lxn/walk"
    // . "github.com/lxn/walk/declarative"
)

var (
    ltu, ltp, ltn, lte, ltphone *walk.LineEdit
    setting                     *walk.Composite
    addUserBtn                  *walk.PushButton
    db                          *walk.DataBinder
    ep                          walk.ErrorPresenter
    ltj                         *walk.NumberEdit
)

type MyMainWindow struct {
    *walk.MainWindow
}

var (
    UserConfig = &User{}
    configFile = "config.data"
)

type User struct {
    UserName  string
    Password  string
    JobNumber int
    Name      string
    Email     string
    Phone     string
}
