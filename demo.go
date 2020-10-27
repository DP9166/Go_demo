package main

import "fmt"

type A interface {
	Foo()
}

type B interface {
	A
	Bar()
}

type T struct {}

func (t T) Foo() {
	fmt.Println("call Foo function from interface A.")
}

func (t T) Bar() {
	fmt.Println("call Bar function from interface B.")
}

func main() {
	//var animal = oop.NewAnimal("小狗")
	//var ianimal oop.IAnimal = oop.Dog{*animal}
	//fmt.Println(reflect.TypeOf(ianimal))
	//if dog, ok := ianimal.(oop.Dog); ok {
	//	fmt.Println(dog.GetName())
	//}
	//fmt.Println(reflect.TypeOf(ianimal))
}