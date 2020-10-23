package main

import (
	"calc/src/animal"
	"fmt"
)

func main()  {
	ani := animal.NewAnimal("狗")
	dog := animal.Dog{ani}
	fmt.Println(dog.GetName())
}
