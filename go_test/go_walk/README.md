# 1.user32.go 里定义的方法和常量
    HWND CreateWindow(LPCTSTR lpClassName,LPCTSTR lpWindowName,DWORD dwStyle,int x,int y,int nWidth，int nHeight，HWND hWndParent,HMENU hMenu，HANDLE hlnstance，LPVOID lpParam）；

    其中dwStyle指定WS_OVERLAPPEDWINDOW&~WS_MAXINIZEBOX(这样做是取消最大化的参数)
    其实你可以选中WS_OVERLAPPEDWINDOW，点击GOTO DEFIONTION,在出现WS_OVERLAPPEDWINDOW的定义为，其中就有WS_MAXINIZEBOX，&～的意思是去掉WS_MAXINIZEBOX

    func InitWindow ¶

    func InitWindow(window, parent Window, className string, style, exStyle uint32) error
    GetSystemMetrics constants

    ❖
    const (
        SW_HIDE            = 0
        SW_NORMAL          = 1
        SW_SHOWNORMAL      = 1
        SW_SHOWMINIMIZED   = 2
        SW_MAXIMIZE        = 3
        SW_SHOWMAXIMIZED   = 3
        SW_SHOWNOACTIVATE  = 4
        SW_SHOW            = 5
        SW_MINIMIZE        = 6
        SW_SHOWMINNOACTIVE = 7
        SW_SHOWNA          = 8
        SW_RESTORE         = 9
        SW_SHOWDEFAULT     = 10
        SW_FORCEMINIMIZE   = 11
    )

    win.WS_OVERLAPPEDWINDOW, 这个就是设置标题栏的


# 2. walk增加禁止最大化、最小化、固定窗体大小
    修改 lxn/walk/mainwindow.go
    1> 增加全局变量
       var winStyle uint32 = win.WS_OVERLAPPEDWINDOW
    2> 修改 NewMainWindow() 函数下的 win.WS_OVERLAPPEDWINDOW 为 winStyle
        func NewMainWindow() (*MainWindow, error) {
        mw := new(MainWindow)

        if err := InitWindow(
            mw,
            nil,
            mainWindowWindowClass,
            winStyle, //     WS_MINIMIZEBOX = 0X00020000 WS_MAXIMIZEBOX = 0X00010000 WS_SIZEBOX = 0X00040000
            win.WS_EX_CONTROLPARENT); err != nil {

            return nil, err
        }
    3> Men()方法前增加以下三个方法：
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
    4> 使用方法：在Crete Mainwindow之前使用
        mw := new(MyMainWindow)
        mw.SetMaximizeBox(false)
        mw.SetMinimizeBox(true)
        mw.SetFixedSize(true)

# 3. 获取居中屏幕坐标
    func ScreenCenter(w, h int) (x, y int) {
        fmt.Println(win.GetSystemMetrics(win.SM_CXSCREEN))
        fmt.Println(win.GetSystemMetrics(win.SM_CYSCREEN))
        srcWidth := win.GetSystemMetrics(win.SM_CXSCREEN)
        srcHeight := win.GetSystemMetrics(win.SM_CYSCREEN)
        x = (int(srcWidth) - w) / 2
        y = (int(srcHeight) - h) / 2
        fmt.Println("screencenter:")
        fmt.Println(x, y)
        return
    }

    //构建窗体时
    case WM_CREATE:
    int scrWidth,scrHeight;
    RECT rect;
    //获得屏幕尺寸
    scrWidth=GetSystemMetrics(SM_CXSCREEN);
    scrHeight=GetSystemMetrics(SM_CYSCREEN);
    //获取窗体尺寸
    GetWindowRect(hWnd,&rect);
    rect.left=(scrWidth-rect.right)/2;
    rect.top=(scrHeight-rect.bottom)/2;

# 4. 增加屏幕居中属性
    修改 lxn\walk\declarative\mainwindow.go
    1> 结构体 MainWindow 增加属性
    type MainWindow struct {
        ... ...
        ScreenCenter     bool
    }
    2> 增加函数：
    // 获取居中坐标
    func ScreenCenter(w, h int) (x, y int) {
        srcWidth := win.GetSystemMetrics(win.SM_CXSCREEN)
        srcHeight := win.GetSystemMetrics(win.SM_CYSCREEN)
        x = (int(srcWidth) - w) / 2
        y = (int(srcHeight) - h) / 2
        // fmt.Println("dec main:", srcWidth, srcHeight, w, h, x, y)
        return
    }
    3> 在 return builder.InitWidget(tlwi, w, func() error 下增加坐标初始化居中位置
        return builder.InitWidget(tlwi, w, func() error {
        // 主窗体是否居中
        if mw.ScreenCenter {
            x, y := ScreenCenter(mw.MinSize.Width, mw.MinSize.Height)
            fmt.Println("dec main : w h", mw.MinSize.Width, mw.MinSize.Height, mw.MaxSize.Width, mw.MaxSize.Height, mw.Size.toW().Width, mw.Size.toW().Height)
            if err := w.SetX(x); err != nil {
                return err
            }
            if err := w.SetY(y); err != nil {
                return err
            }
        }
    4> 使用方法：ScreenCenter: true,
        if err := (MainWindow{
        AssignTo: &mw.MainWindow,
        DataBinder: DataBinder{
            AssignTo:       &db,
            DataSource:     UserConfig,
            ErrorPresenter: ErrorPresenterRef{&ep},
        },
        Title:        "iMan - 测试程序",
        ScreenCenter: true,
        MinSize:      Size{300, 262},
        Layout:       VBox{Spacing: 2},
        Children: []Widget{

# 4. walk 中dialog 增加限制窗口最大化、最小化、固定窗体大小.屏幕居中
    1> 修改 lxn/walk/dialog.go,把 win.WS_CAPTION|win.WS_SYSMENU|win.WS_THICKFRAME, 修改为 winStyle,
        if err := InitWindow(
        dlg,
        owner,
        dialogWindowClass,
        winStyle,
        // win.WS_CAPTION|win.WS_SYSMENU|win.WS_THICKFRAME,
    2> 在 DefaultButton 方法前加入三个方法：SetMinimizeBox SetMaximizeBox SetFixedSize
        func (dlg *Dialog) SetMinimizeBox(minbox bool) {
            if !minbox {
                winStyle = winStyle - win.WS_MINIMIZEBOX
            }
        }

        func (dlg *Dialog) SetMaximizeBox(maxbox bool) {
            if !maxbox {
                winStyle = winStyle - win.WS_MAXIMIZEBOX
            }
        }

        func (dlg *Dialog) SetFixedSize(fixed bool) {
            if fixed {
                winStyle = winStyle - win.WS_SIZEBOX
            }
        }
    3> 屏幕居中,修改 lxn\walk\declarative\dialog.go
        1>> 结构体 Dialog 增加属性
            type Dialog struct {
                    ... ...
                    ScreenCenter     bool
            }
        2>> 在 return builder.InitWidget(tlwi, w, func() error 下增加坐标初始化居中位置
            return builder.InitWidget(tlwi, w, func() error {
            // 主窗体是否居中
            if d.ScreenCenter {
                x, y := ScreenCenter(d.MinSize.Width, d.MinSize.Height)
                // fmt.Println("dec main : w h", d.MinSize.Width, d.MinSize.Height, d.MaxSize.Width, d.MaxSize.Height, d.Size.toW().Width, d.Size.toW().Height)
                if err := w.SetX(x); err != nil {
                    return err
                }
                if err := w.SetY(y); err != nil {
                    return err
                }
            }
        3>> 使用方法
            return Dialog{
            AssignTo:      &dlg,
            Title:         "Animal Details",
            DefaultButton: &acceptPB,
            CancelButton:  &cancelPB,
            DataBinder: DataBinder{
                AssignTo:       &db,
                DataSource:     animal,
                ErrorPresenter: ErrorPresenterRef{&ep},
            },
            MinSize:      Size{300, 300},
            ScreenCenter: true,

        4>> 修改 lxn\walk\dialog.go 增加函数 ScreenCenter
            // 获取居中坐标
            func ScreenCenter(w, h int) (x, y int) {
                srcWidth := win.GetSystemMetrics(win.SM_CXSCREEN)
                srcHeight := win.GetSystemMetrics(win.SM_CYSCREEN)
                x = (int(srcWidth) - w) / 2
                y = (int(srcHeight) - h) / 2
                // fmt.Println("dec main:", srcWidth, srcHeight, w, h, x, y)
                return
            }
        5>> 增加方法
            func (dlg *Dialog) SetScreenCenter(center bool) {
                if center {
                    screenStyleX, screenStyleY := ScreenCenter(dlg.Width(), dlg.Height())
                    dlg.SetX(screenStyleX)
                    dlg.SetY(screenStyleY)
                    // fmt.Println("walk dialog:", dlg.Width(), dlg.Height(), screenStyleX, screenStyleY)
                }
            }
        6>> 使用方法 在dlg.Run() 前使用
            dlg.SetScreenCenter(false)
            return dlg.Run(), nil

