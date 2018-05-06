package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
	"html/template"
)

type Sted struct {
	Sted string
}

type WeatherData struct {
	Sted        string
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Currently struct {
		Time                int     `json:"time"`
		Summary             string  `json:"summary"`
		Icon                string  `json:"icon"`
		PrecipIntensity     float64    `json:"precipIntensity"`
		PrecipProbability   float64     `json:"precipProbability"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		Humidity            float64 `json:"humidity"`
		Pressure            float64 `json:"pressure"`
		WindSpeed           float64 `json:"windSpeed"`
	} `json:"currently"`
}

/*Variabel for å oppdatere placeholder i index.html
Starter med "Søk etter sted..."*/
var by = Sted {
	Sted: "Søk etter sted...",
}

var weath WeatherData

//variable holding Latitude and Longitude
var latLng string

//API-url with custom API-key
const weatherUrlWithKey = "https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50"
//Query-parametre for å få norsk værdata med metric units
const langAndUnits = "?lang=nb&units=si"

func main() {
	server()
}

func index (w http.ResponseWriter, r *http.Request) {
	//Lager template av HTML-filene
	tmpl, err := template.ParseFiles("index.html")
	httpErrorHandling(w, err)

	//Henter det som blir submitet fra html-formen og håndterer feil Method
	getForm(r)

	//Fjerner paranteser fra koordinatene slik at de skal passe inn i darksky-urlen
	latLng = latLngFormat(latLng)

	//Kjører HTML-templaten til writeren og mater den med interfacet
	if err := tmpl.Execute(w, by); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runForecast (w http.ResponseWriter, r *http.Request) {
	//Hvis en bruker går på /forecast uten å submitte koordinater
	if len(latLng) <= 0 {
		t, _ := template.ParseFiles("feil.html")
		t.Execute(w, nil)
	} else {

	fmt.Println(r.URL.Path)
	//Setter sammen API-urlen sammen med koordinatene for å få en fungerende url
	fmt.Println(latLng)
	//Setter sammen API-urlen sammen med koordinatene
	latAndLang := latLng + langAndUnits
	strSlice := []string{weatherUrlWithKey, latAndLang}
	getUrl := strings.Join(strSlice, "/")
	fmt.Println("Weather data collected from: ", getUrl)

	getAndUnmarshal(getUrl, &weath)

	t, err := template.ParseFiles("forecast.html", "above25.html")
	errorHandling(err)
	t.Execute(w, weath)
	}
}

//Server-funksjon som håndterer paths
func server() {
	http.HandleFunc("/", index) // setting router rule
	http.HandleFunc("/forecast", runForecast)
	panic(http.ListenAndServe(":8080", nil))
}

//Funksjon for å hente koordinatene og navn til det stedet brukeren ønsker værdata fra
func getForm(r *http.Request){
	r.ParseForm()
	//Sjekker at brukeren gjør et POST-kall til serveren
	if r.Method == "POST" {
		//Henter data fra form-tags fra HTMLen
		latLng = r.Form["kords"][0]
		stedsNavn:= r.Form["by"][0]
		//Legger
		by = Sted {
			Sted: stedsNavn,
		}

		weath = WeatherData{
			Sted: stedsNavn,
		}

		//Printer informasjonen hentet fra brukeren til serveren
		fmt.Println(latLng)
		fmt.Println(stedsNavn)
	}
}

func httpErrorHandling(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


//Enkel feilhåndtering
func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Formaterer den uformaterte urlen som inneholder koordinater fra brukeren
func latLngFormat(url string) string {
	first := strings.Replace(url, "(", "", -1)
	second := strings.Replace(first, ")", "", -1)
	lats := strings.Replace(second, " ", "", -1)
	return lats
}

//Funksjon for å hente og unmarshalle JSON
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