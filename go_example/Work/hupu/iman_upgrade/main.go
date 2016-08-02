package main

import (
    "fmt"
    "os"

    "path"

    "github.com/Unknwon/goconfig"
)

func main() {
    // 获取当前setup路径
    filepath := os.Args[0]
    setupPath, _ := path.Split(filepath)
    fmt.Println(setupPath)

    // 安装xinetd
    xinetInserr := GetRpmInfo("xinetd-2.3.14-39.el6_4.i686")
    if xinetInserr != nil {
        err := InstallRpm(setupPath + "rpm/xinetd-2.3.14-39.el6_4.i686.rpm")
        if err != nil {
            fmt.Println(err)
        }
    }
    // 安装tftp
    tftpInserr := GetRpmInfo("tftp-server-0.49-7.el6.i686")
    if tftpInserr != nil {
        err := InstallRpm(setupPath + "rpm/tftp-server-0.49-7.el6.i686.rpm")
        if err != nil {
            fmt.Println(err)
        }
    }
    // 修改tftp配置文件
    tftperr := CopyFile(setupPath+"config/tftp", "/etc/xinetd.d/tftp")
    if tftperr != nil {
        fmt.Println(tftperr)
    }

    // 创建tftp根目录
    os.Mkdir("/data/tftproot", 000)
    os.Chmod("/data/tftproot", 0777)

    // 启动tftp服务
    StartServers("xinetd")

    /******************************************************/

    // 修改iptables防火墙配置
    configerr := CopyFile(setupPath+"config/iptables-config", "/etc/sysconfig/iptables-config")
    if configerr != nil {
        fmt.Println(configerr)
    }
    // 修改iptables防火墙配置
    iptableserr := CopyFile(setupPath+"config/iptables", "/etc/sysconfig/iptables")
    if iptableserr != nil {
        fmt.Println(iptableserr)
    }

    // 重启iptables防火墙
    StartServers("iptables")

    /******************************************************/

    // 修改数据库接收数据的大小
    mysql, mysqlerr := goconfig.LoadConfigFile("/etc/my.cnf")
    if mysqlerr != nil {
        fmt.Println(mysqlerr)
    }
    mysql.SetValue("mysqld", "max_allowed_packet", "32M")
    goconfig.SaveConfigFile(mysql, "/etc/my.cnf")

    /******************************************************/

    // 升级tomcat重启文件
    tomcatErr := CopyFile(setupPath+"config/tomcat", "/etc/init.d/tomcat")
    if tomcatErr != nil {
        fmt.Println(tomcatErr)
    }
    os.Chmod("/etc/init.d/tomcat", 0755)

    webErr := CopyFile(setupPath+"config/web.xml", "/nac/web/tomcat/conf/web.xml")
    if webErr != nil {
        fmt.Println(webErr)
    }
    os.Chmod("/nac/web/tomcat/conf/web.xml", 0755)

}
