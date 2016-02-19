package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	// "time"
)

type proxy struct {
	country   string
	ip        string
	port      string
	position  string
	anonymous string
	pTtype    string
	time      string
}

func findProxy(url string) {
	doc, err := goquery.NewDocument(url)
	checkError(err)

	title := new(proxy)
	var content = make([]proxy, 5, 10)

	doc.Find("#ip_list").Each(func(i int, t *goquery.Selection) {
		t.Find(".subtitle").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				title.country = s.Find("th").Eq(0).Text()
				title.ip = s.Find("th").Eq(1).Text()
				title.port = s.Find("th").Eq(2).Text()
				title.position = s.Find("th").Eq(3).Text()
				title.anonymous = s.Find("th").Eq(4).Text()
				title.pTtype = s.Find("th").Eq(5).Text()
				title.time = s.Find("th").Eq(6).Text()
			}
		})

		t.Find("tbody").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(n int, s *goquery.Selection) {
				if !s.Find("tr").Is(".subtitle") || i != 0 {
					coun, _ := s.Find("td > img").Attr("alt")
					content = append(content, proxy{
						coun,
						s.Find("td").Eq(1).Text(),
						s.Find("td").Eq(2).Text(),
						s.Find("td").Eq(3).Text(),
						s.Find("td").Eq(4).Text(),
						s.Find("td").Eq(5).Text(),
						s.Find("td").Eq(6).Text()})
				}
			})
		})
	})

	fmt.Printf("%-4s %-16s %-6s %-12s %-6s %-10s \n", title.country, title.ip, title.port,
		title.position, title.pTtype, title.time)

	for _, value := range content {
		if value.country == "Cn" {
			fmt.Printf("%3s %16s %10s %10s %+10s %+10s \n", value.country, value.ip, value.port,
				value.position, value.pTtype, value.time)
		}
	}

}

func main() {
	url := "http://www.xicidaili.com/"
	findProxy(url)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
