package main
import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z int = int(f)



	fmt.Print(x, y, z, f)
}
