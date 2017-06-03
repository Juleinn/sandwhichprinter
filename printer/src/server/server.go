package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"parser"
	"printer"
	)

func print_sandwich(rw http.ResponseWriter, req *http.Request)  {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		fmt.Println("Unable to parse sandwich")
	}
	fmt.Println(string(body))
	// parse xml file then send to printer
	sandwich, err := parser.ParseXML(body)
	if err != nil {
		fmt.Println("error parsing sent xml : ", err)
		// tell sender failure
		fmt.Fprintf(rw, "Error parsing xml\n")
	} else {
		fmt.Println("Printing sandwich : ", sandwich)
		// send it to the printer
		printer.Print(sandwich)
		fmt.Fprintf(rw, "Sandwich will be printer soon\n")
	}
}

func Start(port string){
	http.HandleFunc("/print", print_sandwich)
	http.ListenAndServe(":" + port, nil)
}
