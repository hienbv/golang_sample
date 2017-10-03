package main
import (
	"fmt"
	"os"
	"strconv"
)

type Vector struct {
	number int
}

func (v Vector) uocso(X int) bool {
	return (X % v.number == 0)
}

func ds_uocso(y int) []int {
	var output []int
	if y <= 0 {
		return output
	}
	if y == 1 {
		return append(output, 1)
	}

	//fmt.Println(len(output))
	for i:=1; i<y;i++ {
		if (Vector{i}).uocso(y) {
			output = append(output, i)
		}

	}
	return output
}

func main() {
	//v := Vector{3}
	argsWithProg := os.Args

	len, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println("Param is error")
		return
	}

	fmt.Println("DS so hoan hao: ")
	for i:=1; i<=len;i++ {
		ds := ds_uocso(i)
		sum := 0
		for _, num := range ds {
			sum += num
		}

		if sum == i {
			fmt.Println(i, ds)
		}
	}

}