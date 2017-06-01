package main

import(
  "fmt"
  "parser"
  "os"
  "printer"
)

func main(){
  if len(os.Args) != 2 {
    fmt.Println("Use : main filename")
    return
  }
  fmt.Println("Hello world. Parser. ")
  sdw_file := os.Args[1]
  sandwich, err := parser.Parse(sdw_file)
  if(err != nil){
    fmt.Println("Unable to parse sandwich properly")
    os.Exit(1)
  }
  printer.Print(sandwich)
}
