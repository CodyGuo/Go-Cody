// +build windows
package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

func newDocument() *goquery.Document {
	resp, err := http.Get(URL)
	checkErr(err)
	defer resp.Body.Close()
	dec := mahonia.NewDecoder("gbk")
	rd := dec.NewReader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(rd)
	checkErr(err)

	return doc
}
