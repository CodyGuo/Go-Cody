package main

import (
	"fmt"
	"os"

	"strconv"

	"strings"

	"github.com/Luxurioust/excelize"
)

type Employee struct {
	Name       string
	Department string
	JobNumber  int
	PunchCard  []map[string]string
}

type CheckWork struct {
	SheetName string
	Employees []Employee
}

func main() {
	xlsx, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get all the rows in a sheet.
	sheetName := xlsx.GetSheetName(2)
	fmt.Printf("sheetName --> %s\n", sheetName)

	rows := xlsx.GetRows("Sheet2")
	fmt.Printf("rows --> %d\n", len(rows))
	// var checkWork = new(CheckWork)
	// checkWork.SheetName = sheetName

	employee := new(Employee)
	lenInfo := 0
	for n, row := range rows {
		if n < 3 {
			continue
		}
		if lenInfo == 3 {
			fmt.Printf("n = %d, lenInfo =%d, employee --> %v\n", n, lenInfo, employee)
			employee = new(Employee)
			lenInfo = 0
		}
		lenInfo++
		employee = employeePrint(lenInfo, row, employee)
	}
	// fmt.Printf("[employee] --> %v\n", employee)

}

func employeePrint(lenInfo int, row []string, employee *Employee) *Employee {
	// fmt.Printf("lenInfo --> %d\n", lenInfo)
	switch lenInfo {
	case 1:
		employee.PunchCard = make([]map[string]string, len(row)+1)
	case 2:
		for nn, colCell := range row {
			// fmt.Printf("nn = %d, colCell --> %s\n", nn, colCell)
			switch nn {
			case 2:
				jobNum, _ := strconv.Atoi(colCell)
				employee.JobNumber = jobNum
			case 10:
				employee.Name = colCell
			case 20:
				employee.Department = colCell
			default:
				continue
			}
		}
	case 3:
		for nn, colCell := range row {
			// fmt.Printf("3 nn = %d, colCell --> %s\n", nn, colCell)

			col := strings.TrimSpace(colCell)
			if col != "" {
				employee.PunchCard[nn] = map[string]string{col: col}
			}
		}
	}
	return employee
}
