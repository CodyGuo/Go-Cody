package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f map[string]interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// m := f.(map[string]interface{})
	for k, v := range f {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "Is string", vv)
		case int:
			fmt.Println(k, "Is int", vv)
		case float64:
			fmt.Println(k, "Is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type i dont know now a nanare.")
			fmt.Printf("%#T\n", vv)
		}
	}
}
