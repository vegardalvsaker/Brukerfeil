package main

import (
	"net/http"
	"time"
	"fmt"
)
func main() {
	//Håndterer hva som skjer når en klient aksesserer "/"
	http.HandleFunc("/", HelloClient)

	srvr := http.Server {
		Addr:         ":8080",
		ReadTimeout:  10* time.Second,
		WriteTimeout: 10* time.Second,
	}
	panic(srvr.ListenAndServe())
}

func HelloClient (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client")
}
