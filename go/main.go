package main

import "fmt"

func main() {

	// you can also declare a variable within an if else foo := "bar"; condition
	age := 2
	// if else
	if age > 2 {
		fmt.Println("you're old")
	} else {
		fmt.Printf("You are too you because you are %d years old\n", age)
	}

	// switch will break by default, must use fallthrough to go to the next
	// condition
	switch age {
	case 1:
		{
			fmt.Printf("You're %d!\n", age)
		}
	case 2:
		{
			fmt.Printf("You're %d!\n", age)
		}
		fallthrough
	default:
		{
			fmt.Printf("And... you're %d!\n", age)
		}

	}

	// you can also use and empty switch with a bunch of boolean expressions
	name := "John"

	switch {
	case name == "John":
		fmt.Printf("Name is %s yippy!\n", name)
		fallthrough
	case age > 0:
		fmt.Println("You're alive!! Yay!")
		fallthrough
	default:
		fmt.Println("This is the end!")
	}

	// loops we have for and for in to get the index for each element in
	// the collection
	for i := 3; i > 0; i-- {
		fmt.Printf("%d\n", i)
	}
	fmt.Println("Blast off!!!")

	myStuff := []string{"foo", "bar", "baz"}

	for index, value := range myStuff {
		fmt.Printf("%d is %s\n", index, value)
	}
}
