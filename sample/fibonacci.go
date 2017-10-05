package main
import (
	"fmt"
    "time"
)


func fibonacci(number int) int {
    if number < 2 {
        return number
    }
    
    return fibonacci(number - 1) + fibonacci(number - 2)
}

func makeTimestamp() float64 {
    return float64(time.Now().UnixNano()) / float64(time.Second)
}

func main() {	
    time1 := makeTimestamp();
    fmt.Printf("%d \n", fibonacci(40))
    
    fmt.Printf("%f \n", (makeTimestamp() -time1))
}