package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	INFO  = "[INFO ] "
	ERROR = "[ERROR] "
)

func main() {
	flag.Parse()

	pack, err := NewPack()
	checkError(err)

	if err := pack.Pack(); err != nil {
		fmt.Printf("%sPack helper failed, %s\n", ERROR, err)
		os.Exit(1)
	} else {
		fmt.Printf("%sPack helper success, ready to upload helper to compile the Linux server.\n", INFO)
	}

	if err := pack.Upload(); err != nil {
		fmt.Printf("%sHelper to upload failed, %s\n", ERROR, err)
		os.Exit(1)
	} else {
		fmt.Printf("%sHelper to upload success.", INFO)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runCmd(cmd string) error {
	lists := strings.Split(cmd, " ")

	return exec.Command(lists[0], lists[1:]...).Run()
}
