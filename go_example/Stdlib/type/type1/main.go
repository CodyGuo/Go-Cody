// 类型断言
package main

import (
    "fmt"
)

func main() {
    var num string

    switch interface{}(num).(type) {
    case int:
        fmt.Println("num is int.")
    case int8:
        fmt.Println("num is int8.")
    case string:
        fmt.Println("num is string.")
    default:
        fmt.Println("num is other type.")
    }

    // fmt.Println(interface{}(num).(int8))
}
