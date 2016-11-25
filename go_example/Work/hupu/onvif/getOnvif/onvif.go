package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

var deviceInfo Envelope

type Onvif struct {
	IP       string
	Username string
	Password string
	Created  string
	Nonce64  string

	body io.ReadWriter
}

func NewOnvif() *Onvif {
	return new(Onvif)
}

func (onv *Onvif) OnvifDevice() (string, error) {
	layout := "2006-01-02T15:04:05Z"
	onv.Created = time.Now().UTC().Format(layout)
	nonce := getNonce(24)
	onv.Nonce64 = base64.StdEncoding.EncodeToString([]byte(nonce))
	passwdSha1 := sha1Encryption(fmt.Sprintf("%s%s%s", nonce, onv.Created, onv.Password))
	onv.Password = base64.StdEncoding.EncodeToString(passwdSha1)

	onvifXML := template.New("onvif")
	onvifXML.Parse(loginXML)

	var buf bytes.Buffer
	err := onvifXML.Execute(&buf, onv)
	if err != nil {
		return "", err
	}
	loginXML := buf.String()
	loginXML = strings.Replace(loginXML, "&lt;", "<", -1)

	getDeviceXML := loginXML + deviceInfoXML
	onv.body = bytes.NewBuffer([]byte(getDeviceXML))
	if err := onv.GetDeviceInfo(); err != nil {
		return "", err
	}

	getNetworkXML := loginXML + networkInterfacesXML
	onv.body = bytes.NewBuffer([]byte(getNetworkXML))
	if err = onv.GetNetworkInterfaces(); err != nil {
		return "", err
	}

	result, err := json.MarshalIndent(deviceInfo, "", "    ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (onv *Onvif) GetDeviceInfo() error {
	data, err := onv.getRequest()
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, &deviceInfo)
	if err != nil {
		return err
	}
	return nil
}

func (onv *Onvif) GetNetworkInterfaces() error {
	data, err := onv.getRequest()
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, &deviceInfo)
	if err != nil {
		return err
	}
	return nil
}

func (onv *Onvif) getRequest() ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*3)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 3))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	resp, err := client.Post("http://"+onv.IP+"/onvif/device_service", "application/soap+xml;charset=utf-8", onv.body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(data), "NotAuthorized") {
		return nil, fmt.Errorf("getRequest IP:%s %s", onv.IP, "NotAuthorized")
	}

	return data, nil
}

func sha1Encryption(str string) []byte {
	sha := sha1.New()
	io.WriteString(sha, str)
	return sha.Sum(nil)
}

func getNonce(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=-"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
