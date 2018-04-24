package main

import (
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"log"
	"html/template"
)
type Barnehager struct {
	Entries []struct {
		Latitude  string `json:"latitude"`
		Navn      string `json:"navn"`
		ID        string `json:"id"`
		Longitude string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

type PlastSortering *struct {
	Entries []struct {
		Navn        string `json:"navn"`
		Plast       string `json:"plast"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

type FylkesNummer struct {
	ContainedItems []struct {
		ID string `json: "id"`
		Description string `json: "description"`
		Owner string `json: "owner"`
		Label string `json: label`
	}
}

type BussHoldeplass struct {
	BusStops []struct {
		FullName     string  `json:"FullName"`
		ShortName    string  `json:"ShortName"`
	} `json:"BusStops"`
}

type Parkering struct {
	Dato                string `json:"Dato"`
	Klokkeslett         string `json:"Klokkeslett"`
	Sted                string `json:"Sted"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	AntallLedigePlasser string `json:"Antall_ledige_plasser"`
	}

	//variablaer av structene for å holde på JSON-dataen
var barnehage Barnehager
var holdeplass BussHoldeplass
var fylker FylkesNummer
var plastSort PlastSortering
var park []Parkering

const navbar = "html/header.html"

func main() {
	http.HandleFunc("/", Hello1)
	http.HandleFunc("/2", Hello2)
	http.HandleFunc("/3", Hello3)
	http.HandleFunc("/4", Hello4)
	http.HandleFunc("/5", Hello5)


	srvr := http.Server {
		Addr:         ":8080",
		ReadTimeout:  10* time.Second,
	}
	panic(srvr.ListenAndServe())
}

func Hello1 (w http.ResponseWriter, r *http.Request) {
	getAndUnmarshal("https://hotell.difi.no/api/json/bergen/lekeplasser?", &barnehage)
	readAndExecuteHTML(w, "html/barnehage.html", barnehage)
}

func Hello2 (w http.ResponseWriter, r *http.Request) {
	getAndUnmarshal("https://hotell.difi.no/api/json/stavanger/miljostasjoner", &plastSort)
	readAndExecuteHTML(w,"html/index.html", plastSort)
}

func Hello3 (w http.ResponseWriter, r *http.Request) {
	getAndUnmarshal("https://register.geonorge.no/api/subregister/sosi-kodelister/kartverket/fylkesnummer-alle.json?", &fylker)
	readAndExecuteHTML(w, "html/fylker.html", fylker)
}

func Hello4 (w http.ResponseWriter, r *http.Request){
	getAndUnmarshal("http://sanntidsappservice-web-prod.azurewebsites.net/busstops?format=json", &holdeplass)
	readAndExecuteHTML(w,"html/holdeplass.html", holdeplass )
}

func Hello5(w http.ResponseWriter, r *http.Request){
	getAndUnmarshal("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json", &park)
	readAndExecuteHTML(w, "html/parkering.html", park)
}

/*
 getAndUnmarshal gjør et API-kall og leser JSON-dataen og legger det inn i structs-variablene.
 */
func getAndUnmarshal(s string, v interface{}) {
	//Henter jsondataen i s
	res, err := http.Get(s)
	defer res.Body.Close()
	errorHandling(err)
	//Leser kroppen til jsondataen
	jsonBytes, err2 := ioutil.ReadAll(res.Body)
	errorHandling(err2)
	//Legger jsondaten inn i variabelen v
	errorHandling(json.Unmarshal(jsonBytes, &v))
}

/*
	readAndExecute leser en string som er en filepath til en html-fil og kjører
	html-filen sammen med struct-variabelen
*/
func readAndExecuteHTML(w http.ResponseWriter,s string, v interface{}) {
	//Lager template av HTML-filene
	tmpl, err := template.ParseFiles(s, navbar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Kjører HTML-templaten til writeren og mater den med interfacet
	if err := tmpl.Execute(w, v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
/*
	Enkel errorhandler
 */
func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}