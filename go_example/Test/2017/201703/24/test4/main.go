package main

import (
	"time"

	"github.com/codyguo/logs"
)

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["update_at"] = time.Now()
	}
}

func main() {
	foo := map[string]interface{}{
		"Mark": 20,
	}
	timeMap(foo)
	logs.Notice(foo)

}
