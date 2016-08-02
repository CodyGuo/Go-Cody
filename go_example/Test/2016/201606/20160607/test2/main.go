package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	var inputNum int
	fmt.Print("请输入数字N：")
	fmt.Scan(&inputNum)

	sumLen, err := maxNumPow(inputNum)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < sumLen; i++ {
		fmt.Printf("%d", i)
		if (i + 1) != sumLen {
			fmt.Print(", ")
		}
	}
}

func maxNumAdd(value int) {
	maxNum := math.Pow10(value)
	sumLen := int(maxNum)
	for i := 0; ; i++ {
		if i > sumLen {
			break
		}
		// fmt.Printf("%d", i)
		if (i + 1) != sumLen {
			// fmt.Print(", ")
		}
	}
}

func maxNumStr(value int) (int, error) {
	if value <= 0 {
		return 0, nil
	}

	start := "1"
	for i := 0; i < value; i++ {
		start += fmt.Sprintf("%d", 0)
	}

	sumLen, err := strconv.Atoi(start)
	for i := 0; i < sumLen; i++ {
		// fmt.Printf("%d", i)
		if (i + 1) != sumLen {
			// fmt.Print(", ")
		}
	}
	return sumLen, err
}

func maxNumPow(value int) (int, error) {
	tmp := math.Pow10(value)
	sumLen := int(tmp)
	// sumLen := int(math.Pow10(value))
	for i := 0; i < sumLen; i++ {
		// fmt.Printf("%d", i)
		if (i + 1) != sumLen {
			// fmt.Print(", ")
		}
	}
	return sumLen, nil
}
