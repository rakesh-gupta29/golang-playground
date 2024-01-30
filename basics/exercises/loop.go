package exercises

import "fmt"

func LoopsAndDataCollection() {
	// first and last condition here is optional
	// for each iteration, a new set of data is passed which is later garbage collected.

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
