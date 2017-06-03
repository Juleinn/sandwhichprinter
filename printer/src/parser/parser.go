package parser

import(
  "encoding/xml"
  "os"
  "fmt"
)

type Sandwich struct{
  Slices    []string `xml:"slice"`
  Garnishes []string `xml:"garnish"`
}

func ParseFile(filename string) (Sandwich, error){
  fmt.Println("Parsing file : ", filename)
  var sandwich Sandwich
  file, err := os.Open(filename)
  if err != nil {
    return sandwich, err
  }
  defer file.Close()

  if err := xml.NewDecoder(file).Decode(&sandwich); err != nil {
    return sandwich, err
  }

  return sandwich, nil
}

func ParseXML(data []byte)(Sandwich, error){
	fmt.Println("Parsing data : \n", string(data));
	var sandwich Sandwich
	if err := xml.Unmarshal(data, &sandwich); err != nil{
		return sandwich, err
	}

	return sandwich, nil
}
