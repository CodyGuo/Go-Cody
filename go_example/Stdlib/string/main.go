package main

import (
    "fmt"
)

func main() {
    var str string     // 声明一个字符串
    str = "cody.guo"   // 赋值
    ch := str[0]       // 获取第一个字符
    lenStr := len(str) // 字符串的长度, len是内置函数,lenStr = 8

    fmt.Println(str, ch, lenStr)
    /*
       cody.guo 99 8
    */

    fmt.Println("----------------------------------")
    var str2 string
    str2 = "cody.guo郭肖"

    for i := 0; i < len(str2); i++ {
        fmt.Println(str2[i])
    }

    for i, s := range str2 {
        fmt.Println(i, "Unicode(", s, ") string=", string(s))
    }

    r := []rune(str2)
    fmt.Println("rune=", r)

    for i := 0; i < len(r); i++ {
        fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
    }

    fmt.Println("---------------------------")
    str3 := "cody郭"
    str4 := "codyguo"
    fmt.Println("len(", str3, ")=", len(str3))
    fmt.Println("len(", str4, ")=", len(str4))
    fmt.Println("str3[0]=", string(str[0]))

}
