package main
import (
	"fmt"
)

func fibonacci () func() uint64 {
	x,y := uint64(0),uint64(1)
	return func () uint64 {
		x,y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	
	for i:=0; i<50; i++ {
		fmt.Println(f())
	}
}

