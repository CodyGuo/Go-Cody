// UTF-8 字符转码 为 GBK
package main

import (
    "bytes"
    "fmt"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
)

func main() {
    str := "缂栫▼缁冧範棰樼粷瀵圭粡鍏"
    dst := Decode(str)
    fmt.Println(dst)
}

func Encode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
    if err == nil {
        dst = string(data)
    }
    return
}
func Decode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
    if err == nil {
        dst = string(data)
    }
    return
}
