package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func showMeData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	go func() {
		n := 0
		for {
			if n%1000000 == 0 {
				// These should all print to browser and terminal
				fmt.Printf("Loop count = %v\n", n)
				fmt.Fprintf(w, "Loop count = %v\n", n)
			}
			select {
			case <-ctx.Done():
				// Line won't show in browser as request context has been cancelled
				fmt.Fprintln(w, "Request Context expired")
				fmt.Println("Request Context expired")
				return
			default:
				n++

			}
		}
	}()
	time.Sleep(1 * time.Second)
}

func main() {
	http.HandleFunc("/showMeData", showMeData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
