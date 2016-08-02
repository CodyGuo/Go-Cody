// Multi-SendEmail project main.go
package main

import (
    "bufio"
    "encoding/gob"
    "errors"
    "fmt"
    "io"
    "net/smtp"
    "os"
    "strconv"
    "strings"
    "time"
)

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

type ShuJu struct {
    Name    string
    Pwd     string
    Host    string
    Subject string
    Body    string
    Send    string
}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
    fmt.Println("Send to " + to)
    //fmt.Println(user, password, host, to, subject, body, mailtype)
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/html;charset=UTF-8"
    } else {
        content_type = "Content-Type: text/plain;charset=UTF-8"
    }
    body = strings.TrimSpace(body)
    msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    if err != nil {
        fmt.Println(err.Error())
    }
    return err
}

func readLine2Array(filename string) ([]string, error) {
    result := make([]string, 0)
    file, err := os.Open(filename)
    if err != nil {
        return result, errors.New("Open file failed.")
    }
    defer file.Close()
    bf := bufio.NewReader(file)
    for {
        line, isPrefix, err1 := bf.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                return result, errors.New("ReadLine no finish")
            }
            break
        }
        if isPrefix {
            return result, errors.New("Line is too long")
        }
        str := string(line)
        result = append(result, str)
    }
    return result, nil
}

func DelArrayVar(arr []string, str string) []string {
    str = strings.TrimSpace(str)
    for i, v := range arr {
        v = strings.TrimSpace(v)
        if v == str {
            if i == len(arr) {
                return arr[0 : i-1]
            }
            if i == 0 {
                return arr[1:len(arr)]
            }
            a1 := arr[0:i]
            a2 := arr[i+1 : len(arr)]
            return append(a1, a2...)
        }
    }
    return arr
}

func LoadData() {
    fmt.Println("LoadData")
    file, err := os.Open("data.dat")
    defer file.Close()
    if err != nil {
        fmt.Println(err.Error())
        SJ.Name = "用户名"
        SJ.Pwd = "用户密码"
        SJ.Host = "SMTP服务器:端口"
        SJ.Subject = "邮件主题"
        SJ.Body = "邮件内容"
        SJ.Send = "要发送的邮箱，每行一个"
        return
    }

    dec := gob.NewDecoder(file)
    err2 := dec.Decode(&SJ)
    if err2 != nil {
        fmt.Println(err2.Error())
        SJ.Name = "用户名"
        SJ.Pwd = "用户密码"
        SJ.Host = "SMTP服务器:端口"
        SJ.Subject = "邮件主题"
        SJ.Body = "邮件内容"
        SJ.Send = "要发送的邮箱，每行一个"
    }
}

func SaveData() {
    fmt.Println("SaveData")
    file, err := os.Create("data.dat")
    defer file.Close()
    if err != nil {
        fmt.Println(err)
    }
    enc := gob.NewEncoder(file)
    err2 := enc.Encode(SJ)
    if err2 != nil {
        fmt.Println(err2)
    }
}

var SJ ShuJu
var runing bool
var chEnd chan bool

func main() {
    LoadData()
    chEnd = make(chan bool)
    var emails, body, msgbox *walk.TextEdit
    var user, password, host, subject *walk.LineEdit
    var startBtn *walk.PushButton
    MainWindow{
        Title:   "邮件群发器 By 一曲忧伤",
        MinSize: Size{800, 600},
        Layout:  HBox{},
        Children: []Widget{
            TextEdit{AssignTo: &emails, Text: SJ.Send, ToolTipText: "待发送邮件列表，每行一个"},
            VSplitter{
                Children: []Widget{
                    LineEdit{AssignTo: &user, Text: SJ.Name, CueBanner: "请输入邮箱用户名"},
                    LineEdit{AssignTo: &password, Text: SJ.Pwd, PasswordMode: true, CueBanner: "请输入邮箱登录密码"},
                    LineEdit{AssignTo: &host, Text: SJ.Host, CueBanner: "SMTP服务器:端口"},
                    LineEdit{AssignTo: &subject, Text: SJ.Subject, CueBanner: "请输入邮件主题……"},
                    TextEdit{AssignTo: &body, Text: SJ.Body, ToolTipText: "请输入邮件内容", ColumnSpan: 2},
                    TextEdit{AssignTo: &msgbox, ReadOnly: true},
                    PushButton{
                        AssignTo: &startBtn,
                        Text:     "开始群发",
                        OnClicked: func() {
                            SJ.Name = user.Text()
                            SJ.Pwd = password.Text()
                            SJ.Host = host.Text()
                            SJ.Subject = subject.Text()
                            SJ.Body = body.Text()
                            SJ.Send = emails.Text()
                            SaveData()

                            if runing == false {
                                runing = true
                                startBtn.SetText("停止发送")
                                go sendThread(msgbox, emails)
                            } else {
                                runing = false
                                startBtn.SetText("开始群发")
                            }
                        },
                    },
                },
            },
        },
    }.Run()
}

func sendThread(msgbox, es *walk.TextEdit) {
    sendTo := strings.Split(SJ.Send, "\r\n")
    susscess := 0
    count := len(sendTo)
    for index, to := range sendTo {
        if runing == false {
            break
        }
        msgbox.SetText("发送到" + to + "..." + strconv.Itoa((index/count)*100) + "%")
        err := SendMail(SJ.Name, SJ.Pwd, SJ.Host, to, SJ.Subject, SJ.Body, "html")
        if err != nil {
            msgbox.AppendText("\r\n失败:" + err.Error() + "\r\n")
            if err.Error() == "550 Mailbox not found or access denied" {
                SJ.Send = strings.Join(DelArrayVar(strings.Split(SJ.Send, "\r\n"), to), "\r\n")
                es.SetText(SJ.Send)
            }
            time.Sleep(1 * time.Second)
            continue
        } else {
            susscess++
            msgbox.AppendText("\r\n发送成功!")
            SJ.Send = strings.Join(DelArrayVar(strings.Split(SJ.Send, "\r\n"), to), "\r\n")
            es.SetText(SJ.Send)
        }
        time.Sleep(1 * time.Second)
    }
    SaveData()
    msgbox.AppendText("停止发送!成功 " + strconv.Itoa(susscess) + " 条\r\n")
}
