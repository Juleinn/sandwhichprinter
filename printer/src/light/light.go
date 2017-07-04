package light
/*
#cgo CFLAGS: -I/home/pi/sandwichprinter/include
#cgo LDFLAGS: /home/pi/sandwichprinter/printer/libmotor_api.a -lwiringPi
#include <wiringPi.h>
*/
import "C"

import "time"

/* Structure describing a light */
type Light struct{
	pin			int
	running		chan bool
	upTime		int // ms
	downTime	int // ms
}

func (l *Light) Init(p int) {
	C.pinMode(C.int(p), C.OUTPUT)
	l.pin = p;
	l.running = make(chan bool);
	// init to off state then start goroutine
	l.upTime = 500;
	l.downTime = 500;
	// start routine
	go l.LightRun();
}

func (l *Light) SetBlink(upTime int, downTime int){
	l.upTime = upTime;
	l.downTime = downTime;
}

func (l *Light) Stop(){
	l.running<-false
}

func (l *Light) LightRun() {
	cont := true;
	for cont {
		select{
			case cont = <-l.running:
				break;
			default:{
				// blink light here
				C.digitalWrite(C.int(l.pin), C.HIGH);
				time.Sleep(time.Duration(l.upTime) * time.Millisecond);
				C.digitalWrite(C.int(l.pin), C.LOW);
				time.Sleep(time.Duration(l.downTime) * time.Millisecond);
				break;
			}
		}
	}
}




