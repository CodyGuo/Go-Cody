package main

import (
	"testing"
)

var str = `&lt;?xml version="1.0" encoding="utf-8"?>

&lt;s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope">
  &lt;s:Header>
    &lt;wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
      <wsse:UsernameToken>
        <wsse:Username>{{.Username}}</wsse:Username>
        &lt;wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">{{.Password}}</wsse:Password>
        &lt;wsse:Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">{{.Nonce64}}</wsse:Nonce>
       &lt;wsu:Created>{{.Created}}</wsu:Created>
      &lt;/wsse:UsernameToken>
    &lt;/wsse:Security>
  &lt;/s:Header>
  &lt;s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    &lt;{{.SoapReuest}} xmlns="http://www.onvif.org/ver10/device/wsdl"></{{.SoapReuest}}>
  &lt;/s:Body>
</s:Envelope>`

func BenchmarkUnescapString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unescapeString1(str)
	}
}

func BenchmarkUnescapString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unescapeString2(str)
	}
}
