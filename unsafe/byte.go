package main


import (
	"fmt"

)


func main () {

	val := 4
	fmt.Printf("val: %d\n", val)

	b:=[4]byte{}
	b[0] = byte(val)
	b[1] = byte(val >>8)
	b[2] = byte(val >>16)
	b[3] = byte(val >>24)

	for i:=0; i< len(b); i++ {
		fmt.Printf(" %d: %d\n",i, b[i])
	}
}
