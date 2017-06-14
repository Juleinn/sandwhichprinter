/**
\file motor.h
\brief motor control functions with end of track detection
*/
#ifndef MOTOR_H_
#define MOTOR_H_

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


/** 
\brief initializes motor with end of track sensors
\param pin1, pin2 the GPIO for running motor in directions 1,2
\param sens1, sens2 the GPIO input pin for end of track sensor 1,2
\return the generated motor descriptor */
Motor motor_init(int pin1, int pin2, int sens1, int sens2);

/**
\brief sets the speed for a motor (unimplemented now)
\param m a pointer to an initialized motor structure
\param speed the new speed to be set
*/
void motor_set_speed(Motor * m, int speed);

#define FORWARD 1
#define BACKWARDS 2

/** 
\brief runs the motor in the given direction until it hits its end-of-track sensor. Only checks for sensor 1 if run in direction 1, sensor 2 if run in direction 2
\param m an initialized motor descriptor
\param direction the direction (BACKWARDS=2, FORWARD=1)
*/
void motor_run_stop(Motor m, int direction);

/** 
\brief runs the motor for a specified amount of time without checking for end-of-track sensor detection
\param m an initialized motor descriptor
\param direction the direction
\param ms the amount of time to run the command in ms
*/
void motor_run_time(Motor m, int direction, int ms);


// time beewteen to sensor updates (avoids threading consuming all processor time)
#define SAMPLE_DELAY 1	// 1 ms

#endif
