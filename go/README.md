# Go
学习Go过程中遇到的坑
教程来自：http://bbs.csdn.net/topics/390601048/

## 1.1 交叉编译
-------------------------------------------------------------------
* 在安装好go后，建立交叉编译环境还需要一个重要的工具链gcc，推荐使用mingw，下载地址如下http://sourceforge.net/projects/mingw/files/Installer/mingw-get-inst/mingw-get-inst-20120426/mingw-get-inst-20120426.exe/download
* 安装好后，下文假设安装在D:\MinGW下，将D:\MinGW\bin添加到系统环境变量 %PATH% 中。
* 假设Go安装在D:\go下面，将下面的批处理文件放置到D:\go\src下后执行。
    set CGO_ENABLED=0
    ::x86块
    set GOARCH=386
    set GOOS=windows
    call make.bat --no-clean

    set GOOS=linux
    call make.bat --no-clean

    set GOOS=freebsd
    call make.bat --no-clean

    set GOOS=darwin
    call make.bat --no-clean
    ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

    ::x64块
    set GOARCH=amd64
    set GOOS=linux
    call make.bat --no-clean
    ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

    ::arm块
    set GOARCH=arm
    set GOOS=linux
    call make.bat --no-clean
    ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

    set GOARCH=386
    set GOOS=windows
    go get github.com/nsf/gocode
    pause
-------------------------------------------------------------------
* 运行完毕将会产生交叉编译环境列表如下(不完全，请根据自己需要修改)
* x86的windows/linux/darwin(mac os)/freebsd
* x64的linux
* arm的linux(android)
* 另外还将安装gocode用于代码提示。
-------------------------------------------------------------------
* 最后提供一个例子用于在windows上交叉编译x86的linux可执行文件

    set GOPATH=你的工程目录
    set GOARCH=386
    set GOOS=linux
    go build
    pause

* 将上述批处理文件放置到你的.go源文件所在目录下运行，即可产生对应平台的可执行文件。
* 修改GOARCH及GOOS来产生对应平台的可执行文件，可以自行完善批处理文件，做到编译、strip、upx、打包一条龙。