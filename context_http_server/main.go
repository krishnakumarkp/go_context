package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("handler started")
	defer log.Printf("Handler Ended")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hellow world!")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func handlerNoContext(w http.ResponseWriter, r *http.Request) {
	log.Printf("handler started")
	defer log.Printf("Handler Ended")
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "hellow world!")

}
