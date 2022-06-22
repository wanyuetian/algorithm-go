package main

import "fmt"

type Car interface {
	Drive()
}

type Truck struct{}

func (t Truck) Drive() {}

type Bike struct{}

func (b Bike) Drive() {}

func main() {
	var c Car = Truck{}
	b := c.(Bike)
	fmt.Println(b)
	//m := make(map[int]int, 100)
}
