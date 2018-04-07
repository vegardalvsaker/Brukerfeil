package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	//"encoding/json"
	//"html/template"
	"encoding/json"
	"log"
)

type JsonStruct struct {
	Entries []struct {
		Latitude    string `json:"latitude"`
		Navn        string `json:"navn"`
		Plast       string `json:"plast"`
		GlassMetall string `json:"glass_metall"`
		TekstilSko  string `json:"tekstil_sko,omitempty"`
		Longitude   string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

var m JsonStruct
var myClient = &http.Client{Timeout:10 * time.Second}

type Profile struct {
	Name     string
	Hobbies []string
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", Hello1)
	router.HandleFunc("/2", Hello2)
	router.HandleFunc("/3", Hello3)
	router.HandleFunc("/4", Hello4)
	router.HandleFunc("/5", Hello5)


	srvr := http.Server {
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10* time.Second,
		WriteTimeout: 10* time.Second,
	}
	panic(srvr.ListenAndServe())
}

func Hello1 (w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, client")

	foo1 := JsonStruct{}
	getJson("https://hotell.difi.no/api/json/stavanger/miljostasjoner", foo1)
	js, err := json.Marshal(foo1)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	//w.Write(js)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	if err := json.Unmarshal(js, &m); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "lol", &m)

}

func Hello2 (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client2")
}
func Hello3 (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client3")
}
func Hello4 (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client4")
}
func Hello5 (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client5")
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}