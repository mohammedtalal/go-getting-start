package main

import (
	"fmt"
)

func arraysQues() {
	x := [5]int{42, 43, 44, 45}
	for key, val := range x {
		fmt.Println(key, val)
	}

	fmt.Printf("%T \n", x)
}

func sliceQues() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for k, v := range x {
		fmt.Println(k, v)
	}
	fmt.Printf("%T \n", x)
}

func sliceFromSlice() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func playWithSlice() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

/*
* output [42, 43, 44, 48, 49, 50, 51]
 */
func sliceDeleteUsingAppend() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func sliceLenAndCap() {
	x := []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}
	fmt.Printf("len of %d and cap of x : %d \n", len(x), cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func sliceOfSlice() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xx := [][]string{x1, x2}
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(xx)

	for _, val := range xx {
		for k2, val2 := range val {
			fmt.Println(k2, val2)
		}
		// fmt.Println(k, val)
	}
}

func playWithMap() {
	m := map[string][]string{
		`talal`:           []string{`developer`, `cs`, `mohamed`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`} // add new record
	delete(m, `no_dr`)                                           // delete one record

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}

func main() {
	// arraysQues()
	// sliceQues()
	// sliceFromSlice()
	// playWithSlice()
	// sliceDeleteUsingAppend()
	// sliceLenAndCap()
	// sliceOfSlice()
	// playWithMap()
}
