package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(Getrand())
}

func Getrand() int {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	return source.Int()
}

func getNonce1() int {
	// str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=-"
	// bytes := []byte(str)
	// result := []byte{}
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < n; i++ {
	// 	result = append(result, bytes[r.Intn(len(bytes))])
	// }
	// return string(result)
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	return source.Int()
}

func getNonce2() string {
	// str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=-"
	// bytes := []byte(str)
	// result := []byte{}
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < n; i++ {
	// 	result = append(result, bytes[r.Intn(len(bytes))])
	// }
	// return string(result)
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := source.Int()
	return strconv.Itoa(nonce)
}

func getNonce3(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=-"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func toString1() string {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := source.Int()
	return string(nonce)
}

func toString2() string {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := source.Int()
	return strconv.Itoa((nonce))
}
