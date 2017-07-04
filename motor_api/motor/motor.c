#include "motor.h"

Motor motor_init(int pin1, int pin2, int sens1, int sens2)
{
	Motor m;
	m.pin1 = pin1;
	m.pin2 = pin2;
	m.sensor1 = sens1;
	m.sensor2 = sens2;

	m.speed = 100;

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
	int upTime = (CYCLE_TIME * m.speed) / 100;
	int downTime = CYCLE_TIME - upTime;

	// wait until corresponding sensor hit
	while(!digitalRead(sensor))
	{
		// update every SAMPLE_DELAY ms
		//	delay(SAMPLE_DELAY);
		digitalWrite(pin, HIGH);
		usleep(upTime);
		digitalWrite(pin, LOW);
		usleep(downTime);
		
	}

	digitalWrite(pin, LOW);

}

void motor_run_time(Motor m, int direction, int ms)
{
	// blind run : won't stop if hits sensor 
	int pin = direction == FORWARD ? m.pin1 : m.pin2;
	
	int upTime = (CYCLE_TIME * m.speed) / 100;
	int downTime = CYCLE_TIME - upTime;

	while(ms > 0)
	{
		digitalWrite(pin, HIGH);
		usleep(upTime);
		digitalWrite(pin, LOW);
		usleep(downTime);
		ms -= 10;
	}

	digitalWrite(pin, LOW);

}

void motor_set_speed(Motor * m, int speed)
{
	// 0 < speed < 100
	m->speed = speed > 100 ? 100 : speed < 0 ? 0 : speed;
}


