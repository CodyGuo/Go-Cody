package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL = "http://zz.house.qq.com/xfbj.htm"
)

func queryQQHouse() {
	doc := newDocument()
	doc.Find(".box").Each(func(i int, s *goquery.Selection) {
		areas := s.Find("strong")
		area := areas.Text()
		fmt.Printf("area --> %s: \n", area)
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			name := s.Find("a").Text()
			name = strings.TrimSpace(name)
			price := s.Find(".price").Text()
			price = strings.TrimSpace(price)
			direction := s.Find(".direction").Text()
			direction = strings.TrimSpace(direction)
			fmt.Printf("\tname --> %s | price --> %s | direction --> %s \n", name, price, direction)
		})
	})
}

func main() {
	queryQQHouse()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
