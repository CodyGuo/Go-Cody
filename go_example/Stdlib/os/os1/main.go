package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.DevNull)
    fmt.Println(os.ModeAppend)
    fmt.Println(os.ModeCharDevice)
    fmt.Println(os.ModeDevice)
    fmt.Println(os.ModeDir)
    fmt.Println(os.ModeExclusive)
    fmt.Println(os.ModeNamedPipe)
    fmt.Println(os.ModePerm)
    fmt.Println(os.ModeSetgid)
    fmt.Println(os.ModeSetuid)
    fmt.Println(os.ModeSocket)
    fmt.Println(os.ModeSticky)
    fmt.Println(os.ModeSymlink)
    fmt.Println(os.ModeTemporary)
    fmt.Println(os.ModeType)
    fmt.Println(os.O_APPEND)
    fmt.Println(os.O_CREATE)
    fmt.Println(string(os.PathListSeparator))
    fmt.Println(string(os.PathSeparator))
    fmt.Println(os.SEEK_CUR)
}
