# Go Walk ui
walk ui 设计
https://github.com/lxn/walk

# 目录

## 1.1 mainwindow 新增部分功能
   说明：文件是cody.go，放在lxn/walk目录下
   // 在 NewMainWindowCody 前使用
   SetMinimizeBox 设置是否可以最小化
   SetMaximizeBox 设置是否可以最大化
   SetFixedSize   设置窗体是否固定大小
   NewMainWindowCody  创建自定义 mainwindow
   // 在 NewMainWindowCody 后使用
   SwitchToThisWindow 窗体到最前
   SetForegroundWindow 窗体到前台
   // 在设置完窗体大小后使用
   SetScreenCenter 设置窗体是否在屏幕居中

## 1.2 Dialog新增部分功能
   说明：文件是cody.go，放在lxn/walk目录下
   // 在 NewDialogCody 前使用
   SetMinimizeBox 设置是否可以最小化
   SetMaximizeBox 设置是否可以最大化
   SetFixedSize   设置窗体是否固定大小
   NewDialogCody  创建自定义dialog
   // 在 NewDialogCody 后使用
   SwitchToThisWindow 窗体到最前
   SetForegroundWindow 窗体到前台
   // 在 SetClientSize 后使用
   SetScreenCenter 设置窗体是否在屏幕居中

## 1.3 declarative 新增部分功能
   说明：文件是codytive.go，放在lxn/walk/declarative目录下
   CreateCody