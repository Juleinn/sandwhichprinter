all:
	make -C motor/
	make -C stepper/
	gcc -c demo.c -o demo.o
	gcc -c printer.c -o printer.o
	gcc demo.o motor/motor.o stepper/stepper.o -o demo -lwiringPi -lpthread
	gcc printer.o motor/motor.o stepper/stepper.o -o printer -lwiringPi -lpthread
