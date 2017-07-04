#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>

#define LED_PIN 14
#define CTRL_PIN_1 18
#define CTRL_PIN_2 15
#define TIME 1000

#define STEP_MOTOR_PIN_1 14
#define STEP_MOTOR_PIN_2 15

#define MOTOR_1_PIN_1 18
#define MOTOR_1_PIN_2 25

#define MOTOR_2_PIN_1 23
#define MOTOR_2_PIN_2 24

#define MOTOR_3_PIN_1 12
#define MOTOR_3_PIN_2 16

#include "motor/motor.h"
#include "stepper/stepper.h"

int main(int argc, char ** argv)
{
	wiringPiSetupGpio();
	Motor disk_drv = motor_init(23, 24, 5, 6);
	Motor conveyor = motor_init(25, 18, 0, 0);
	Motor extruder = motor_init(12, 16, 19, 19);
	Stepper head = stepper_init(14, 15, 13, 20);

	motor_run_stop_async(&disk_drv, FORWARD);
	motor_run_time_async(&conveyor, BACKWARDS, 3000);

	
	motor_wait(disk_drv);
	motor_wait(conveyor);
	
	motor_run_stop_async(&disk_drv, BACKWARDS);
	motor_run_time_async(&conveyor, FORWARD, 3000);

	motor_wait(disk_drv);
	motor_wait(conveyor);

	
	/*set_speed(&head, 100);
	// recalibrate
	step(head, 2000);	// will stop before 2000 steps
	int max = step(head, -2000);
	printf("%d steps\n", max);


	// init done
	// pull out first one
	motor_run_stop(disk_drv, BACKWARDS);
	motor_run_stop(disk_drv, FORWARD);

	// run conveyor under head
	motor_run_time_async(conveyor, BACKWARDS, 1000);

	// print 
	// move head midway
	step(head, max/2);

	// extrude some matter
	motor_run_time(extruder, FORWARD, 500);

	// run motor back
	motor_run_time(conveyor, FORWARD, 3000);
	
	// pull out second one
	motor_run_stop(disk_drv, BACKWARDS);
	motor_run_stop(disk_drv, FORWARD);
	
	// run complete product out
	motor_run_time(conveyor, BACKWARDS, 5000);

	printf("Done.\n");*/
}
