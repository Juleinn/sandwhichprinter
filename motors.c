#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>

int main(int argc, char ** argv)
{
	wiringPiSetupGpio();

	int pin = atoi(argv[1]);

	pinMode(18, OUTPUT);
	

	digitalWrite(pin, LOW);

	int i=0;
	for(i=0; i>-1; i++)
	{
		int j;
		for(j=0; j<200; j++)
		{
		digitalWrite(pin, HIGH);
		delay(TIME);
		digitalWrite(pin, LOW);
		delay(TIME);
		}
	}
	printf("Done.\n");
}
