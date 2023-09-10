package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(fn string, ln string) Instructor {
	return Instructor{FirstName: fn, LastName: ln}
}

func (i Instructor) GetDetails() {
	fmt.Printf("The instructor is a nice person named, %s %s, and they have an id of %d, as well as a score of %d\n", i.FirstName, i.LastName, i.Id, i.Score)
}
