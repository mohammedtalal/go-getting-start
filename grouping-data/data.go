package main

import "fmt"

func arrLen() {
	var arr [5]int
	arr[2] = 100
	fmt.Printf("%x\n length for arr = %x\n", arr, len(arr))
}

func compositeLiteral() {
	// slice allow u to roup together values of the same types
	a := 3.22
	x := []int{1, 2, 3, 4, int(a)}
	fmt.Println(x)
}

func rangeClause() {
	x := []int{5, 6, 7, 8, 9}
	for i, v := range x {
		fmt.Println(i, v)
	}
}

func sliceASlice() {
	x := []int{5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(x)
	fmt.Println(x[2:])
	fmt.Println(x[2:6])

}

func appendToSlice() {
	x := []int{5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(x)

	x = append(x, 55, 66, 77)
	fmt.Println(x)

	y := []int{30, 40, 50, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func removeFromSliceUsingAppend() {
	x := []int{5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(x)

	y := []int{30, 40, 50, 60, 88, 99}
	x = append(x, y...)
	fmt.Println(x)

	// remove from x
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func sliceWithMake() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func multiDimensionalArr() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp[0][2])
	fmt.Println(xp[1][3])
}

func mapSyntax() {
	m := map[string]int{
		"age":       24,
		"colleague": 4,
	}

	fmt.Println(m)
	fmt.Println(m["age"])

	v, ok := m["egypt"]
	fmt.Println(v, ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(ok)
		fmt.Println(v)
	}
}

func main() {
	// arrLen()
	// compositeLiteral()
	// rangeClause()
	// sliceASlice()
	// appendToSlice()
	// removeFromSliceUsingAppend()
	// sliceWithMake()
	// multiDimensionalArr()
	// mapSyntax()
}
