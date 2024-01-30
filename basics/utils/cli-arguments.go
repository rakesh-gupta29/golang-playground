package utils

import (
	"fmt"
	"os"
)

// print all the CLI args given
func PrintArgs() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
