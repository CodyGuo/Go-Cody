# Go walk
学习`gowi`过程中遇到的坑

## 1.1 Onclicke

OnClicked: mw.On...Click,
Go中方法也可以赋值，跟 a = 1 没什么差别
但是你加 ()，是方法返回值赋值

    Controls: []Control{
        Button{
            AssignTo:  &mw.button1,
            ID:        IDD_BUTTON1,
            OnClicked: mw.OnButton1Clicked,
        },