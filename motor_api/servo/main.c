#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>
#include "servo.h"

int main()
{
	wiringPiSetupGpio();

	Servo servo = servo_init(5);

	servo_move(servo, 90);
	servo_move(servo, 160);
}
