package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("main_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var out bytes.Buffer
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s := sc.Text()

		if strings.Contains(s, "github.com/lxn") {
			fmt.Printf("reflase -> %s\n", s)
			s = strings.Replace(s, "github.com/lxn", "github.com/codyguo", -1)
		}

		fmt.Fprintln(&out, s)
	}

	w, _ := os.Create("test.go")
	w.Write(out.Bytes())
	w.Sync()

	w.Close()
}
