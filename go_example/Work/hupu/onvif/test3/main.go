package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	Created := "2016-11-25T03:44:37Z"
	Nonce := "2145383247"
	Password := "admin2016"
	// Nonce = "H25iU9gudByWLzvCfi/7Fg=="
	// Created = "2016-11-25T02:03:39Z"
	Nonce = base64.StdEncoding.EncodeToString([]byte(Nonce))
	fmt.Println(Nonce)
	// passwdSha1 := sha1.Sum([]byte(Nonce + Created + Password))
	passwdSha1 := sha1Encryption(fmt.Sprintf("%s%s%s", "2145383247", Created, Password))
	fmt.Printf("%s\n", passwdSha1)
	// passwdSha1 := []byte(fmt.Sprintf("%s%s%s", Nonce, Created, Password))

	Password = base64.StdEncoding.EncodeToString(passwdSha1)
	fmt.Println(Password)

}

func sha1Encryption(str string) []byte {
	sha := sha1.New()
	io.WriteString(sha, str)
	// sha.Write([]byte(str))
	return sha.Sum(nil)
}
