package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (this *User) Greetrings() string {
	return fmt.Sprintf("Hi %s from %s", this.Name, this.Location)
}

type Player struct {
	User
	GameId int
}

func main() {
	// p := Player{&User{1, "codyguo", "zh"}, 11}
	p := Player{User{1, "codyguo", "zh"}, 11}

	fmt.Println(p.Greetrings())
}
