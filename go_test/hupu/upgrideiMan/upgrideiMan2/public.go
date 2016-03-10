package main

import (
	"log"
	"regexp"
	"strings"
)

const (
	line        = 10    // '\n'
	enter       = 13    //'\r'
	space       = 32    // ' '
	commaEN     = 44    // ','
	commaZH     = 65292 // '，'
	semicolonEN = 59    // ';'
	semicolonZH = 65307 // '；'

	width      = 450
	height     = 380
	lineHeight = 25
)

const (
	ipRegxp = `^((25[0-5]|2[0-4]\d|1\d{2}|\d?\d)\.){3}(25[0-5]|2[0-4]\d|1\d{2}|\d?\d)$`
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func regex(s string) bool {
	re, err := regexp.Compile(ipRegxp)
	checkError(err)
	return re.MatchString(s)
}

// 转换为list
func stringToList(sip string) []string {
	m := func(c rune) rune {
		switch c {
		case line, enter:
			return commaEN
		}
		return c
	}
	sip = strings.Map(m, sip)

	f := func(c rune) bool {
		switch c {
		case space, commaEN, commaZH, semicolonEN, semicolonZH:
			return true
		}
		return false
	}
	return strings.FieldsFunc(sip, f)
}

// IP地址验证
func checkSIP(sipList []string) ([]string, []string) {
	var okSIP, errSIP []string
	if len(sipList) != 0 {
		for _, sip := range sipList {
			if regex(sip) {
				okSIP = append(okSIP, sip)
			} else {
				errSIP = append(errSIP, sip)
			}
		}
		return okSIP, errSIP
	}
	return nil, nil
}
