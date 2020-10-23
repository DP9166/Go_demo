package main

import (
	"calc/src/oop2"
	oop1 "calc/src/opp1"
	"fmt"
)

func main() {
	var num1 oop1.Number = 1
	var num2 oop2.Number2 = &num1

	if num3, ok := num2.(oop1.Number1); ok {
		fmt.Println(num3.Equal(1))
	}
}