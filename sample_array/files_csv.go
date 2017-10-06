package main
import (
	"io"
    "os"
    "fmt"
    "encoding/csv"
)

func main() {
	f, err := os.Open("users.csv")    
    if err != nil {
        fmt.Println("File not found!")
    }
    
    defer f.Close()
    
    reader := csv.NewReader(f)
    reader.Comma = ','
    lineCount := 0
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
			return
        }
        
        fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
        for i := 0; i < len(record); i++ {
			fmt.Println(" ", record[i])
		}
        fmt.Println()
        lineCount++
    }
}

