package main

import (
	"fmt"

	"guitarkeegan.com/go/data"
	d "guitarkeegan.com/go/data"
)

func main() {
	instructor := d.Instructor{Id: 110, LastName: "Mine"}
	instructor.FirstName = "Joe"

	instructor.GetDetails()

	me := d.NewInstructor("G", "K")

	fmt.Printf("%v", me)

	goClass := data.Course{Id: 100, Name: "Go Class", Instructor: me}

	about := goClass.String()

	fmt.Printf("%v", about)

}
