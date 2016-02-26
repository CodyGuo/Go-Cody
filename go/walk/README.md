# Go Walk ui
walk ui 设计
https://github.com/lxn/walk

# 目录

## 1.1 Dialog新增部分功能
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