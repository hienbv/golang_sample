package main
import (
	"fmt"
)

type Vector struct {
	lat float64
	long float64
}

var m map[string]Vector

var m2 = map[string]Vector{
	"hanam": Vector{
		2.00001, 6.31111,
	},
	"hanoi": Vector{
		4.00001, 5.31111,
	},
}

func main() {
	m = make(map[string]Vector)
	
	m["hanoi"] = Vector{4.00001, 5.31111}
	
	fmt.Println(m)
	
	delete(m2, "hanoi")	
	fmt.Println(m2)
	
	var test = make(map[string]string)
	test["a"] = "42"
	test["b"] = "23"
	elem, ok := test["a"]
	fmt.Println("Value: ", elem, " Present?: ", ok)
	
	delete(test, "b")	
	elem, ok = test["b"]
	fmt.Println("Value: ", elem, " Present?: ", ok)
}

