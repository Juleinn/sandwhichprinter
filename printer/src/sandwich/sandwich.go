package main

import(
		"fmt"
		"parser"
		"os"
		"printer"
		"server"
      )

func main(){
	/* 	launch with sandwich stored as a file with 	sandwich -f filename.sdw 	*
	 *	launch with online service with 			sandwich -s port 			*/
	if len(os.Args) < 3 {
		fmt.Println("Use :\n-f [filename] or -s [port]")
		fmt.Println("(optional) -p [working path]")
		fmt.Println("Params must be in right order")
		return
	}

	if len(os.Args) == 5{
		// set up working path accordingly
		err := os.Chdir(os.Args[4]);
		if err != nil {
			fmt.Println("Unable to change working dir to : ", os.Args[4]);
			fmt.Println(err);
			return;
		}
	}

	if os.Args[1] == "-s" {
		// fmt.Println("Starting server...")
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
