package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
)
func main() {

	router := mux.NewRouter()
	//Håndterer hva som skjer når en klient aksesserer "/ "
	router.HandleFunc("/", HelloClient)

	srvr := http.Server {
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10* time.Second,
		WriteTimeout: 10* time.Second,
	}
	panic(srvr.ListenAndServe())
}

func HelloClient (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client")
}
