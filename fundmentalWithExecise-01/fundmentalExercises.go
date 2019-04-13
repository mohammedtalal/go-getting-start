package main

import (
	"fmt"
)

func main() {
	// printDecimalBinaryHex()
	// writeExpressionsAndAssignValue()
	// typedAndUntypedConstant()
	// shiftsTheBits()
	// stringLiteral()
	printLastFourYearUsingIota()
}

func test() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte

	a = ar[2:]
	b = ar[3:5]
	fmt.Printf("%d\n %d\n", a, b)
}

/* Exercise 1 */
func printDecimalBinaryHex() {
	x := 42
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}

/* Exercise 2 */
func writeExpressionsAndAssignValue() {
	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 43)
	d := (42 != 42)
	e := (42 > 42)
	f := (42 < 42)

	fmt.Println(a, b, c, d, e, f)
}

/* Exercise 3 */
func typedAndUntypedConstant() {
	const (
		a     = 42
		b int = 43
	)
	fmt.Println(a, b)
	fmt.Printf("%T\t %T", a, b)
}

/* Exercise 4 */
func shiftsTheBits() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}

/* Exercise 5 */
func stringLiteral() {
	a := `here is something
		as 
		a
		raw string
		"literal"
		another bla`
	fmt.Println(a)
}

/* Exercise 6 */
func printLastFourYearUsingIota() {
	const (
		a = 2017 + iota
		b = 2017 + iota
		c = 2017 + iota
		d = 2017 + iota
	)
	fmt.Println(a, b, c, d)
}


