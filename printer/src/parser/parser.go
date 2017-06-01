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

func Parse(filename string) (Sandwich, error){
  fmt.Println("Parsing ", filename)
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
