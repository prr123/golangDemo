// readInp
// program to read an input
// Author: prr, azul software
// Date: 30 May 2023
// copyright 2023 prr azulsoftware
//

package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
)

func main() {
    fmt.Println("Enter a string")
    reader := bufio.NewReader(os.Stdin)
    str, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("string: %s\n", str)

    fmt.Printf("Enter a number: ")
    str, err = reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
	numStr := str[:len(str)-1]
	num, err := strconv.Atoi(numStr)
	if err!= nil {log.Fatal(err)}
    fmt.Printf("number: %d\n", num)
}
