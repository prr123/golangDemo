package main


import (
//	"os"
	"fmt"
	"unsafe"
	"reflect"
)

const BYTES_IN_INT32 = 4

type st struct {
	num uint32
	txt2 [24]byte
//	txt []byte
}

func main() {

	stAr := [5]st{}

	for i:=0; i< 5; i++ {
		stAr[i].num = uint32(i)
		str := fmt.Sprintf("hello %d",i+1)
//		stAr[i].txt = []byte(str)
		for j:=0; j<len(str); j++ {
			stAr[i].txt2[j] = str[j]
		}
	}

	for i:=0; i< 5; i++ {
		fmt.Printf(" %d: %d %s\n", i+1, stAr[i].num, stAr[i].txt2)
	}

	totLen := 0
	for i:=0; i< 5; i++ {
		stLen := int(unsafe.Sizeof(stAr[i]))
		fmt.Printf(" size st[%d]: %d\n", i+1, stLen)
		totLen += stLen
	}
//	bytlen := len(intAr)* BYTES_IN_INT32

	blen := len(stAr)*int(unsafe.Sizeof(stAr[0]))
	fmt.Printf(" size star: %d %d\n",blen, totLen)


	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&stAr[0])), Len:  totLen, Cap: totLen}
	b := *(*[]byte)(unsafe.Pointer(&hdr))

	fmt.Printf("Array of Struct stAr size: %d\n", len(b))

	x :=st{}
	xLen := int(unsafe.Sizeof(x))
	numSt := totLen/xLen
	fmt.Printf("Struct Size: %d Struct Count: %d\n", xLen, numSt)

//	tB := []byte{}
	tstNum := uint32(0)
/*
	tB := b[xLen:xLen+4]
	fmt.Println("tB:")
	for j:=0; j<4; j++ {
		fmt.Printf("%d:",tB[j])
	}
	fmt.Println()
*/
	for i:=0; i<numSt; i++ {
		tB := b[i*xLen:i*xLen+4]
//		fmt.Printf("tB: %d %d\n", len(tB), cap(tB))
/*
		fmt.Println("tB:")
		for j:=0; j<4; j++ {
			fmt.Printf("%d:",tB[j])
		}
		fmt.Println()
*/
		tP := unsafe.Pointer(&tB[0])
		tstNum = *(*uint32)(tP)

		tS := b[i*xLen+4:i*xLen+24]
		tstStr := string(tS)
		fmt.Printf("num[%d]: %d %s\n", i, tstNum, tstStr)

	}


}
