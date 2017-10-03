package main
import (
    "fmt"
    "os"
	"strconv"   
)

func ucln(x,y int) int {
    var r,b int
    if x > y {
        r = x % y
        b = y
    } else {
        r = y % x
        b = x
    }   
    
    if r == 0 {
        return b
    } else {
        return ucln(b, r)
    }
}

func bcnn(x,y int) int {
    return abs(x*y)/ucln(x, y)
}

func abs(x int) int {
    if x < 0 {
        return x * (-1)
    }
    
    return x
}

func main() {
    argsWithProg := os.Args
    x, err := strconv.Atoi(argsWithProg[1])
    y, err := strconv.Atoi(argsWithProg[2])
    
    if err != nil {
		fmt.Println("Param is error")
		return
	}
    
    fmt.Printf("UCLN: %d\n", ucln(x, y))
    
    fmt.Printf("BCNN: %d\n", bcnn(x, y))
}
