package printer
/*
#cgo CFLAGS: -I/home/pi/sandwichprinter/include
#cgo LDFLAGS: /home/pi/sandwichprinter/printer/libmotor_api.a -lwiringPi
#include "motor.h"
#include "stepper.h"
*/
import "C"

import (
  "fmt"
  "parser"
)

func Print(sandwich parser.Sandwich){
	fmt.Println("Printing sandwich : ", sandwich)
	var m C.Motor
	m = C.motor_init(1, 5, 3, 2)
	fmt.Println("Motor run stop : ")
	C.motor_run_time(m, 1, 1000)
}


