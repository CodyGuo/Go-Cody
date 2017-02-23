package main

import "fmt"

// Sender接口
type Sender interface {
	Send()
}

type MailSend struct {
}

func (this MailSend) Send() {
	fmt.Println("this is mailsend!")
}

type SmsSend struct {
}

func (this SmsSend) Send() {
	fmt.Println("this is smsSend!")
}

type qqSend struct {
}

func (this qqSend) Send() {
	fmt.Println("this is qqSend!")
}

type SendMailFactory struct {
}

func (this SendMailFactory) Produce() Sender {
	return new(MailSend)
}

type SendSmsFactory struct {
}

func (this SendSmsFactory) Produce() Sender {
	return new(SmsSend)
}

type SendQQFactory struct {
}

func (this SendQQFactory) Produce() Sender {
	return new(qqSend)
}

type Provider interface {
	Produce() Sender
}

func main() {
	var provider Provider
	provider = new(SendMailFactory)
	sender := provider.Produce()
	sender.Send()

	provider = new(SendQQFactory)
	sender = provider.Produce()
	sender.Send()
}
