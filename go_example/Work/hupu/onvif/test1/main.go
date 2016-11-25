package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	passwd = "vNfYT8i+QjRG6ytv4QpNxc0pvCk="
)

var (
	IP       string
	Username string
	Password string
	Created  string
	Nonce    string
)

var (
	loginXML             string
	deviceInfoXML        string
	networkInterfacesXML string
)

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=-"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

type Envelope struct {
	Body Body `xml:"Body"`
}
type Body struct {
	DeviceInformation            DeviceInformation `xml:"GetDeviceInformationResponse"`
	GetNetworkInterfacesResponse GetNetworkInterfacesResponse
}

type DeviceInformation struct {
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
}

type GetNetworkInterfacesResponse struct {
	NetworkInterfaces NetworkInterfaces
}

type NetworkInterfaces struct {
	Info Info
}

type Info struct {
	Name      string
	HwAddress string
}

var deviceInfo Envelope

func main() {
	Username = "admin"
	Password = ""
	IP = "10.10.2.15"

	layout := "2006-01-02T15:04:05Z"
	Created = time.Now().UTC().Format(layout)
	Nonce = GetRandomString(24)
	Nonce64 := base64.StdEncoding.EncodeToString([]byte(Nonce))
	passwdSha1 := sha1Encryption(fmt.Sprintf("%s%s%s", Nonce, Created, Password))
	Password = base64.StdEncoding.EncodeToString(passwdSha1)

	loginXML = `<?xml version="1.0" encoding="utf-8"?>

<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope">
  <s:Header>
    <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
      <wsse:UsernameToken>
        <wsse:Username>` + Username + `</wsse:Username>
        <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">` + Password + `</wsse:Password>
        <wsse:Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">` + Nonce64 + `</wsse:Nonce>
        <wsu:Created>` + Created + `</wsu:Created>
      </wsse:UsernameToken>
    </wsse:Security>
  </s:Header>
  `
	deviceInfoXML = `<s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <GetDeviceInformation xmlns="http://www.onvif.org/ver10/device/wsdl"></GetDeviceInformation>
  </s:Body>
</s:Envelope>`

	networkInterfacesXML = `<s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <GetNetworkInterfaces xmlns="http://www.onvif.org/ver10/device/wsdl"></GetNetworkInterfaces>
  </s:Body>
</s:Envelope>`
	loginDeviceXML := loginXML + deviceInfoXML
	body := bytes.NewBuffer([]byte(loginDeviceXML))
	GetDeviceInfo(IP, body)

	networkXML := loginXML + networkInterfacesXML
	body = bytes.NewBuffer([]byte(networkXML))
	GetNetworkInterfaces(IP, body)

	result, _ := json.MarshalIndent(deviceInfo, "", "    ")
	fmt.Println(string(result))
}

func GetDeviceInfo(ip string, body io.ReadWriter) {
	resp, err := http.Post("http://"+ip+"/onvif/device_service", "application/soap+xml;charset=utf-8", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = xml.Unmarshal(data, &deviceInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func GetNetworkInterfaces(ip string, body io.ReadWriter) {
	resp, err := http.Post("http://"+ip+"/onvif/device_service", "application/soap+xml;charset=utf-8", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = xml.Unmarshal(data, &deviceInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func sha1Encryption(str string) []byte {
	sha := sha1.New()
	io.WriteString(sha, str)
	return sha.Sum(nil)
}
