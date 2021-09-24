package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var fuga string

func main()  {
	const HTTPStatusOK = 200
	firstName, lastName := "john", "hoge" 
	age := 22
	updateName(&firstName)
	fmt.Println(firstName, lastName, age)

	var integer32 int32 = 32767
	var integer16 int16 = 127
	fmt.Println(int32(integer16) + integer32)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)

	sum, mul := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum", sum)
	fmt.Println("Mul", mul)
}

func sum(number1 string, number2 string) (result int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	result = int1 + int2
	mul = int1 * int2
	return 
}

func updateName(name *string)  {
	*name = "fuff"
}