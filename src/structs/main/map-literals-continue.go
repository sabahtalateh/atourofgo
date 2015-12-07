package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.22, 20.12,},
	"Google": {37.22, 58.12},
}

func main() {
	fmt.Println(m)
}
