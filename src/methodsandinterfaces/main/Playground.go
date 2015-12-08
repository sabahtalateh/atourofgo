package main

import "fmt"

type Car struct {
	wheels int
}

func (c *Car) ride() (string, error) {
	if (c.wheels >= 100) {
		return "err", TooMuchWheels(c.wheels)
	}
	return fmt.Sprintf("riding with %v wheels", c.wheels), nil
}

type TooMuchWheels int

func (e TooMuchWheels) Error() string {
	return fmt.Sprintf("Too much wheels %v", int(e))
}

type SuperCar Car

func main() {
	c := Car{220}
	fmt.Println(c.ride())
}






