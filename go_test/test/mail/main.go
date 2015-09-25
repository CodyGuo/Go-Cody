package main

import (
    "fmt"
    "ioutil"
    "net/smtp"
    "path/filepath"
    "strings"
)

type Attachment struct {
    Filename string
    Data     []byte
    Inline   bool
}

func SendToMail(user, password, host, to, subject, body, file, mailtype string) error {
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
    } else {
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }

    data, err := ioutil.ReadFile(file)
    if err != nil {
        return err
    }

    _, filename := filepath.Split(file)

    var Attachments map[string]*Attachment

    Attachments[filename] = &Attachment{
        Filename: filename,
        Data:     data,
        Inline:   inline,
    }

    msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    return err
}

func main() {
    user := "rebon@tec-development.com"
    password := "aptech"
    host := "192.168.119.141:25"
    to := "guoxiao219@126.com"

    subject := "使用Golang发送邮件"

    body := `
        <html>
        <body>
        <h3>
        "Test send to email"
        </h3>
        </body>
        </html>
        `
    fmt.Println("send email")
    err := SendToMail(user, password, host, to, subject, body, "html")
    if err != nil {
        fmt.Println("Send mail error!")
        fmt.Println(err)
    } else {
        fmt.Println("Send mail success!")
    }

}
