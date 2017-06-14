all:
	make -C motor/
	make -C stepper/
	ar rcs libmotor_api.a motor/motor.o stepper/stepper.o
	ranlib libmotor_api.a
