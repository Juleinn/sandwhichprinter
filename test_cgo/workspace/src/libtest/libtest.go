package libtest
/*
#cgo CFLAGS: -I/home/anton/test_cgo
#cgo LDFLAGS: -L/home/anton/test_cgo/ libtest.a
#include <stdio.h>
#include <stdlib.h>
#include "libtest.h"
*/
import ("C")

func Test(){
  C.libtestfunc()
}
