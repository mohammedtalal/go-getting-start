package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	active bool
}

func simpleStruct() {
	p1 := person{
		first: "mohamed",
		last:  "talal",
		age:   25,
	}

	p2 := person{
		first: "john",
		last:  "wiki",
		age:   25,
	}
	fmt.Println(p1, p2)

	fmt.Println(p1.first, p2.last)
	fmt.Println(p2.first, p2.last)
}

func embeddedStruct() {
	sa := secretAgent{
		person: person{
			first: "mohamed",
			last:  "talal",
			age:   25,
		},
		active: false,
	}
	if sa.active {
		fmt.Printf("My full name is %v and i am %v years old", sa.person.first+sa.person.last, sa.person.age)
	} else {
		fmt.Println("You are not active user")
	}

}

func anonymousStruct() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "mohamed",
		last:  "talal",
		age:   25,
	}
	fmt.Println(p1)
}

func main() {
	// simpleStruct()
	// embeddedStruct()
	// anonymousStruct()
}
