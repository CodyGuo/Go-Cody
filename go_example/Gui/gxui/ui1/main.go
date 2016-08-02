// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "github.com/google/gxui"
    "github.com/google/gxui/drivers/gl"
    "github.com/google/gxui/math"
    "github.com/google/gxui/samples/flags"
)

func appMain(driver gxui.Driver) {
    theme := flags.CreateTheme(driver)

    window := theme.CreateWindow(800, 600, "iMan升级")
    window.OnClose(driver.Terminate)
    window.SetScale(flags.DefaultScaleFactor)
    window.SetPadding(math.Spacing{L: 50, R: 50, T: 50, B: 50})

    button := theme.CreateButton()
    button.SetHorizontalAlignment(gxui.AlignCenter)
    button.SetSizeMode(gxui.Fill)

    toggle := func() {
        fullscreen := !window.Fullscreen()
        window.SetFullscreen(fullscreen)
        if fullscreen {
            button.SetText("窗口化")
        } else {
            button.SetText("全屏")
        }
    }

    box := theme.CreateTextBox()
    box.SetText("盒子")

    button.SetText("全屏")
    button.OnClick(func(gxui.MouseEvent) { toggle() })
    window.AddChild(button)
    window.AddChild(box)
}

func main() {
    gl.StartDriver(appMain)
}
