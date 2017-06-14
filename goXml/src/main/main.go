package main

import(
	"fmt"
	"encoding/xml"
	"os"
)

type Slice struct{
	XMLName	xml.Name	`xml:"slice"`
	Name	string		`xml:",chardata"`
	Img		string		`xml:"img,attr"`
}

type Stock struct{
	XMLName	xml.Name	`xml:"stock"`
	Slices []Slice		`xml:"slice"`
}

func main(){
	var stock Stock;
	file, _ := os.Open("example.xml");

	_ = xml.NewDecoder(file).Decode(&stock);

	fmt.Println(stock)
	fmt.Println(stock.Slices[0].Name)
	fmt.Println(stock.Slices[0].Img)
}
