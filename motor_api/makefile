all:
	make -C motor/
	make -C stepper/
	make -C servo/
	ar rcs libmotor_api.a motor/motor.o stepper/stepper.o servo/servo.o
	ranlib libmotor_api.a
