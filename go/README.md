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

### 1.2 最小编译go程序
    把Go程序变小的办法是：
    go build -ldflags "-s -w"

    相关解释：
    -s去掉符号表,panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果，
    -w去掉DWARF调试信息，得到的程序就不能用gdb调试了。 -s和-w也可以分开使用.
### 1.3 更新git上fork的代码
    当然，那是完全不用命令行的办法，其实我还是更推荐命令行，流程如下：

    首先要先确定一下是否建立了主repo的远程源：
    git remote -v
    如果里面只能看到你自己的两个源(fetch 和 push)，那就需要添加主repo的源：
    git remote add upstream URL
    git remote -v
    然后你就能看到upstream了。
    如果想与主repo合并：
    git fetch upstream
    git merge upstream/master
    提交
    git push origin master

### 1.4 windows程序（UAC）以管理员身份运行
    1> go get github.com/akavel/rsrc
    2> 把windows目录下的nac.manifest 文件拷贝到当前windows项目根目录
    3> rsrc -manifest nac.manifest -o nac.syso
    4> go build

    nac.mainfest的内容为：

    <?xml version="1.0" encoding="UTF-8" standalone="yes"?>
    <assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
    <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
    <requestedPrivileges>
    <requestedExecutionLevel level="requireAdministrator"/>
    </requestedPrivileges>
    </security>
    </trustInfo>
    </assembly>

