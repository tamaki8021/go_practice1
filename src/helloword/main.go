package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	// "time"
	// "log"
)

type Parson struct {
	ID int
	FirstName string `json: "name"`
	LastName string
	Address string `json: "address,omitempty"`
}

type Employee struct {
	ID int
	FirstName string 
	LastName string
	Address string 
	// Parson
	// ManagerID int
}

type triangle struct {
	size int
}

type square struct {
	size int
}

type upperstring string

type coloredTriangle struct {
	triangle
	color string
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

type Circle struct {
	radius float64
}

type Stringer interface {
	String() string
}

type dollars float32

type database map[string]dollars

func main()  {
	// 例外処理
	// defer func()  {
	// 	handler := recover()
	// 		if handler != nil {
	// 			fmt.Println("main(): recover", handler)
	// 		}
	// }()

	// 呼び出し元でのエラー処理
	// employee, err := getInformation(1001)
	// if errors.Is(err, ErrNotFound) {
	// 	fmt.Printf("NOT FOUND: %v\n", err)
	// } else {
	// 	fmt.Print(employee)
	// }

	// ログをファイルに抽出
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// log.SetOutput(file)
	// log.Print("Hey, I'm a log!")

	// t := triangle{3}
	// s := square{4}
	// t.doubleSize()
	// fmt.Println("Size:", t.size)
	// fmt.Println("Perimeter:", t.perimeter())
	// fmt.Println("Perimeter(square):", s.perimeter())

	// s := upperstring("Learning Go!")
	// fmt.Println(s)
	// fmt.Println(s.Upper())

	// t := coloredTriangle{triangle{3}, "blue"}
	// fmt.Println("Size:", t.size)
	// fmt.Println("Perimeter:", t.perimeter())

	var s Shape = Square{3}
	printInfomation(s)

	c := Circle{3}
	printInfomation(c)

	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func printInfomation(s Shape)  {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}
func (t triangle) perimeter() int {
	return t.size * 3
}

var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil

	// 一時的なエラーの場合に、再試行ロジックを実行する方法
	// for tries := 0; tries < 3; tries++ {
	// 	employee, err := apiCallEmpoyee(1000)
	// 	if err == nil {
	// 		return employee, nil
	// 	}

	// 	fmt.Println("Server is not responding, retrying ...")
	// 	time.Sleep(time.Second * 2)
	// }
	// return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmpoyee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
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

// func employee()  {
// 	employees := []Employee{
// 		{Parson: Parson{LastName: "Done", FirstName: "Johe"}}, {
// 			Parson:    Parson{LastName: "Cambell", FirstName: "Devid"},
// 			ManagerID: 0,
// 		}}

// 	data, _ := json.Marshal(employees)
// 	fmt.Println(data)

// 	var decoded []Employee
// 	json.Unmarshal(data, &decoded)
// 	fmt.Printf("%v", decoded)

// 	employeCopy := &employe
// 	employeCopy.FirstName = "David"
// 	fmt.Println(employees)
// }

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