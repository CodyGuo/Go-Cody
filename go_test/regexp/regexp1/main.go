package main

import (
	"fmt"
	"regexp"
)

var digitsRegexp = regexp.MustCompile(`(\d+)\D+(\d+)`)

func main() {
	someString := "1000abcd123"
	fmt.Println(digitsRegexp.FindStringSubmatch(someString))

	str := `
apnic|TW|ipv4|140.133.0.0|65536|19921229|allocated
apnic|TW|ipv4|140.134.0.0|65536|19920902|allocated
apnic|TW|ipv4|140.135.0.0|65536|19900324|allocated
apnic|CK|ipv4|140.136.0.0|131072|19900324|allocated
apnic|WN|ipv4|140.138.0.0|65536|19921229|allocated
apnic|CN|ipv4|140.143.0.0|65536|20110310|allocated
apnic|WN|ipv4|140.149.0.0|65536|20110311|allocated
apnic|AU|ipv4|140.159.0.0|65536|19900502|allocated
apnic|CK|ipv4|140.168.0.0|65536|19900508|allocated
apnic|NZ|ipv4|140.200.0.0|65536|19900611|allocated
apnic|CN|ipv4|140.205.0.0|65536|20110311|allocated
apnic|CN|ipv4|140.206.0.0|131072|20110309|allocated
apnic|TH|ipv4|140.210.0.0|65536|20110314|allocated
apnic|CK|ipv4|140.210.0.0|65536|20110314|allocated
apnic|WN|ipv4|140.210.0.0|65536|20110314|allocated
apnic|AK|ipv4|140.210.0.0|65536|20110314|allocated
apnic|TH|ipv4|140.210.0.0|65536|20110314|allocated
apnic|AK|ipv4|140.210.0.0|65536|20110314|allocated`

	re, _ := regexp.Compile(`(apnic\|([TA][WHK]|CN)\|ipv4)\|([0-9|\.]{1,15})\|(\d+)\|(\d+)\|([a-z]+)`)
	result := re.FindAllString(str, -1)
	for _, s := range result {
		fmt.Println(s)
	}
}
