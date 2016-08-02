package main

import (
    "encoding/gob"
    "log"
    "net/smtp"
    "os"
    "time"
)
import (
    "github.com/codyguo/email"
)

var MS MailSend

type MailSend struct {
    SendList []string
    UserName string
    Passwd   string
    Smtp     string
    Port     string
    Subject  string
    Body     string
    Adjunct  string
}

func (ms *MailSend) LoadData() {
    file, err := os.Open("data.dat")
    defer file.Close()
    if err != nil {
        ms.SendList = []string{"要发送的邮箱，每行一个"}
        ms.UserName = "用户名"
        ms.Passwd = "用户密码"
        ms.Smtp = "SMTP服务器"
        ms.Port = "25"
        ms.Subject = "邮件主题"
        ms.Body = "邮件内容"
        ms.Adjunct = ""

        return
    }

    dec := gob.NewDecoder(file)
    err = dec.Decode(ms)
    if err != nil {
        ms.SendList = []string{"要发送的邮箱，每行一个"}
        ms.UserName = "用户名"
        ms.Passwd = "用户密码"
        ms.Smtp = "SMTP服务器"
        ms.Port = "25"
        ms.Subject = "邮件主题"
        ms.Body = "邮件内容"
        ms.Adjunct = ""
    }

    file.Close()
}

func (ms *MailSend) SaveData() {

    configFile := "data.dat"
    if ms.Exist(configFile) {
        err := os.Remove(configFile)
        if err != nil {
            os.Exit(1)
        }
    }

    file, err := os.Create(configFile)
    defer file.Close()
    if err != nil {
        log.Println(err)
    }

    enc := gob.NewEncoder(file)
    err = enc.Encode(ms)
    if err != nil {
        log.Println(err)
    }

    file.Close()
}

func (ms *MailSend) SendMail() (err error) {
    nowTime := time.Now()
    layout := "2006-01-02 03:04:05"

    m := email.NewMessage(ms.Subject+" "+nowTime.Format(layout), ms.Body)
    m.From = ms.UserName
    m.To = []string{"hp131@hupu.net"}
    if ms.Adjunct != "" {
        err = m.Attach(ms.Adjunct)
        if err != nil {
            return err
            os.Exit(1)
        }
    }

    err = email.Send(ms.Smtp+":"+ms.Port, smtp.PlainAuth("", ms.UserName, ms.Passwd, ms.Smtp), m)
    if err != nil {
        return err
        os.Exit(1)
    }
    ms.Adjunct = ""

    return nil
}

func (ms *MailSend) Exist(filename string) bool {
    _, err := os.Stat(filename)

    return err == nil || os.IsExist(err)
}
