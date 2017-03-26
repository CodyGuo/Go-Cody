package main

import (
	"github.com/codyguo/logs"
)

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelase(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "CodyGuo", Genre: "hupu", Songs: 17}
	logs.Noticef("%s relase their %dth song", me.Name, newRelase(me))
	logs.Noticef("%s has a total of %d songs.", me.Name, me.Songs)
}
