package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answear"] = 42
	fmt.Println("The value:", m["Answear"])

	m["Answear"] = 48
	fmt.Println("The value:", m["Answear"])

	delete(m, "Answear")
	fmt.Println("The value:", m["Answear"])

	v, ok := m["Answear"]
	fmt.Println("The value:", v, "Present?", ok)
}
