package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main()  {
	// 例外処理
	defer func()  {
		handler := recover()
			if handler != nil {
				fmt.Println("main(): recover", handler)
			}
	}()

	var num int

	fmt.Scanln(&num)
	fibonacci(num)
}

// フィボナッチ配列
func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i - 1] + nums[i - 2]
	}

	return nums
}

type Parson struct {
	ID int
	FirstName string `json: "name"`
	LastName string
	Address string `json: "address,omitempty"`
}

type Employee struct {
	Parson
	ManagerID int
}

func employee()  {
	employees := []Employee{
		{Parson: Parson{LastName: "Done", FirstName: "Johe"}}, {
			Parson:    Parson{LastName: "Cambell", FirstName: "Devid"},
			ManagerID: 0,
		}}

	data, _ := json.Marshal(employees)
	fmt.Println(data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)

	// employeCopy := &employe
	// employeCopy.FirstName = "David"
	fmt.Println(employees)
}

func fmap()  {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	fmt.Println(studentsAge)
	fmt.Println(studentsAge["john"])

	age, exist := studentsAge["chi"]
	if exist {
		fmt.Println(age)
	} else {
		fmt.Println("couldn'T be found")
	}

	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func sliceCopy()  {
	letters := []string{"A", "B", "C", "D"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Bfter", slice2)
}

func arrays()  {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}

	fmt.Println("\nAll at once:", twoD)
}

func array()  {
	var a [4]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(len(a))
	fmt.Println(a[len(a) -1])

	cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[99])
}

func panic() string {
	val := 0

	for {
			fmt.Print("Enter number: ")
			fmt.Scanf("%d", &val)

			switch {
			case val < 0:
				// panic("You entered a negative number!")
			case val == 0:
					fmt.Println("0 is neither negative nor positive")
			default:
					fmt.Println("You entered:", val)
			}
	}
}

// 例外処理
func highlow(high int, low int)  {
	if high < low {
		fmt.Println("Panic!")
		// panic("highlow() low greater than hig")
	}

	defer fmt.Println("Deferred: highlow(", high, ",", low, ",")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low + 1)
}

func fizzBuzz()  {
	for i := 1; i < 100; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}
	if number > 1 {
		return true
	} else {
		return false
	}
}

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