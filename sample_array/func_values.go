package main
import (
	"fmt"
	"math"
)

func compute(add func(float64, float64) float64, a,b float64) float64{
	return add(a,b)
}

func add(x,y float64) float64 {
	return x+y
}

func main() {
	fmt.Println(compute(math.Pow, 4, 5))
	fmt.Println(compute(add, 4, 5))
}

