/**
 \file servo.h
 \brief servomotor control
 * */

#ifndef SERVO_H_
#define SERVO_H_

#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>

#define TIME_180 	1200000 	// 1200 ms for full movement
#define SERV_CYCLE_TIME 	21000 	// 21 ms a cycle

struct Servo{
	int pin;	// pin for servo control
	int angle;	// last reached angle
};
typedef struct Servo Servo;

/**
 \brief initializes the servo with a given pin. WiringPi must me initialized
 \param pin the pin to initialize
 */
Servo servo_init(int pin);

/**
 \brief moves servo to given angle. Releases tension on servo after a timeout proportional to the difference of angle
 \param servo the pointer to an initialized servo
 \param angle the angle to reach
 * */
void servo_move(Servo servo, int angle);



#endif
