package main

import (
	"calc/src/animal"
	"fmt"
)

func main() {
	var animal = animal.NewAnimal("小狗")

	fmt.Println(animal.GetName())
	//var ianimal oop.IAnimal = oop.Dog{animal}
	//if dog, ok := ianimal.(oop.Dog); ok {
	//	fmt.Println(dog.GetName())
	//}
}