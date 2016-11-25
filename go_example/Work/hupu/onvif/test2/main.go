package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// 	xmlns = `<?xml version="1.0" encoding="utf-8"?>
	// <s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope">
	//   <s:Header>
	//     <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	//       <wsse:UsernameToken>
	//         <wsse:Username>admin</wsse:Username>
	//         <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">vNfYT8i+QjRG6ytv4QpNxc0pvCk=</wsse:Password>
	//         <wsse:Nonce>34sJ+Xd6Yk03VZz/LysSTg==</wsse:Nonce>
	//         <wsu:Created>2016-11-22T10:04:31Z</wsu:Created>
	//       </wsse:UsernameToken>
	//     </wsse:Security>
	//   </s:Header>
	//   <s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	//     <GetDeviceInformation xmlns="http://www.onvif.org/ver10/device/wsdl"></GetDeviceInformation>
	//   </s:Body>
	// </s:Envelope>`

	xmlns = `<?xml version="1.0" encoding="utf-8"?>

<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope">
  <s:Header>
    <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
      <wsse:UsernameToken>
        <wsse:Username>admin</wsse:Username>
        <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">VRunYJdH2rEnJfCVia80NYSYP7E=</wsse:Password>
        <wsse:Nonce>rq2w8ZBjyI7uMk5ia+zqg5d1</wsse:Nonce>
        <wsu:Created>2016-11-25T01:58:19Z</wsu:Created>
      </wsse:UsernameToken>
    </wsse:Security>
  </s:Header>
  <s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <GetDeviceInformation xmlns="http://www.onvif.org/ver10/device/wsdl"></GetDeviceInformation>
  </s:Body>
</s:Envelope>`
)

func main() {
	body := bytes.NewBuffer([]byte(xmlns))
	resp, err := http.Post("http://10.10.3.16/onvif/device_service", "application/soap+xml;charset=utf-8", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}
