#include "motor.h"

Motor motor_init(int pin1, int pin2, int sens1, int sens2)
{
	Motor m;
	m.pin1 = pin1;
	m.pin2 = pin2;
	m.sensor1 = sens1;
	m.sensor2 = sens2;
	
	// init pinm and menmor pinm
	pinMode(pin1, OUTPUT);
	pinMode(pin2, OUTPUT);
	digitalWrite(pin1, LOW);
	digitalWrite(pin2, LOW);

	pinMode(sens1, INPUT);
	pinMode(sens2, INPUT);
	pullUpDnControl(sens1, PUD_DOWN);
	pullUpDnControl(sens2, PUD_DOWN);

	return m;	
}

void motor_run_stop(Motor m, int direction)
{

	int pin = direction == FORWARD ? m.pin1 : m.pin2;
	int sensor = direction == FORWARD ? m.sensor1 : m.sensor2;
	
	// discard if no sensor set
	if(pin == 0)
		printf("Warning : no end of track sensor set\n");

	// run motor
	digitalWrite(pin, HIGH);

	// wait until corresponding sensor hit
	while(!digitalRead(sensor))
	{
		// update every SAMPLE_DELAY ms
		delay(SAMPLE_DELAY);
	}
	

	// have a small delay before stop
	delay(400);
	// stop motor
	digitalWrite(pin, LOW);
}

void motor_run_time(Motor m, int direction, int ms)
{
	// blind run : won't stop if hits sensor 
	int pin = direction == FORWARD ? m.pin1 : m.pin2;
	// run motor
	digitalWrite(pin, HIGH);
	// delay
	delay(ms);
	// stop motor
	digitalWrite(pin, LOW);
}


