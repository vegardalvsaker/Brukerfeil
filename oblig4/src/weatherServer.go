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
	Sted      string
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

//Server-funksjon som håndterer paths
func server() {
	http.HandleFunc("/", index)
	http.HandleFunc("/forecast", runForecast)
	panic(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Webserver for værdata startet!")
	fmt.Println("Lytter for http requests på localhost:8080")
	server()
}

func index (w http.ResponseWriter, r *http.Request) {
	//Resetter by-variabelen for å fjerne linken på framsiden hvis du velger å gå tilbake til index fra /forecast.
	by = Sted {
		Sted: "Søk etter sted...",
	}
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
	latLng = "===9((&%#'``#¤%agsdg"
	if len(latLng) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		t, err := template.ParseFiles("feil.html")
		httpErrorHandling(w, err)
		t.Execute(w, nil)

		//http.Error(w,"Du må søke etter et sted før du kan hente værmeldingen! \nGå tilbake for å gjøre et søk", http.StatusBadRequest)

	} else {
		//Setter sammen API-urlen sammen med koordinatene
		url := joinUrlAndCoord()

		//Server prints for debugging
		fmt.Println(r.URL.Path)
		fmt.Println(latLng)
		fmt.Println("Weather data collected from: ", url)

		getAndUnmarshal(url, &weath, w)

		//Funksjon for å velge ut "skreddersydde" meldinger til sluttbrukeren
		executeTilbakemelding(w)

	}


}

func joinUrlAndCoord() string {
	latAndLang := latLng + langAndUnits
	strSlice := []string{weatherUrlWithKey, latAndLang}
	return strings.Join(strSlice, "/")
}

//Funksjon for å hente koordinatene og navn til det stedet brukeren ønsker værdata fra
func getForm(r *http.Request){
	r.ParseForm()
	//Sjekker at brukeren gjør et POST-request til serveren
	if r.Method == "POST" {
		//Henter data fra form-tags fra HTMLen
		latLng = r.Form["kords"][0]
		stedsNavn:= r.Form["by"][0]
		//Legger stedet som blir søkt på i to forskjellige structs
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
	}
}

//Enkel feilhåndtering
func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Fjerner paranteser fra koordinatene slik at de skal passe inn i darksky-urlen
func latLngFormat(s string) string {
	if strings.ContainsAny(s, "()") {
		first := strings.Replace(s, "(", "", -1)
		second := strings.Replace(first, ")", "", -1)
		lats := strings.Replace(second, " ", "", -1)
		return lats
	} else {
		fmt.Errorf("Koordinatene mangler paranteser som pre- og suffix; den så slik ut %v", s)
		return s
	}
}

//Funksjon for å hente og unmarshalle JSON
func getAndUnmarshal(s string, v interface{}, w http.ResponseWriter) {
	//Henter jsondataen i s
	res, err := http.Get(s)
	defer res.Body.Close()
	httpErrorHandling(w, err)
	//Leser kroppen til jsondataen
	jsonBytes, err2 := ioutil.ReadAll(res.Body)
	errorHandling(err2)
	//Legger jsondaten inn i variabelen v
	httpErrorHandling(w, json.Unmarshal(jsonBytes, &v))
}

func executeTilbakemelding(w http.ResponseWriter) {
	//Hvis det ikke regner, gjør temperaturvurderinger
	if weath.Currently.WindSpeed < 8 {
		if weath.Currently.PrecipIntensity < 1 {
			if weath.Currently.Temperature >= 16 {
				t, err := template.ParseFiles("forecast.html", "msg/above20.html")
				httpErrorHandling(w, err)
				t.Execute(w, weath)
			} else if weath.Currently.Temperature < 16 && weath.Currently.Temperature > 10 {
				t, err := template.ParseFiles("forecast.html", "msg/below20.html")
				httpErrorHandling(w, err)
				t.Execute(w, weath)
			} else if weath.Currently.Temperature <= 0 {
				t, err := template.ParseFiles("forecast.html", "msg/below0.html")
				httpErrorHandling(w, err)
				t.Execute(w, weath)
			} else {
				t, err := template.ParseFiles("forecast.html", "msg/above0.html")
				httpErrorHandling(w, err)
				t.Execute(w, weath)
			}
		} else {
			if weath.Currently.PrecipIntensity > 1 && weath.Currently.PrecipIntensity < 6 {
				t, err := template.ParseFiles("forecast.html", "msg/above1mm.html")
				httpErrorHandling(w, err)
				t.Execute(w, weath)
			} else if weath.Currently.PrecipIntensity > 6 {
				t, err := template.ParseFiles("forecast.html", "msg/above6mm.html")
				httpErrorHandling(w, err)
				t.Execute(w, weath)
			} else {
				t, err := template.ParseFiles("forecast.html", "msg/below0.html")
				httpErrorHandling(w, err)
				t.Execute(w, weath)
			}
		}
	} else if weath.Currently.WindSpeed >= 8 && weath.Currently.WindSpeed < 21 {
		t, err := template.ParseFiles("forecast.html", "msg/kuling.html")
		httpErrorHandling(w, err)
		t.Execute(w, weath)
	} else if weath.Currently.WindSpeed >= 21 && weath.Currently.WindSpeed < 32.7 {
		t, err := template.ParseFiles("forecast.html", "msg/storm.html")
		httpErrorHandling(w, err)
		t.Execute(w, weath)
	} else {
		t, err := template.ParseFiles("forecast.html", "msg/orkan.html")
		httpErrorHandling(w, err)
		t.Execute(w, weath)
	}
}