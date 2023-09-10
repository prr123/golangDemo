// -- test to time lookup for func table

package main

import (
	"fmt"
)


type handler = func(s1, s2 string)


func main() {

	funcA:= func(a,b string) {
		fmt.Printf("hello %s and %s from A\n", a, b)
	}
	funcB:= func(a,b string) {
		fmt.Printf("hello %s and %s from B\n", a, b)
	}
	var funcC handler = func(a,b string) {
		fmt.Printf("hello %s and %s from C\n", a, b)
	}

	funcTab := [3]handler {
		funcA,
		funcB,
		funcC,
	}

	funcTab[0]("a1", "a2")
	funcTab[1]("b1", "b2")
	funcTab[2]("c1", "c2")

	fmt.Println("end!")
}

