package walk

import (
	"github.com/lxn/win"
)

var winStyle uint32 = win.WS_OVERLAPPEDWINDOW

// 获取居中坐标
func screenCenter(w, h int) (x, y int) {
	srcWidth := win.GetSystemMetrics(win.SM_CXSCREEN)
	srcHeight := win.GetSystemMetrics(win.SM_CYSCREEN)
	x = (int(srcWidth) - w) / 2
	y = (int(srcHeight) - h) / 2
	// fmt.Println("dec main:", srcWidth, srcHeight, w, h, x, y)
	return
}

// mainwindow
func NewMainWindowCody() (*MainWindow, error) {
	mw := new(MainWindow)

	if err := InitWindow(
		mw,
		nil,
		mainWindowWindowClass,
		winStyle,
		// win.WS_OVERLAPPEDWINDOW,
		win.WS_EX_CONTROLPARENT); err != nil {

		return nil, err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			mw.Dispose()
		}
	}()

	mw.SetPersistent(true)

	var err error

	if mw.menu, err = newMenuBar(mw.hWnd); err != nil {
		return nil, err
	}
	if !win.SetMenu(mw.hWnd, mw.menu.hMenu) {
		return nil, lastError("SetMenu")
	}

	tb, err := NewToolBar(mw)
	if err != nil {
		return nil, err
	}
	mw.SetToolBar(tb)

	if mw.statusBar, err = NewStatusBar(mw); err != nil {
		return nil, err
	}
	mw.statusBar.parent = nil
	mw.Children().Remove(mw.statusBar)
	mw.statusBar.parent = mw
	win.SetParent(mw.statusBar.hWnd, mw.hWnd)

	// This forces display of focus rectangles, as soon as the user starts to type.
	mw.SendMessage(win.WM_CHANGEUISTATE, win.UIS_INITIALIZE, 0)

	succeeded = true

	return mw, nil
}

func (mw *MainWindow) SetMinimizeBox(minbox bool) {
	if !minbox {
		winStyle = winStyle - win.WS_MINIMIZEBOX
	}
}

func (mw *MainWindow) SetMaximizeBox(maxbox bool) {
	if !maxbox {
		winStyle = winStyle - win.WS_MAXIMIZEBOX
	}
}

func (mw *MainWindow) SetFixedSize(fixed bool) {
	if fixed {
		winStyle = winStyle - win.WS_SIZEBOX
	}
}

func (mw *MainWindow) SetScreenCenter(center bool) {
	if center {
		screenStyleX, screenStyleY := screenCenter(mw.Width(), mw.Height())
		mw.SetX(screenStyleX)
		mw.SetY(screenStyleY)
	}
}

// dialog

func NewDialogCody(owner Form) (*Dialog, error) {
	return newDialogWithStyleCody(owner, 0)
}

func newDialogWithStyleCody(owner Form, style uint32) (*Dialog, error) {
	dlg := &Dialog{
		FormBase: FormBase{
			owner: owner,
		},
	}

	if err := InitWindow(
		dlg,
		owner,
		dialogWindowClass,
		winStyle|style,
		// win.WS_CAPTION|win.WS_SYSMENU|style,
		0); err != nil {
		return nil, err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			dlg.Dispose()
		}
	}()

	dlg.centerInOwnerWhenRun = owner != nil

	// This forces display of focus rectangles, as soon as the user starts to type.
	dlg.SendMessage(win.WM_CHANGEUISTATE, win.UIS_INITIALIZE, 0)

	dlg.result = DlgCmdNone

	succeeded = true

	return dlg, nil
}

// 设置最小化
func (dlg *Dialog) SetMinimizeBox(minbox bool) {
	if !minbox {
		winStyle = winStyle - win.WS_MINIMIZEBOX
	}
}

// 设置最大化
func (dlg *Dialog) SetMaximizeBox(maxbox bool) {
	if !maxbox {
		winStyle = winStyle - win.WS_MAXIMIZEBOX
	}
}

// 设置窗体固定大小
func (dlg *Dialog) SetFixedSize(fixed bool) {
	if fixed {
		winStyle = winStyle - win.WS_SIZEBOX
	}
}

// 该函数将焦点切换指定的窗口，并将其带到前台。
func (wb *WindowBase) SwitchToThisWindow(switchto bool) {
	win.ShowWindow(wb.hWnd, win.SW_RESTORE)
}

// SetForegroundWindow函数将创建指定窗口的线程设置到前台，并且激活该窗口。
// 键盘输入转向该窗口，并为用户改各种可视的记号。系统给创建前台窗口的线程分配的权限稍高于其他线程。
func (wb *WindowBase) SetForegroundWindow() error {
	if !win.SetForegroundWindow(wb.hWnd) {
		return lastError("SetForegroundWindow")
	}

	return nil
}

// 设置窗体居中
func (dlg *Dialog) SetScreenCenter(center bool) {
	if center {
		screenStyleX, screenStyleY := screenCenter(dlg.Width(), dlg.Height())
		dlg.SetX(screenStyleX)
		dlg.SetY(screenStyleY)
	}
}
