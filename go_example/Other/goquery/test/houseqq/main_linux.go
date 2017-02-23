// +build linux
package main

import "github.com/PuerkitoBio/goquery"

func newDocument() *goquery.Document {
	doc, err := goquery.NewDocument(URL)
	checkErr(err)

	return doc
}
