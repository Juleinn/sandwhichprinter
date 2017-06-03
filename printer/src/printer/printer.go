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

/* A few global variables for the printer to work */
var conveyor C.Motor
var disk_drive C.Motor
var extruder1 C.Motor

var initialized bool

func Init(){
//	conveyor = C.motor_init()
//	disk_drive = C.motor_init()
//	extruder1 = C.motor_init()
}

func push_slice(){
	// push motor stop and backwards accordingly
}

func Print(sandwich parser.Sandwich){
	fmt.Println("Printing sandwich : ", sandwich)
	var m C.Motor
	m = C.motor_init(1, 5, 3, 2)
	fmt.Println("Motor run stop : ")
	C.motor_run_time(m, 1, 1000)
}


