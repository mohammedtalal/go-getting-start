package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// firstForLoop()
	// pointer()
	// shortFor()
	// printASCII()
	// ifStatement()
	// isEven()
	// switchCase()
	// scanner()
	// interfaces()
	// readingInConsole()
}

/* FIrst loop */
func firstForLoop() {
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}
}

func pointer() {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("address of variable: %x\n", &a)

	fmt.Printf("Address stored in ip variable: %x\n", ip)
	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

}

func shortFor() {
	x := 2
	// for {
	// 	if x < 10 {
	// 		break
	// 	}
	// 	fmt.Println(x)
	// 	x++
	// }
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}

func printASCII() {
	for i := 30; i < 40; i++ {
		fmt.Printf("%x\t %d\t %#x\t %#U\n", i, i, i, i)
	}
}

func ifStatement() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
}

func isEven() {
	x := 10
	if x%2 == 0 {
		fmt.Println("value of x is even")
	}
}

func switchCase() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
	case (4 == 4):
		fmt.Println("also true, does it print?")
	}
}

func readingInConsole() {
	for {
		fmt.Print("Enter Name: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println("Welcome,", text)
		fmt.Print("Press 'C' To Continue: ")
		choice, _ := reader.ReadString('\n')
		if !strings.EqualFold(strings.TrimSpace(choice), "c") {
			fmt.Println("Good Bye")
			return
		}
	}
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func interfaces(a interface{}) {
	fmt.Println(a)
}
