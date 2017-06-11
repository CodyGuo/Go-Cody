package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	hostsURL = "https://raw.githubusercontent.com/racaljk/hosts/master/hosts"
)

func downloadHosts(url string) error {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	hosts, err := os.Create("hosts")
	defer hosts.Close()
	if err != nil {
		return err
	}

	io.Copy(hosts, resp.Body)
	return nil
}

func getLastUpdate(name string) []byte {
	file, err := os.Open(name)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lastUpate []byte
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil && err != io.EOF {
			return nil
		}
		if err == io.EOF {
			break
		}
		if bytes.HasPrefix(line, []byte("# Last updated:")) {
			lastUpate = bytes.Split(line, []byte(": "))[1]
			break
		}
	}
	return lastUpate
}

func main() {
	downloadHosts(hostsURL)
	lastUpate := getLastUpdate("hosts")
	fmt.Printf("lastUpate -->{%s}\n", lastUpate)
}
