package main

import (
	"fmt"

)

func main()  {
	x := 27
	if x % 2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x) 
	}

	switch num := 15; {
	case num < 50:
		fmt.Println("%d is less then 50\n", num)
		fallthrough
	case num > 100:
		fmt.Println("%d is less then 100\n", num)
		fallthrough
	case num < 200:
		fmt.Println("%d is less then 200\n", num)
	}
}