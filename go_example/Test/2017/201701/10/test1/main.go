package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

func main() {
	var va1 map[string]interface{}

	if err := json.Unmarshal([]byte(input), &va1); err != nil {
		log.Fatal(err)
	}

	fmt.Println(va1)
	for k, v := range va1 {
		fmt.Println(k, reflect.TypeOf(v))
	}
}
