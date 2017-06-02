package libtest
/*
#cgo CFLAGS: -I/home/pi/sandwichprinter/test_cgo
#cgo LDFLAGS: /home/pi/sandwichprinter/test_cgo/libtest.a
#include "libtest.h"
*/
import ("C")

func Test(){
  C.libtestfunc()

