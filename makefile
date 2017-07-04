all:
	make -C motor_api
	cp motor_api/libmotor_api.a printer/
	cp motor_api/motor/motor.h include/
	cp motor_api/stepper/stepper.h include/
	cp motor_api/servo/servo.h include/	
