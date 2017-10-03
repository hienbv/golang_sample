package main
import (
	"fmt"
	"os"
	"strconv"
    "math"
    "errors"
)

type Param struct {
	a float64
    b float64
    c float64
}

func (p Param) delta() float64 {
	return math.Pow(p.b, 2) - (4*p.a*p.c)
}

func giaiPT(a,b,c float64) (x1, x2 float64, err error) {
    param := Param{a, b, c}
    delta := param.delta()
    fmt.Printf("Delta: %f\n", delta)
    
    if delta == 0 {
        x1, x2 = param.b/(2*param.a), param.b/(2*param.a)
    } else if delta < 0 {        
        err = errors.New("Vo nghiem")
        return
    } else {
        deltaSqrt := math.Sqrt(delta)
        x1 = (-b - deltaSqrt)/(2*a)
        x2 = (-b + deltaSqrt)/(2*a)
    }
    
    return
}

func main() {
	//v := Vector{3}
	argsWithProg := os.Args

	a, err := strconv.ParseFloat(argsWithProg[1], 64)
    b, err := strconv.ParseFloat(argsWithProg[2], 64)
    c, err := strconv.ParseFloat(argsWithProg[3], 64)
    
    x1,x2,err := giaiPT(a,b,c)
    
    if err != nil {
		fmt.Println(err)
		return
	} else{
        fmt.Printf("Nghiem pt: %2f, %2f\n", x1, x2)
    }
}