// Echo2 prints its command line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i) + ": " + arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
