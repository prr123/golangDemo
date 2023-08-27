package main


import (
//	"os"
	"fmt"
	"unsafe"
	"reflect"
)

const BYTES_IN_INT32 = 4

func main() {

	intAr := [5]uint32{}

	for i:=0; i< 5; i++ {
		intAr[i] = uint32(i)
	}

	for i:=0; i< 5; i++ {
		fmt.Printf(" %d: %d\n", i+1, intAr[i])
	}

	bytlen := len(intAr)* BYTES_IN_INT32
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&intAr[0])), Len:  bytlen, Cap: bytlen}
	b := *(*[]byte)(unsafe.Pointer(&hdr))

	fmt.Printf(" len b: %d\n", len(b))
	for i :=0; i< len(b); i++ {
		fmt.Printf(" %2d: %d\n", i+1, b[i:i+1])
	}

	bAr := make([]byte, len(intAr)*4)
	for i:=0; i< 5; i++ {
		ival := intAr[i]
		bAr[i*4] = byte(ival)
    	bAr[i*4+1] = byte(ival >>8)
   	 	bAr[i*4+2] = byte(ival >>16)
    	bAr[i*4+3] = byte(ival >>24)

	}

	fmt.Printf(" len bAr: %d\n", len(bAr))
	for i :=0; i< len(bAr); i++ {
		fmt.Printf(" %2d: %d\n", i+1, bAr[i:i+1])
	}

}
