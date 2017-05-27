#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>

#define LED_PIN 14
#define CTRL_PIN_1 18
#define CTRL_PIN_2 15
#define TIME 1000

int main(int argc, char ** argv)
{
	wiringPiSetupGpio();

	int pin = atoi(argv[1]);

	pinMode(pin, INPUT);
	pullUpDnControl(pin, PUD_DOWN);
	int i=0;
	for(i=0; i>-1; i++)
	{
		int res = digitalRead(pin);
		printf("pin = %d\n", res);
	}
	printf("Done.\n");
}
