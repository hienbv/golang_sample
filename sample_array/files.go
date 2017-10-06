package main
import (
	"log"
    "os"
    "fmt"
)

func main() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_APPEND, 0755)
    //f.WriteString("hienbvxxx\n")
    aaa, err := os.Lstat("notes.txt")
    fmt.Println(aaa.ModTime())
    if err != nil {
        log.Fatal(err)
    }
    
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}

