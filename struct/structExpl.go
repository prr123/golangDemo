// program sructExpl.go
// program that illustrates anonymous structures
// Author: prr, azul software
// Date 7 June 2023
// copyright 2023 prr, azul software
//

package main

import (
	"os"
	"fmt"
)

type anon struct {
	test
	Nam string
}

type test struct {
	Name string
	count int
}

func main() {

	numArgs := len(os.Args)

	useStr := "structExpl [help]"
	helpStr := "program that demonstrates the use of anonymous structures and fields\n"

	if numArgs > 2 {
		if os.Args[1] == "help" {
			fmt.Printf("usage is: %s\n", useStr)
			fmt.Printf("help:\n%s\n", helpStr)
			os.Exit(1)
		}
	}

	// begin
	tst := test {
		Name: "hello",
		count: 1,
	}

	an := anon {
		test: tst,
		Nam: "AnonTest",
	}

	PrintTest(tst)
	PrintAnon(an)

	fmt.Printf("success\n")
}

func PrintTest(tst test) {
	fmt.Println("********* test **********")
	fmt.Printf("  Name: %s\n", tst.Name)
	fmt.Printf("  count: %d\n", tst.count)

	fmt.Println("********* end test ******")
}

func PrintAnon(anon anon) {
	fmt.Println("*********** anon ************")
	fmt.Printf("test: %v\n", anon.test)
	PrintTest(anon.test)
	fmt.Printf("Name: %s\n", anon.Nam)

	fmt.Println("********* end anon **********")

}
