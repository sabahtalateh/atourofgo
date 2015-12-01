package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(9302491230981)
	fmt.Print("My rand is ", rand.Intn(10))
}
