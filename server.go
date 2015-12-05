package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Listening on 3012")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, it works"))
	})
	http.ListenAndServe(":3012", mux)
}

