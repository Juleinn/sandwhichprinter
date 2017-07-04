#include "servo.h"

Servo servo_init(int pin)
{
	Servo servo;
	servo.pin = pin;
	
	pinMode(servo.pin, OUTPUT);
	digitalWrite(servo.pin, LOW);

	return servo;
}

void servo_move(Servo servo, int angle)
{
	// compute 21 ms cycles
	// = (angle / 180) * TOTAL_TIME / SERV_CYCLE_TIME
	// int cycles = ( angle * TIME_180 ) / ( 180 * SERV_CYCLE_TIME);
	int cycles = 60;	// 60 cycles for full movement

	int upTime = (int) ((((float) angle / 180.0f) * 1600.0f) + 600.0f);

	printf("Cycles = %d", cycles);

	int i;
	for(i=0; i<cycles; i++)
	{
		digitalWrite(servo.pin, HIGH);
		usleep(upTime);
		digitalWrite(servo.pin, LOW);
		usleep(SERV_CYCLE_TIME - upTime);
	}

	return;
}

