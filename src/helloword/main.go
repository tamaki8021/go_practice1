package main

import (
	"fmt"
	"io"
	"os"
	// "io"
	// "os"
)

func main()  {
	// 例外処理
	// defer func()  {
	// 	handler := recover()
	// 		if handler != nil {
	// 			fmt.Println("main(): recover", handler)
	// 		}
	// }()

	createfile()

	// highlow(2, 0)
	// fizzBuzz()
	// fmt.Println("Program number less than 20")

	// for number := 1; number <= 20; number++ {
	// 	if findprimes(number) {
	// 		fmt.Printf("%v ", number)
	// 	}
	// }
}

// 例外処理
// func highlow(high int, low int)  {
// 	if high < low {
// 		fmt.Println("Panic!")
// 		panic("highlow() low greater than high")
// 	}

// 	defer fmt.Println("Deferred: highlow(", high, ",", low, ",")
// 	fmt.Println("Call: highlow(", high, ",", low, ")")

// 	highlow(high, low + 1)
// }

// func fizzBuzz()  {
// 	for i := 1; i < 100; i++ {
// 		switch {
// 		case i % 3 == 0 && i % 5 == 0:
// 			fmt.Println("FizzBuzz")
// 		case i % 3 == 0:
// 			fmt.Println("Fizz")
// 		case i % 5 == 0:
// 			fmt.Println("Buzz")
// 		default:
// 			fmt.Println(i)
// 		}
// 	}
// }

// func findprimes(number int) bool {
// 	for i := 2; i < number; i++ {
// 		if number % i == 0 {
// 			return false
// 		}
// 	}
// 	if number > 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func createfile()  {
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file.")
	}

	newfile.Sync()
}