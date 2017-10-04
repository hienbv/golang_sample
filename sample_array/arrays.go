package main
import (
	"fmt"
)

func main() {
	var balance = []int {2,3,4,5,6,7,8}
	var slice = balance[1:3] // 3 pt đầu tiên của mảng, lấy bắt đầu từ vị trí 1
							 // Slices are like references to arrays
	
	slice[0] = 100
	
	fmt.Println(balance[:2])
	fmt.Println(balance[1:])
	
	balance[6] = 1;
	balance = balance[:4]
	
	fmt.Println(len(balance))
	fmt.Println(cap(balance))
	
	cakes := make([]int, 10, 10)
	fmt.Printf("%s len=%d, cap=%d %v \n", "cakes", len(cakes), cap(cakes), cakes)
	
	cakes = append(cakes, 10)
	cakes = append(cakes, 11)
	fmt.Printf("%s len=%d, cap=%d %v \n", "cakes", len(cakes), cap(cakes), cakes)
	
	for k,v := range cakes {
		fmt.Printf("elem[%d]=%d \n", k, v)
	}
	fmt.Printf("----------------------\n")
	for _,v := range cakes {
		fmt.Printf("elem[]=%d \n", v)
	}
	
	fmt.Printf("----------------------\n")
	
	for k,_ := range cakes {
		fmt.Printf("elem[%d]=%d \n", k, cakes[k])
	}
	
}

