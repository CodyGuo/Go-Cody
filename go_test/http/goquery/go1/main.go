package main

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".newHotB > .cc > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".s2").Text()
		value := s.Find("li > a").Text()
		// fmt.Println(value)
		fmt.Printf("Review %d: %s -> [%s]\n", i, title, value)
	})

}

func main() {
	timeNow := time.Now()
	url := "http://10.10.2.222/bbs"
	ExampleScrape(url)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("[INFO] 抓取 [%s] 用时%s.\n", url, fmt.Sprint(time.Since(timeNow)))
}
