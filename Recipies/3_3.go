package main
import (
	"fmt"
)

type Student struct {
	Name string
	Age int 
	Subject[]
}

type Subject struct {
	SName string
	SChapters int
}

var stud = []Student {
			Student {
				Name :"Anupam",
				Age :23,
				 []Subject {
					Subject {
						SName:"Physics",
						SChapters:12,
					},
					Subject {
						SName :"Chemistry",
						SChapters:10,
					},
				},
			},
			Student {
				Name:"Deepoo",
				Age:24,
				[]Subject {
					Subject {
						SName : "Maths",
						SChapters:14,
					},
					Subject {
						SName:"Computers",
						SChapters:15,
					},
				},
			},
	}

type PersonName struct {
	Name string
}

var c = PersonName {
		Name:"Anupam",
}

func main () {
//	fmt.Println (stud)
	fmt.Println((&stud[1]).printName(1))
	fmt.Println (stud[1].Subject[0].printName(1))
	fmt.Println ((&c).printName())
}

func (s Student) printName (x int) (name string) {
	name = s.Name
	fmt.Println (name)
	return
}

func (s *Subject) printName (x int) (name string) {
	name = s.SName
	fmt.Println (name)
	return
}

func (p *PersonName) printName () (name string) {
	name = p.Name
	return
}
