package stock

import(
	"encoding/xml"
	"encoding/json"
	"os"
	"fmt"
)

type Slice struct{
	XMLName	xml.Name	`xml:"slice"`
	Name	string		`xml:",chardata"`
	Img		string		`xml:"img,attr"`
}

type Garnish struct{
	XMLName	xml.Name	`xml:"garnish"`
	Name	string		`xml:",chardata"`
	Img		string		`xml:"img,attr"`
}

type Stock struct{
	XMLName	xml.Name	`xml:"stock"`
	Slices	[]Slice		`xml:"slice"`
	Garnishes []Garnish `xml:"garnish"`
}

/* instance of the above structure loaded by Load()
 * Must be reloaded upon file changes */
var stock Stock;

/* Set to true after successfull Load() */
var initialized bool;

/* Initializes the stock (loads from file) 
 * May be called multiple times to reload */
func Load(){
	file, err := os.Open("site/stock.xml"); // this needs to accept another file later
	if err != nil{
		fmt.Println("Unable to open stock file");
		return;
	}
	// parse xml
	err = xml.NewDecoder(file).Decode(&stock);
	if err != nil{
		fmt.Println("Unable to parse stock file");
		return;
	}
	initialized = true;
}

/* Returns an instance of the stock as a structure */
func Get() Stock{
	return stock;
}

/* Returns an json string containing the string in json
 * format */
func GetJSON()([]byte, error){
	// marshal stock & return
	data, err := json.Marshal(&stock)
	if err != nil{
		fmt.Println("Unable to marshal stock");
		return nil, err;
	}

	return data, nil;
}

/* Sets the stock according to an XML string representing 
 * the new stock and saves it to the file */
func SetSON(json string){
	
}
