package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.22, 55.69,
	},
	"Google": Vertex{
		44.36, 33.12,
	},
}

func main() {
	fmt.Println(m)
}
