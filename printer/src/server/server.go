package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"parser"
	"printer"
	"encoding/json"
	)

func print_sandwich(rw http.ResponseWriter, req *http.Request)  {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		fmt.Println("Unable to parse sandwich")
	}
	// parse xml file then send to printer
	sandwich, err := parser.ParseXML(body)
	if err != nil {
		fmt.Println("error parsing sent xml : ", err)
		// tell sender failure
		fmt.Fprintf(rw, "Error parsing xml\n")
	} else {
		// send it to the printer
		printer.Print(sandwich)
		fmt.Fprintf(rw, "Sandwich will be printer soon\n")
	}
}

func get_sandwiches(rw http.ResponseWriter, req *http.Request) {
	sandwiches := printer.Get_sandwiches()
	jsonTxt, err := json.Marshal(sandwiches)
	if err == nil{
		fmt.Fprintf(rw, "%s", string(jsonTxt))
	} else {
		fmt.Println("{'error': 'unable to get sandwiches'")
	}
}

func static(rw http.ResponseWriter, req * http.Request){
	path := "site/" + req.URL.Path[len("/static/"):]
	fmt.Println(path)
	body, err := ioutil.ReadFile(path)
	if err == nil{
		fmt.Fprintf(rw, "%s", body)
	} else {
		fmt.Println(err)
		fmt.Fprintf(rw, "Unable to get document : ", path)
	}
}

func Start(port string){
	// initialize printer
	printer.Init()
	http.HandleFunc("/print", print_sandwich)
	http.HandleFunc("/static/", static)
	http.HandleFunc("/sandwiches", get_sandwiches)
	http.ListenAndServe(":" + port, nil)
}
