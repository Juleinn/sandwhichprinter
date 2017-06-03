package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
	)

func print_sandwich(rw http.ResponseWriter, req *http.Request)  {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		fmt.Println("Unable to parse sandwich")
	}
	fmt.Println(string(body))

}

func start(port int){
	http.HandleFunc("/print", print_sandwich)
	http.ListenAndServe(":8080", nil)
}
