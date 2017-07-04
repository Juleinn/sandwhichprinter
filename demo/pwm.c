#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>
#include <unistd.h>


#include "motor/motor.h"
#include "stepper/stepper.h"

int main(int argc, char ** argv)
{
	wiringPiSetupGpio();

	#define TOTALCYCLE 1000
	int timeUp = atoi(argv[1]);
	int timeDown = TOTALCYCLE - timeUp;

	// a few pwm steps 
	pinMode(16, OUTPUT);
	pinMode(12, OUTPUT);

	digitalWrite(16, LOW);
	digitalWrite(12, LOW);

	int i;
	for(i=0; i<500; i++){
		// set a 50% pwm
		// with total cycle time of 2ms
		digitalWrite(16, HIGH);
		usleep(timeUp);
		digitalWrite(16, LOW);
		usleep(timeDown);

	}

	for(i=0; i<500; i++){
		// set a 50% pwm
		// with total cycle time of 2ms
		digitalWrite(12, HIGH);
		usleep(timeUp);
		digitalWrite(12, LOW);
		usleep(timeDown);

	}


	printf("Done.\n");
}
