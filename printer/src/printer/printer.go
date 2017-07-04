package printer
/*
#cgo CFLAGS: -I/home/pi/sandwichprinter/include
#cgo LDFLAGS: /home/pi/sandwichprinter/printer/libmotor_api.a -lwiringPi
#include "motor.h"
#include "stepper.h"
#include "servo.h"
 */
import "C"

import (
	"parser"
	"fmt"
	"light"
	"errors"
//	"time"
       )

/* A few global variables for the printer to work */
var conveyor	C.Motor
var disk_drive	C.Motor
var extruder1	C.Motor
var extruder2	C.Servo
var head		C.Stepper

var bulb		light.Light
var yLED		light.Light
var gLED		light.Light

var initialized bool

var sandwiches []parser.Sandwich
var sandwiches_states chan bool

var stepper_width int

func Init(){
	if !initialized {
		// setup the GPIO
		C.wiringPiSetupGpio()

		// init lights
		gLED.Init(9);
		gLED.SetBlink(0, 1000);
		yLED.Init(10);
		yLED.SetBlink(1000, 0);
		bulb.Init(11);
		bulb.SetBlink(0, 1000);

		// init pins for motors and sensors
		disk_drive = C.motor_init(23, 24, 13, 6);
		conveyor = C.motor_init(18, 25, 21, 21);
		extruder1 = C.motor_init(16, 12, 19, 19);
		extruder2 = C.servo_init(5);
		head = C.stepper_init(14, 15, 20, 26);
		C.set_speed(&head, 100);


		// calibrate stepper head
		C.step(head, 2000);
		stepper_width =  int(C.step(head, -2000));

		// reset extruder head
		C.motor_set_speed(&extruder1, 35);
		C.motor_run_stop(extruder1, C.BACKWARDS);

		// reset 2nd extruder head
		C.servo_move(extruder2, 160);

		// initialize channel
		sandwiches_states = make(chan bool, 100)

		initialized = true

		yLED.SetBlink(0, 1000);
		gLED.SetBlink(1000, 0);

		// launch asynchronous print loop
		go printer_loop()
	} else {
		fmt.Println("Warning : already initialized");
	}
}

func push_slice(){
	C.motor_run_stop(disk_drive, C.BACKWARDS);
	// run a little longer to really push to the end of track
	C.motor_run_time(disk_drive, C.BACKWARDS, 500);
	C.motor_run_stop(disk_drive, C.FORWARD);
}

func printer_loop(){
		for {
		<-sandwiches_states
		// fmt.Println("Printing sandwich : ", sandwiches[0])
		// clear first entry of the fifo
		// print_sandwich_fifo()
		print(sandwiches[0]);
		sandwiches = sandwiches[1:]
	}
}

func print(sandwich parser.Sandwich){
	// set lightbulb
	bulb.SetBlink(1000, 0);
	// we don't want to print on the conveyor so discard any sandwich with no bread
	if len(sandwich.Slices) == 0{
		fmt.Println("Print warning : discarding sandwich because there is no bread");
		return
	}

	push_slice();
	C.motor_run_time(conveyor, C.FORWARD, 4000);

	// enter actual printing process
	// three channels required : one for the conveyor, one for the head and one for the extruder

	// push head back to the beginning
	C.step(head, -2000);


	// center head to aim for the bread
	C.step(head, 50);

	waitHead := make(chan bool)
	go func(){
		// step along while extruding
		C.step(head, C.int(stepper_width-200));
		waitHead<-true
	}()

	// extrude matter for a given amount of time
	C.motor_run_time(extruder1, C.FORWARD, 500);

	<-waitHead

	// run the conveyor for a while
	C.motor_run_time(conveyor, C.FORWARD, 500);

	// start a second pass the other way
	go func(){
		C.step(head, C.int(200-stepper_width));
		waitHead<-true;
	}()

	C.motor_run_time(extruder1, C.FORWARD, 500);
	<-waitHead

	// convey back to push second slice
	C.motor_run_time(conveyor, C.BACKWARDS, 4500);

	push_slice();

	// push the slice out
	C.motor_run_time(conveyor, C.FORWARD, 10000);
	

	bulb.SetBlink(0, 1000);
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


