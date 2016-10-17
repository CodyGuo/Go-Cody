package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
import (
	"github.com/CodyGuo/Go-Cody/go_example/Test/2016/201610/15/test2/mlib"
	"github.com/CodyGuo/Go-Cody/go_example/Test/2016/201610/15/test2/mp"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	if len(tokens) < 2 {
		fmt.Println("USAGE: lib <tool>")
		return
	}

	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id),
				tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			index, _ := strconv.Atoi(tokens[2])
			lib.Remove(index)
		} else {
			fmt.Println("USAGE: lib remove <index>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func hanlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			hanlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
