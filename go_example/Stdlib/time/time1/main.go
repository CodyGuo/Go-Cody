package main

import (
    "log"
    "os"
    "time"
)

func main() {
    timer := time.NewTicker(2 * time.Second)
    count := 0
    for {
        select {
        case <-timer.C:
            go func() {
                if count == 5 {
                    os.Exit(0)
                }
                log.Printf("定时器运行 %d 次.\n", count+1)
                log.Println(time.Now())
                count += 1
            }()
        }
    }
}
