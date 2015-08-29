package main

import (
    "fmt"
    "strconv"
)

/* 大白加大白等于白胖胖.求大,白,胖各是数字多少？ */
func main() {
    for d := 0; d < 10; d++ {
        for b := 0; b < 10; b++ {
            for p := 0; p < 10; p++ {
                daBai, _ := strconv.Atoi(strconv.Itoa(d) + strconv.Itoa(b))
                // fmt.Println("大白", daBai)
                baiPangPang, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(p) + strconv.Itoa(p))
                // fmt.Println("白胖胖", baiPangPang)
                if daBai+daBai == baiPangPang && d != 0 && b != 0 && p != 0 {
                    fmt.Println("-------------------大 白 胖--------------------")
                    fmt.Printf("大 = %d, 白 = %d, 胖 = %d\n", d, b, p)
                    fmt.Println("白胖胖: ", baiPangPang)
                }
            }
        }
    }
}
