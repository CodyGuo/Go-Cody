package main

import (
    "github.com/CodyGuo/email"
    "log"
    "net/smtp"
)

func main() {
    m := email.NewMessage("Hi, 领导.", "测试邮件发送,时间、附件都已调通.")
    m.From = "rebon@tec-development.com"
    m.To = []string{"hp131@hupu.net"}
    err := m.Attach(`C:\Users\Administrator\Desktop\4点起床：最养生和高效的时间管理.pdf`)
    if err != nil {
        log.Println(err)
    }
    err = email.Send("192.168.119.141:25", smtp.PlainAuth("", "rebon@tec-development.com", "aptech", "192.168.119.141"), m)
    if err != nil {
        log.Fatal(err)
    } else {
        log.Println("邮件发送成功...")
    }
}
