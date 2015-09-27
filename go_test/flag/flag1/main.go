package main

import (
    "flag"
    "fmt"
    "os"
)

var (
    flagOne string
    debug   bool
)

func init() {
    flag.StringVar(&flagOne, "one", "", "The string example one.")
    flag.BoolVar(&debug, "d", debug, "The bool example debug.")
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of cody.guo error %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

}

func main() {
    flag.Parse()

    if flagOne == "" {
        fmt.Fprintf(os.Stdout, "Usage of cody.guo ok %s:\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(1)
    }

    fmt.Println("参数数量:", flag.NFlag())
    oneFlag := flag.Lookup("one")

    fmt.Println(oneFlag.Name, oneFlag.Value)

    // fmt.Println(len(os.Args))

    if debug {
        fmt.Println("debug is on.")
    } else {
        fmt.Println("debug is off.")
    }

    fmt.Println(flagOne)

    debugFlag := flag.Lookup("d")
    fmt.Println(debugFlag.Name, debugFlag.Value)

}
