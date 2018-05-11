package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Sted struct {
	Sted string
}

type WeatherData struct {
	Sted      string
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Currently struct {
		Summary             string  `json:"summary"`
		PrecipIntensity     float64    `json:"precipIntensity"`
		PrecipProbability   float64     `json:"precipProbability"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		Humidity            float64 `json:"humidity"`
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
	sigintHandtering()
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
		if len(latLng) <= 12 || !isCorrectRune(latLng) {
			w.WriteHeader(http.StatusBadRequest)
			t, err := template.ParseFiles("feil.html")
			httpErrorHandling(w, err)
			t.Execute(w, nil)
		} else {
			//Setter sammen API-urlen sammen med koordinatene
			url := joinUrlAndCoord()

			//Server prints for debugging
			fmt.Println(r.URL.Path)
			fmt.Println(latLng)
			fmt.Println("Weather data collected from: ", url)

			if !getAndUnmarshal(url, &weath, w) {
				w.WriteHeader(http.StatusBadRequest)
				t, err := template.ParseFiles("feil2.html")
				httpErrorHandling(w, err)
				t.Execute(w, nil)
			} else {
				//Funksjon for å velge ut "skreddersydde" meldinger til sluttbrukeren
				executeTilbakemelding(w)
			}
		}
	}

//Funksjon for å hente koordinatene og navn til det stedet brukeren ønsker værdata fra
func getForm(r *http.Request){
	r.ParseForm()
	//Sjekker at brukeren gjør et POST-request til serveren
	if r.Method == "POST" {
		//Henter data fra form-tags fra HTMLen
		//"kords" og "by" er name attributter i <input> tags i index.html
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
		fmt.Println("Bruker input er nå blitt hentet:")
		fmt.Println("Bruker sin adresse:",r.RemoteAddr)
		fmt.Println("Koordinater:",latLng)
		fmt.Println("Sted:", stedsNavn)
		fmt.Println("----------------------------------------------------")
		fmt.Println()
	} else {
		fmt.Printf("HTTP forespørsel er '%v', venter på 'POST' før brukerinput kan tas imot\n", r.Method)
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
//Setter sammen koordinatene til en fungerende URL til API-kall
func joinUrlAndCoord() string {
	latAndLang := latLng + langAndUnits
	strSlice := []string{weatherUrlWithKey, latAndLang}
	return strings.Join(strSlice, "/")
}

//Funksjon for å hente og unmarshalle JSON
func getAndUnmarshal(s string, v interface{}, w http.ResponseWriter) bool {
	//Henter jsondataen i s
	res, err := http.Get(s)
	defer res.Body.Close()
	httpErrorHandling(w, err)
	//Leser kroppen til jsondataen
	jsonBytes, err2 := ioutil.ReadAll(res.Body)
	errorHandling(err2)
	if string(jsonBytes) == `{"code":400,"error":"The given location is invalid."}` || string(jsonBytes) ==`{"code":400,"error":"The given location (or time) is invalid."}` || string(jsonBytes) == "Bad Request"{
		return false
	} else {
		//Legger jsondaten inn i variabelen
		httpErrorHandling(w, json.Unmarshal(jsonBytes, &v))
		return true
	}
}
//Funksjon for å finne ut hvilken tilbakemelding brukeren skal få
func executeTilbakemelding(w http.ResponseWriter) {
	//Hvis det ikke blåser, gjør nedbørsvurderinger
	if weath.Currently.WindSpeed < 8 {
		//Hvis det ikke regner, gjør temperaturvurderinger
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

//En ikke så veldig pen funksjon for å sjekke at en koordinat-string kun har tillatte tegn.
func isCorrectRune(s string) bool {
	for _, r := range s {
		if !((r == '0') || (r == '1') || (r == '2') || (r == '3') || (r == '4') || (r == '5') || (r == '6') || (r == '7') || (r == '8') || (r == '9') || (r == '-') || ( r == '.') || ( r == ',')) {
			return false
		}
	}
	return true
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
//Håndtering av SIG-INT signal for å skru av tjeneren på en trygg måte.
func sigintHandtering() {
	chsig := make(chan os.Signal, 1)
	signal.Notify(chsig, syscall.SIGINT)
	go func() {
		for sig := range chsig {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Sigint received. Shutting down safely")
				time.Sleep(3e+9)
				os.Exit(0)
			}
		}
	}()
}