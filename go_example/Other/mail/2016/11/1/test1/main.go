package main

import (
	// "encoding/base64"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/scorredoira/email"
)

func main() {
	// compose the message
	if len(os.Args) < 5 {
		os.Stdout.WriteString("帮助说明:\n\tsendMail.exe hupu@hupu.net passwd smtp.hupu.net 25")
		return
	}

	fromTo := os.Args[1]
	m := email.NewMessage("Hi", "this is iMan test email.\n监控平台邮件服务器测试成功.")
	m.From = mail.Address{Name: "监控平台测试邮件", Address: fromTo}
	m.To = []string{fromTo}

	// // add attachments
	// if err := m.Attach("main.go"); err != nil {
	// 	log.Fatal(err)
	// }

	// send it
	passwd := os.Args[2]
	smtpConfig := os.Args[3]
	port := os.Args[4]

	// username := base64.StdEncoding.EncodeToString([]byte(fromTo))
	// passwd = base64.StdEncoding.EncodeToString([]byte(passwd))
	server := smtp.ServerInfo{TLS: false}
	auth := smtp.PlainAuth("LOGIN", fromTo, passwd, smtpConfig)
	auth.Start(&server)
	if err := email.Send(smtpConfig+":"+port, auth, m); err != nil {
		log.Fatal(err)
	}
	os.Stdout.WriteString("测试邮件发送成功!\n")
	for {
	}
}
