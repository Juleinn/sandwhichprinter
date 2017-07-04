#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>

#define TIME 1000

int main(int argc, char ** argv)
{
	wiringPiSetupGpio();	

	int i;
	for(i=0; i<40; i++)
	{
		pinMode(i, INPUT);
		pullUpDnControl(i, PUD_DOWN);
		
	}
	
	
	while(1)
	{
		system("clear");
		for(i=0; i<40; i++)
		{
			printf("%d : %d\n", i, digitalRead(i));
		}

	}	

	printf("Done.\n");
}
