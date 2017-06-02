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

#endif
