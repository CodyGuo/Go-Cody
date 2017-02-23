package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "./test.xls"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal("openfile ----->", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
