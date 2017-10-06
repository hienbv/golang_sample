package main
import (
	"fmt"
    "github.com/xuri/excelize"
)

func main() {
	xlsx := excelize.NewFile()
    // Create a new sheet
    index := xlsx.NewSheet("Users")
    // SEt value of a cell
    
    xlsx.SetCellValue("Users", "A1", "ID")
    xlsx.SetCellValue("Users", "B1", "username")
    xlsx.SetCellValue("Users", "C1", "email")
    
    xlsx.SetCellValue("Users", "A2", "1")
    xlsx.SetCellValue("Users", "B2", "hienbv")
    xlsx.SetCellValue("Users", "C2", "hienbv@gmail.com")
    
    xlsx.SetActiveSheet(index)
    
    err := xlsx.SaveAs("users.xlsx")
    if err != nil {
        fmt.Println(err)
    }
}

