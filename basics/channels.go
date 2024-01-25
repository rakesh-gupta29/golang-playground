package channels

// channels: concurrency primitive: primary way of communication between goroutines.

// wait groups are used to make the program wait while the goroutines are executing.
// channels is for pushing data from one routine to another routine.

// they work in form of stacks: last in, first out.

import (
	"fmt"
	"time"
)

func makecall(timer time.Duration, resch chan time.Duration) {
	resch <- timer

}
func main() {

	namesch := make(chan time.Duration)
	start := time.Now()

	go makecall(time.Second*20, namesch)

	go makecall(time.Second*100, namesch)

	nameA := <-namesch
	fmt.Println(nameA)

	timeB := <-namesch
	fmt.Println(timeB)

	fmt.Println("total time taken to run", time.Since(start))
}
