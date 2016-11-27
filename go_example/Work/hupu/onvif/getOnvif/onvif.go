package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

var deviceInfo Envelope

type Onvif struct {
	IP         string
	Username   string
	Password   string
	Created    string
	Nonce64    string
	SoapReuest string
}

func NewOnvif() *Onvif {
	return new(Onvif)
}

func (onv *Onvif) OnvifDevice() (string, error) {
	onv.getEncryptedPassword()
	if err := onv.GetDeviceInfo(); err != nil {
		return "", err
	}

	if err := onv.GetNetworkInterfaces(); err != nil {
		return "", err
	}

	result, err := json.MarshalIndent(deviceInfo, "", "    ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (onv *Onvif) GetDeviceInfo() error {
	onv.SoapReuest = "GetDeviceInformation"
	data, err := onv.getRequest()
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(data, &deviceInfo); err != nil {
		return err
	}
	return nil
}

func (onv *Onvif) GetNetworkInterfaces() error {
	onv.SoapReuest = "GetNetworkInterfaces"
	data, err := onv.getRequest()
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(data, &deviceInfo); err != nil {
		return err
	}
	return nil
}

func (onv *Onvif) getEncryptedPassword() {
	nonce := createNonce()
	onv.Nonce64 = getEncryptedNonce(nonce)
	passwdSha1 := sha1Encryption(fmt.Sprintf("%s%s%s", nonce, onv.Created, onv.Password))
	onv.Password = base64.StdEncoding.EncodeToString(passwdSha1)
	onv.Created = getUTCTime()
}

func (onv *Onvif) getRequest() ([]byte, error) {
	msg, err := onv.createSoapMessage()
	if err != nil {
		return nil, err
	}

	var timeout time.Duration = 3
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*timeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * timeout))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * timeout,
		},
	}
	req, err := http.NewRequest("POST", "http://"+onv.IP+"/onvif/device_service", msg)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Hupu-iMan-Onvif")
	req.Header.Add("Content-Type", "application/soap+xml;charset=utf-8")
	// resp, err := client.Post("http://"+onv.IP+"/onvif/device_service", "application/soap+xml;charset=utf-8", onv.body)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if bytes.Contains(data, []byte("NotAuthorized")) {
		return nil, fmt.Errorf("getRequest IP:%s %s", onv.IP, "NotAuthorized")
	}

	return data, nil
}

func (onv *Onvif) createSoapMessage() (msg *bytes.Buffer, err error) {
	onvifXML := template.New("onvif")
	onvifXML.Parse(soapXML)

	var buf bytes.Buffer
	if err := onvifXML.Execute(&buf, onv); err != nil {
		return nil, err
	}
	soapXML := buf.String()
	// soapXML = strings.Replace(soapXML, "&lt;", "<", -1)
	// UnescapeString 提升速度200%
	soapXML = html.UnescapeString(soapXML)
	msg = bytes.NewBuffer([]byte(soapXML))

	return msg, nil
}

func createNonce() string {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := source.Int()

	return strconv.Itoa(nonce)
}

func getEncryptedNonce(nonce string) string {
	return base64.StdEncoding.EncodeToString([]byte(nonce))
}

func getUTCTime() string {
	layout := "2006-01-02T15:04:05Z"
	return time.Now().UTC().Format(layout)
}

func sha1Encryption(str string) []byte {
	sha := sha1.New()
	io.WriteString(sha, str)
	return sha.Sum(nil)
}
