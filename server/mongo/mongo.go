package mongo

import (
	"fmt"
)

type speaker interface {
	print()
}
type Father struct {
	age int
}

func (f *Father) print() {
	fmt.Printf("Im father and age:%d\n", f.age)
}

type Son struct {
	Father
	job string
}

func (s *Son) print() {
	fmt.Printf("Im Son and age:%d and job:%s", s.age, s.job)
}

func say(s speaker) {
	s.print()
}
func main() {
	s := &Son{
		Father{
			age: 10,
		},
		"coding",
	}
	say(s)
}
