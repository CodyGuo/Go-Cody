package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"strings"
)

var (
	soapXML = `<?xml version="1.0" encoding="utf-8"?>

<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope">
  <s:Header>
    <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
      <wsse:UsernameToken>
        <wsse:Username>{{.Username}}</wsse:Username>
        <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">{{.Password}}</wsse:Password>
        <wsse:Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">{{.Nonce64}}</wsse:Nonce>
        <wsu:Created>{{.Created}}</wsu:Created>
      </wsse:UsernameToken>
    </wsse:Security>
  </s:Header>
  <s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <{{.SoapReuest}} xmlns="http://www.onvif.org/ver10/device/wsdl"></{{.SoapReuest}}>
  </s:Body>
</s:Envelope>`
)

type Onvif struct {
	Username   string
	Password   string
	Created    string
	Nonce64    string
	SoapReuest string
}

func main() {
	onvifXML := template.New("onvif")
	onvifXML.Parse(soapXML)
	onv := new(Onvif)
	onv.Username = "admin"
	onv.Password = "admin2016"
	onv.Created = "2016"
	onv.Nonce64 = "123456"
	onv.SoapReuest = "hello"

	var buf bytes.Buffer
	if err := onvifXML.Execute(&buf, onv); err != nil {
		fmt.Println(err)
		return
	}
	soapXML := buf.String()
	fmt.Println(soapXML)

	fmt.Println(html.UnescapeString(soapXML))
}

func unescapeString1(str string) string {

	return html.UnescapeString(str)
}

func unescapeString2(str string) string {
	return strings.Replace(str, "&lt;", "<", -1)
}
