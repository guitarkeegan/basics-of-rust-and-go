package data

import "fmt"

type Duration float32 // hours

type Course struct {
	Id         int
	Name       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) String() string {
	return fmt.Sprintf("The course is: %v, and the instructor is %v", c.Name, c.Instructor.FirstName)
}
