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


/* Inits a stepper motor for 2 given pins */
Stepper stepper_init(int pin1, int pin2, int sens1, int sens2);

/* Set the speed for a motor */
void set_speed(Stepper * s, int speed);

/* step the motor by the given amount (positive and negative) 
returns the number of steps actually done (stops at sensor cut) */
int step(Stepper s, int steps);


#endif
