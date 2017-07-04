/**
\file stepper.h
\brief contains functions for a 2-pin bound stepper motor (1 pin per step in each direction) with end-of-track detection
*/
#ifndef STEPPER_H_
#define STEPPER_H_

#include <wiringPi.h>
#include <stdio.h>
#include <stdlib.h>

struct Stepper{
	int pin1, pin2;		// control pin
	int sensor1, sensor2;	// end of track sensor pin
	int speed;		// a speed coefficient. A default fast speed would be 100, be it can be set to more than than
};
typedef struct Stepper Stepper;


/**
\brief initializes a stepper motor descriptor structure with output pins and sensor input pins
\param pin1, pin2 the pins for the stepper to step
\param sens1, sens2 the pins for the end of track sensors
*/
Stepper stepper_init(int pin1, int pin2, int sens1, int sens2);

/**
\brief sets the running speed for the stepper motor
\param s an initialized stepper descriptor
\param speed the new speed
*/
void set_speed(Stepper * s, int speed);

/**
\brief step the motor by the given amount (positive and negative) 
returns the number of steps actually done (stops at sensor cut) 
\param s an initialized stepper descriptor
\param steps the number of steps performed
\return the true number of steps performed
*/
int step(Stepper s, int steps);


#endif
