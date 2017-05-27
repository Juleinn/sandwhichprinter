#include "stepper.h"

Stepper stepper_init(int pin1, int pin2, int sens1, int sens2)
{
	Stepper s;
	s.pin1 = pin1;
	s.pin2 = pin2;
	s.sensor1 = sens1;
	s.sensor2 = sens2;
	
	// init pins and sensor pins
	pinMode(pin1, OUTPUT);
	pinMode(pin2, OUTPUT);
	digitalWrite(pin1, LOW);
	digitalWrite(pin2, LOW);

	pinMode(sens1, INPUT);
	pinMode(sens2, INPUT);
	pullUpDnControl(sens1, PUD_DOWN);
	pullUpDnControl(sens2, PUD_DOWN);

	return s;	
}

void set_speed(Stepper * s, int speed)
{
	s->speed = speed;
}

int step(Stepper s, int steps)
{
	// will step for a given number of steps 
	// or stop if it hits a sensor (pin1->sensor1, pin2->sensor2)
	int pin = steps > 0 ? s.pin1 : s.pin2;
	int sensor = steps > 0 ? s.sensor1 : s.sensor2;
	steps = steps > 0 ? steps : -steps;

	int time = 200 / s.speed;
	
	int i;
	for(i=0; i<steps; i++)
	{
		if(digitalRead(sensor))
			return i;
		digitalWrite(pin, HIGH);
		delay(time);
		digitalWrite(pin, LOW);
		delay(time);
	}
	return i;
}

