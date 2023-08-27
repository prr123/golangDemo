package main


import (
//	"os"
	"fmt"
//	"unsafe"
)

func main() {

	intAr := [5]uint32{}

	for i:=0; i< 5; i++ {
		intAr[i] = uint32(i)
	}

	for i:=0; i< 5; i++ {
		fmt.Printf(" %d: %d\n", i+1, intAr[i])
	}
}
