package main
import "fmt"

func main() {
	var a []int
	printSlicee("a", a)

	a = append(a, 0)
	printSlicee("a", a)

	a = append(a, 1)
	printSlicee("a", a)

	a = append(a, 2, 3, 4)
	printSlicee("a", a)
}

func printSlicee(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
