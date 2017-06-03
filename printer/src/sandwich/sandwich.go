package main

import(
		"fmt"
		"parser"
		"os"
		"printer"
		"server"
      )

func main(){
	if len(os.Args) != 3 {
		fmt.Println("Use :\n-f [filename]")
		fmt.Println("-s [port]")
		return
	}

	if os.Args[1] == "-s" {
		fmt.Println("Starting server...")
		server.Start(os.Args[2])
	} else if os.Args[1] == "-f" {
		sdw_file := os.Args[2]
		sandwich, err := parser.ParseFile(sdw_file)
		if(err != nil){
			fmt.Println("Unable to parse sandwich properly")
			os.Exit(1)
		}
		printer.Print(sandwich)
		fmt.Println("Printing file : ", os.Args[2])
	} else {
		fmt.Println("Invalid parameters. Execute with no params for args")
	}
}
