package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
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
		io.WriteString(rw, "Sandwich will be printed soon");
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

func aliasHandler(route string) func(http.ResponseWriter, *http.Request){
	return func(rw http.ResponseWriter, req * http.Request){
		body, err := ioutil.ReadFile(route)
			if err == nil{
			fmt.Fprintf(rw, "%s", body);
		} else {
			fmt.Println(err);
			fmt.Fprintf(rw, "Unable to get home page");
		}
	}
}

func handleHome(rw http.ResponseWriter, req * http.Request){
	body, err := ioutil.ReadFile("site/index.html")
	if err == nil{
		fmt.Fprintf(rw, "%s", body);
	} else {
		fmt.Println(err);
		fmt.Fprintf(rw, "Unable to get home page");
	}
}

func Start(port string){
	// initialize printer
	printer.Init()

	// init server
	// handle all static routes
	fs := http.FileServer(http.Dir("site"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// handle dynamic content
	http.HandleFunc("/sandwiches", get_sandwiches)
	http.HandleFunc("/print", print_sandwich)

	// handle static route aliases (for home)
	homeHandler := aliasHandler("site/index.html");
	orderHandler := aliasHandler("site/order.html");
	configHandler := aliasHandler("site/config.html");
	http.HandleFunc("/", homeHandler);
	http.HandleFunc("/order", orderHandler);
	http.HandleFunc("/config", configHandler);


	http.ListenAndServe(":" + port, nil)
}
