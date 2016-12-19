package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	src := string(body)
	// fmt.Println(src)
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	// fmt.Println(src)
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	// fmt.Println(src)
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src, "")
	fmt.Println(src)
}
