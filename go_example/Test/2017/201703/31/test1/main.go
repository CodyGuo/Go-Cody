package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	sx := fmt.Sprintf("%x", "abc")
	fmt.Println(sx)
	bx, err := hex.DecodeString(sx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bx)
}
