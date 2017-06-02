package printer
/*
#cgo CFLAGS: -I/home/pi/sandwichprinter/printer/
#cgo LDFLAGS: /home/pi/sandwichprinter/printer/libtest.a
#include "libtest.h"
*/
import "C"

import (
  "fmt"
  "parser"
)

func Print(sandwich parser.Sandwich){
	fmt.Println("Printing sandwich : ", sandwich)
	C.printTest()
}
