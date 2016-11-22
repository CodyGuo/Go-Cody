package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	str := `<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope"><s:Header><wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"><wsse:UsernameToken><wsse:Username>admin</wsse:Username><wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">vNfYT8i+QjRG6ytv4QpNxc0pvCk=</wsse:Password><wsse:Nonce>34sJ+Xd6Yk03VZz/LysSTg==</wsse:Nonce><wsu:Created>2016-11-22T10:04:31Z</wsu:Created></wsse:UsernameToken></wsse:Security></s:Header><s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><GetDeviceInformation xmlns="http://www.onvif.org/ver10/device/wsdl"/></s:Body></s:Envelope>`
	body := bytes.NewBuffer([]byte(str))
	resp, err := http.Post("http://10.10.2.15/onvif/device_service", "application/soap+xml;charset=utf-8", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}
