package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "os"

    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

var (
    UserConfig User
)

var (
    configFile = "config.data"
)

type User struct {
    UserName  string
    Password  string
    JobNumber string
    Name      string
    Email     string
    Phone     string
}

func (u *User) LoadConfig() {
    file, err := os.Open(configFile)
    if err != nil {
        log.Println("打开配置文件失败." + err.Error())
        UserConfig.UserName = "用户名"
        UserConfig.Password = "密码"
        UserConfig.JobNumber = "工号"
        UserConfig.Name = "姓名"
        UserConfig.Email = "邮箱"
        UserConfig.Phone = "联系电话"
        return
    }
    defer file.Close()

    dec := gob.NewDecoder(file)
    err = dec.Decode(&UserConfig)
    if err != nil {
        UserConfig.UserName = "用户名"
        UserConfig.Password = "密码"
        UserConfig.JobNumber = "工号"
        UserConfig.Name = "姓名"
        UserConfig.Email = "邮箱"
        UserConfig.Phone = "联系电话"
    }
}

func (u *User) SaveConfig() {
    file, err := os.Create(configFile)
    if err != nil {
        log.Fatal("创建配置文件失败." + err.Error())
    }
    defer file.Close()

    enc := gob.NewEncoder(file)
    err = enc.Encode(UserConfig)
    if err != nil {
        log.Fatal("保存配置文件失败." + err.Error())
    }

}

func main() {
    UserConfig.LoadConfig()
    fmt.Println(UserConfig.UserName)

    UserConfig.UserName = "hp131@hupu.net"
    UserConfig.SaveConfig()

    fmt.Println(UserConfig.UserName, UserConfig.Password, UserConfig.Name)

    fmt.Print("加载完毕.\n")

}
