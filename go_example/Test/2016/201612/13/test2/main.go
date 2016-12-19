package main

import (
	"fmt"
	"io"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("./ssh")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	// cmd.Stdin = os.Stdin
	w, _ := cmd.StdinPipe()
	r, _ := cmd.StdoutPipe()

	// go func() {
	// 	var (
	// 		buf [1 << 16]byte
	// 		t   int
	// 	)
	// 	for {
	// 		n, err := out.Read(buf[t:])
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			return
	// 		}
	// 		t += n
	// 		result := string(buf[:t])
	// 		fmt.Println(result)
	// 	}
	// }()

	cmd.Start()
	result := readResult(r)
	fmt.Print(result)

	w.Write([]byte("ter len 0\n"))
	result = readResult(r)
	fmt.Print(result)

	w.Write([]byte("show arp\n"))
	result = readResult(r)
	fmt.Print(result)
	// cmd.Wait()
}

func readResult(r io.Reader) string {
	var (
		t      int
		buf    [1 << 16]byte
		result string
	)

	for {
		n, _ := r.Read(buf[t:])
		t += n
		result = string(buf[:t])
		if t-1 > 0 {
			end := buf[t-1 : t]
			fmt.Printf("========== [%s] ==============\n", end)
			if end[0] == 62 {
				break
			}
			// if strings.Contains(result, ">") ||
			// 	strings.Contains(result, "#") {
			// 	result = string(buf[:t])
			// 	break
			// }
		} else if t-1 < 0 {
			break
		}

	}
	return result
}
