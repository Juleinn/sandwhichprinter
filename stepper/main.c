#include "stepper.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
	wiringPiSetupGpio();
	Stepper s = stepper_init(14, 15, 13, 20);
	set_speed(&s, 100);
	step(s, 300);
	printf("Done\n");
}
