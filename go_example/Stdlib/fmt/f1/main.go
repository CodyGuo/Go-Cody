// 每一个Scanf的格式化字符串中都必须带有 \n
package main

import (
	"fmt"
)

func main() {
	var str1 string

	fmt.Print("Please input str1: ")
	fmt.Scanf("%s\n", &str1)
	fmt.Println("Output str1: ", str1)

	fmt.Println("----------------------------")

	var str2, str3 string
	fmt.Print("Please input str2-str3: ")
	fmt.Scanf("%s %s\n", &str2, &str3)

	fmt.Println("Output str2-str3: ", str2, str3)

	fmt.Println("----------------------------")

	var str4, str5 string
	fmt.Print("Please input str4-str5: ")
	fmt.Scanf("%s %s", &str4, &str5)

	fmt.Println("Output str4-str5: ", str4, str5)
}
