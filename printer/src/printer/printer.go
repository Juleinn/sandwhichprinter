package printer

import(
  "fmt"
  "parser"
)

func Print(sandwich parser.Sandwich){
  fmt.Println("Printing sandwich : ", sandwich)
}
