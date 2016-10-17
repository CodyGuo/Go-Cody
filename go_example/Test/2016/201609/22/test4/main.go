package main

import (
	"flag"
	"fmt"
)

var (
	user   string
	passwd string
)

func init() {
	flag.StringVar(&user, "user", "root", "set user name.")
	flag.StringVar(&passwd, "passwd", "123456", "set passwd.")
}

func main() {
	flag.Parse()
	fmt.Println(user, passwd)
	fmt.Printf("first arg: [%s]\n", flag.Arg(0))
	fmt.Println(flag.NFlag())
}
