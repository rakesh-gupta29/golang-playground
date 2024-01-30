package exercise_server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count uint32
var mut sync.Mutex

func MountAndListen() {
	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/count", handleVisitorCount)
	fmt.Println("Server is active at :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func healthCheck(w http.ResponseWriter, r *http.Request) {

	mut.Lock()
	count++
	mut.Unlock()

	w.Write([]byte("thanks for visiting the homepage"))

}

func handleVisitorCount(w http.ResponseWriter, r *http.Request) {

	msg := fmt.Sprintf("Homepage has been visited %d times since the server is up", count)
	w.Write([]byte(msg))
}
