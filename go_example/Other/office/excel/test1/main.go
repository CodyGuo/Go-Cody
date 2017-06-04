package main

import (
	"fmt"
	"log"

	"github.com/extrame/xls"
)

func main() {
	excelFileName := "./test.xls"
	xlFile, err := xls.Open(excelFileName, "utf-8")
	if err != nil {
		log.Fatal("openfile ----->", err)
	}
	fmt.Println(xlFile.Author)
	for i := 0; i < xlFile.NumSheets(); i++ {
		sheet := xlFile.GetSheet(i)
		fmt.Println(sheet.Name)
	}

	if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
		fmt.Print("Total Lines ", sheet1.MaxRow, sheet1.Name)
		col1 := sheet1.Row(0).Col(0)
		col2 := sheet1.Row(0).Col(0)
		for i := 0; i <= (int(sheet1.MaxRow)); i++ {
			row1 := sheet1.Row(i)
			col1 = row1.Col(0)
			col2 = row1.Col(1)
			fmt.Print("\n", col1, ",", col2)
		}
	}
}
