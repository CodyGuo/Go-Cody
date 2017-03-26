package main

import (
	"github.com/codyguo/logs"
)

type Bootcamp struct {
	Lat, Lon float64
}

func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	logs.Notice(*x == *y)
}
