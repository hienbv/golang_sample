package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithProg := os.Args

	len, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println("Param is error")
		return
	}

	fmt.Println("_______Bang Cuu Chuong_______")

	for i:=1; i<=len; i++ {
		for j:=1; j<=len; j++ {
			fmt.Printf("%2d*%2d=%3d   ", j, i, (i*j))
		}

		fmt.Println("\n")
	}
}