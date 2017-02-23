// 简单工厂模式 - 邮件和短信
package main

import "fmt"

// 公共接口
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

// 建立工厂
type SendFactory struct {
}

//　普通工厂模式
func (this SendFactory) Produce(type_ string) Sender {
	switch type_ {
	case "mail":
		return this.ProduceMail()
	case "sms":
		return this.ProduceSms()
	default:
		return nil
	}
}

// 多个工厂方法模式
func (this SendFactory) ProduceMail() Sender {
	return new(MailSend)
}

// 多个工厂方法模式
func (this *SendFactory) ProduceSms() Sender {
	return new(SmsSend)
}

func main() {
	factory := new(SendFactory)

	sender := factory.Produce("mail")
	sender.Send()

	sender = factory.Produce("sms")
	if sender != nil {
		sender.Send()
	}

	sender = factory.Produce("qq")
	if sender != nil {
		sender.Send()
	}

	mail := factory.ProduceMail()
	mail.Send()
	sms := factory.ProduceSms()
	sms.Send()

}
