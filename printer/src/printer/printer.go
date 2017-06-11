package printer
/*
#cgo CFLAGS: -I/home/pi/sandwichprinter/include
#cgo LDFLAGS: /home/pi/sandwichprinter/printer/libmotor_api.a -lwiringPi
#include "motor.h"
#include "stepper.h"
 */
import "C"

import (
	"parser"
	"fmt"
	"errors"
	"time"
       )

/* A few global variables for the printer to work */
var conveyor	C.Motor
var disk_drive	C.Motor
var extruder1	C.Motor
var head	C.Stepper

var initialized bool

var sandwiches []parser.Sandwich
var sandwiches_states chan bool

func Init(){
	if !initialized {
		// setup the GPIO
		C.wiringPiSetupGpio()
		// init pins for motors and sensors
		disk_drive = C.motor_init(23, 24, 5, 6);
		conveyor = C.motor_init(25, 18, 0, 0);
		extruder1 = C.motor_init(12, 16, 0, 0);
		head = C.stepper_init(14, 15, 13, 20);
		C.set_speed(&head, 100);
		// calibrate stepper head

		// initialize channel
		sandwiches_states = make(chan bool, 100)

		initialized = true

		// launch asynchronous print loop
		go printer_loop()
	} else {
		fmt.Println("Warning : already initialized");
	}
}

func push_slice(){
	// push motor stop and backwards accordingly
}

func printer_loop(){
	for {
		<-sandwiches_states
		// fmt.Println("Printing sandwich : ", sandwiches[0])
		// clear first entry of the fifo
		// print_sandwich_fifo()
		time.Sleep(30 * time.Second)
		sandwiches = sandwiches[1:]
	}
}

func print(sandwich parser.Sandwich){
	// actual printing of a sandwich
}

func print_sandwich_fifo(){
	fmt.Println("Sandwich FIFO :")
	for i, s := range sandwiches{
		fmt.Println("\t", i, " : ", s);
	}
}

func Get_sandwiches() []parser.Sandwich{
	return sandwiches
}

func Print(sandwich parser.Sandwich) error{
	if !initialized {
		return errors.New("Error : printer not initialized");
	}

	// add sandwich to queue for printing
	sandwiches = append(sandwiches, sandwich)
	sandwiches_states<-true;
	return nil;
}


