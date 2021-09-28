package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go func () {
		for {
			time.Sleep(time.Second)
			fmt.Println(time.Now())
		}
	}()

	http.HandleFunc("/", func(w  http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":12345", nil)
}
