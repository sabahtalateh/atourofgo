package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Mike", 42}
	b := Person{"Vitya", 10}
	fmt.Println(a, b)
}

