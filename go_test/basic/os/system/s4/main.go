package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func logErr(err interface{}) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	cmd := exec.Command("ls", "-l")
	stdout, _ := cmd.StdoutPipe()

	logErr(cmd.Start())

	bs, _ := ioutil.ReadAll(stdout)
	fmt.Println(string(bs))

	logErr(cmd.Wait())

	bs2, err := exec.Command("ls").Output()
	logErr(err)

	fmt.Println(string(bs2))
}
