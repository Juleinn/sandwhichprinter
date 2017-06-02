#ifndef MOTORS_H_
#define MOTORS_H_

#include <wiringPi.h>
#include <stdio.h>
#include <stdlib.h>

struct Motor{
	int pin1, pin2;
	int sensor1, sensor2;
	int speed;			// for later pwm speed modulation
};
typedef struct Motor Motor;

// pin 0 will be considered none
Motor motor_init(int pin1, int pin2, int sens1, int sens2);

void motor_set_speed(Motor * m, int speed);

#define FORWARD 1
#define BACKWARDS 2

// runs the motor until it hits the corresponding sensor
void motor_run_stop(Motor m, int direction);

// run motor for a given amount of ms
void motor_run_time(Motor m, int direction, int ms);


// time beewteen to sensor updates (avoids threading consuming all processor time)
#define SAMPLE_DELAY 1	// 1 ms

#include <wiringPi.h>

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
