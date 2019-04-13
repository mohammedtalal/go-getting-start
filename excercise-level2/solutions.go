package main

import (
	"fmt"
	"time"
)

/*
Print from 1 to 10.000
*/
func ques1() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

/*
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
65
	U+0041 'A'
	U+0041 'A'
	U+0041 'A'
*/
func ques2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}

/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/
func ques3() {
	myBirthDate := 1995
	for myBirthDate <= 2019 {
		fmt.Println(myBirthDate)
		myBirthDate++
	}
}

/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/
func ques4() {
	myBirthDate := 1995
	for {
		if myBirthDate > 2019 {
			break
		}
		fmt.Println(myBirthDate)
		myBirthDate++
	}
}

/*
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
*/
func ques5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}

/*
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
*/
func ques6() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}

/*
Create a program that shows the “if statement” in action.
*/
func ques7() {
	x := true
	if x {
		fmt.Println("that is true")
	} else {
		fmt.Println("not true :( ")
	}
}

/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/
func ques8() {
	x := "Moneypenny"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}

/*
Create a program that uses a switch statement with no switch expression specified.
*/
func ques9() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

/*
Create a program that uses a switch statement with no switch expression specified.
*/
func ques10() {
	favSport := "bingpong"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "bingpong":
		fmt.Println("go to club to play bingPong!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}

/*
Create a program that uses a switch statement with no switch expression specified.
*/
func ques11() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func conversion() {
	var float float64 = 10.1
	fmt.Println(float, "become", int(float))
}

func getDay() {
	today := time.Now().Weekday()

	switch today {
	case time.Saturday:
		fmt.Println("Today is saturday")
	case time.Sunday:
		fmt.Println("Today is sunday")
	case time.Monday:
		fmt.Println("Today is monday")
	case time.Tuesday:
		fmt.Println("Today is tuesday")
	default:
		fmt.Println("next days in week")
	}
}

func main() {
	// ques1()
	// ques2()
	// ques3()
	// ques4()
	// ques5()
	// ques7()
	// ques8()
	// ques9()
	// ques10()
	// ques11()
	// conversion()
	getDay()
	// goDefer()
}
