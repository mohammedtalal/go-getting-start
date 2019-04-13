package main

import "fmt"

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func varadicParams(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
func testDefer() {
	defer foo()
	bar()
}

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func main() {
	// fmt.Println(greet("mohammed ", "talal"))
	// varadicParams(2, 3, 4, 5, 6)
	// fmt.Println("The total is", sum(1, 2, 3, 4, 5, 6))
	// testDefer()
}
