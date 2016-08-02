package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main() {
	var name string
	for {
		fmt.Print("请输入姓名：")
		fmt.Scanf("%s \r\n", &name)
		if name != "" {
			fmt.Printf("你好！%s\n", name)
		}
		// os.Stdin.Truncate(int64(0))
	}
	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Enter Username: ")
	// 	username, _ := reader.ReadString('\n')
	// 	fmt.Println("the username: ", username)
	// }

}
