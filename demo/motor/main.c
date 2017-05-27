#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>

#include "motor.h"

int main(int argc, char** argv)
{
	wiringPiSetupGpio();
	Motor m = motor_init(23, 24, 5, 6);
	motor_run_stop(m, BACKWARDS);
	motor_run_stop(m, FORWARD);	



	Motor m2 = motor_init(25, 18, 0, 0);
	motor_run_time(m2, FORWARD, 1500);
	
	Motor m3 = motor_init(12, 16, 0, 0);
	motor_run_time(m3, FORWARD, 1000);
	
	motor_run_time(m2, BACKWARDS, 1500);
	
	// push out next one
	motor_run_stop(m, BACKWARDS);
	motor_run_stop(m, FORWARD);	
	
	motor_run_time(m2, FORWARD, 5000);

	printf("Done\n");
}
