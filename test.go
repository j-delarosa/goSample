package main

import (
	"fmt"
)

type Animal struct {
	NumberOfLegs int
	Name         string
}

func (doggy Dog) GetLegCount() int {
	return doggy.NumberOfLegs
}

func (doggy Dog) bark() {
	fmt.Printf("%s went woof woof!!\n", doggy.Name)
}

type Dog struct {
	Animal
}

func main() {
	dog := Dog{}
	dog.Name = "timmy"
	dog.NumberOfLegs = 4
	fmt.Println(dog.Name)
	dog.bark()
	fmt.Println(dog.GetLegCount())
}
