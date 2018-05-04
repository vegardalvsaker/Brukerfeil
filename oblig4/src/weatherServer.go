package main

import (
"net/http"
"log"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"strings"

	"fmt"
)

type WeatherData struct {
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

var weath WeatherData

//variable holding Latitude and Longitude
var latLng string

//API-url with custom API-key
const weatherUrlWithKey = "https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50"

func main() {
	//Startkoordinater for kartet på fremsiden
	weath = WeatherData {
		Latitude: 60.4720,
		Longitude: 8.4689,
	}
	server()
}

func index (w http.ResponseWriter, r *http.Request) {
	//Lager template av HTML-filene
	tmpl, err := template.ParseFiles("index.html", "gmapsTemplate.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	//Kjører HTML-templaten til writeren og mater den med interfacet
	if err := tmpl.Execute(w, weath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Sjekker at URLen har nok runes til at det kan være koordinater
	if len(r.URL.Path) > 12 {
		latLngUnformatted := r.URL.Path
		latLng = latLngFormat(latLngUnformatted)
	}
}

func runWeather (w http.ResponseWriter, r *http.Request) {
	//Setter sammen API-urlen sammen med koordinatene for å få en fungerende url
	fmt.Println(latLng)
	//Setter sammen API-urlen sammen med koordinatene
	strSlice := []string{weatherUrlWithKey, latLng}
	getUrl := strings.Join(strSlice, "/")

	getAndUnmarshal(getUrl, &weath)


	//EKSEMPEL:
	//Hvis temperaturen er over så så mye, kjør den relevante html-filen
	if fahrToCels(weath.Currently.Temperature) > 0 {
		t, _ := template.ParseFiles("weathData.html", "above25.html")
		t.Execute(w, weath)
	}


}




//Server-funksjon som håndterer paths
func server() {
	http.HandleFunc("/", index) // setting router rule
	http.HandleFunc("/forecast", runWeather)
	panic(http.ListenAndServe(":8080", nil))
}
//Enkel feilhåndtering
func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
//Formaterer den uformaterte urlen som inneholder koordinater fra brukeren
func latLngFormat(url string) string {
	first := strings.Replace(url, "/(", "", -1)
	second := strings.Replace(first, ")", "", -1)
	lats := strings.Replace(second, " ", "", -1)
	return lats
}
//Funksjon for å gjøre fahrenheit til celsius
func fahrToCels(temp float64) float64 {
	cels := temp - 32
	cels *= 5
	cels /= 9
	return cels
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

