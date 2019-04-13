package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:

first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values,
 ranging over the elements in the slice which stores the favorite flavors.
*/
func playingWithStruct() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}

func playingWithStructAndMap() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, val := range m {
		fmt.Println(val.first)
		fmt.Println(val.last)
		for i, val := range val.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func playingWithMultiStruct() {
	tr := truck{
		vehicle: vehicle{
			doors: "four",
			color: "black",
		},
		fourWheel: true,
	}

	sd := sedan{
		vehicle: vehicle{
			doors: "four double",
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(tr)
	fmt.Println(sd)
	fmt.Println(tr.doors)
	fmt.Println(sd.doors)
}

func anonymousStruct() {
	person := struct {
		name string
		age  int
	}{
		name: "talal",
		age:  25,
	}
	fmt.Println(person)
}

func main() {
	// playingWithStruct()
	// playingWithStructAndMap()
	// playingWithMultiStruct()
	anonymousStruct()
}
