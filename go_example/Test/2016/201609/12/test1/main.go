package main

import (
	"fmt"
)

type People struct {
	Name string
	Sex  string
}

type Teacher struct {
	People
	Number int
}

func (t Teacher) String() string {
	return fmt.Sprintf("Teacher Name: %s Sex: %s Number: %d\n", t.Name, t.Sex, t.Number)
}

func (p People) String() string {
	return fmt.Sprintf("People Name: %s Sex: %s\n", p.Name, p.Sex)
}

func main() {
	var p People = People{
		Name: "codyguo",
		Sex:  "boy",
	}

	var t Teacher = Teacher{People: People{Name: "alice", Sex: "girl"},
		Number: 1,
	}

	fmt.Println(checker(t))

	fmt.Println(p)
	fmt.Println(t)

}

func checker(check interface{}) bool {
	_, ok := check.(fmt.Stringer)

	return ok
}
