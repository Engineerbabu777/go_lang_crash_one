package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello From NZ!!")

	// STRUCT!!
	type Student struct {
		Name       string
		Roll       int
		Percentage float32
	}

	type Department struct {
		Teachers int // TYPE WILL BE INTEGER!
		Students int  // TYPE WILL BE INTEGER!
		Courses  []string  // TYPE WILL BE SLICES OF STRING!
	}

	// CREATING AN INSTANCE OF STUDENT!!
	s := Student{
		Name:       "John",
		Roll:       45672,
		Percentage: 30.00,
	}
	fmt.Println(s.Name)

	// CREATING AN INSTANCE OF DEPARTMENT!
	d := Department{
		Teachers: 25,
		Students: 70,
		Courses: []string{"SE", "CS", "IT"},
	}

	fmt.Println(d.Courses)

}
