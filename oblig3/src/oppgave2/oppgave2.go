package main

import (
	"net/http"
	//"github.com/gorilla/mux"
	"time"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strings"
	"log"
	"html/template"
)

type JsonStruct struct {
	Dato                string `json:"Dato"`
	Klokkeslett         string `json:"Klokkeslett"`
	Sted                string `json:"Sted"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	AntallLedigePlasser string `json:"Antall_ledige_plasser"`
	}

type JsonStruct2 struct {
	Entries []struct {
		Latitude    string `json:"latitudeka"`
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
		Name         string  `json:"Name"`
		ShortName    string  `json:"ShortName"`
		Latitude     float64 `json:"Latitude"`
		Longitude    float64 `json:"Longitude"`
		Zone1        int     `json:"Zone1"`
		Zone2        int     `json:"Zone2"`
		TransferTime int     `json:"TransferTime"`
		Transfer     int     `json:"Transfer"`
		BusStopType  int     `json:"BusStopType"`
		BusStopID    int     `json:"BusStopId"`
	} `json:"BusStops"`

}
var holdeplass BussHoldeplass
var fylker FylkesNummer
var m JsonStruct
var m2 []JsonStruct
var m3 JsonStruct2

func main() {

	http.HandleFunc("/", Hello1)
	http.HandleFunc("/2", Hello2)
	http.HandleFunc("/3", Hello3)
	http.HandleFunc("/4", Hello4)

	srvr := http.Server {
		Addr:         ":8080",
		ReadTimeout:  10* time.Second,
		WriteTimeout: 10* time.Second,
	}
	panic(srvr.ListenAndServe())
}

func Hello1 (w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	jsonBytes,_ := ioutil.ReadAll(res.Body)
	jsonString := string(jsonBytes)
	dec := json.NewDecoder(strings.NewReader(jsonString))

	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%v \n %v\n", m.Dato, m.Klokkeslett)
	}
	defer res.Body.Close()

	if err := json.Unmarshal(jsonBytes, &m2); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "%s \n ", m2[2].Sted)
}

func Hello2 (w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	jsonBytes, _ :=ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(jsonBytes, &m3); err != nil {
		log.Fatal(err)
	}

	b, _ := ioutil.ReadFile("src/oppgave2/index.html")
	streng := string(b)

	t := template.Must(template.New("").Parse(streng))
	if err := t.Execute(w, m3); err != nil {
		log.Fatal(err)
	}
}

func Hello3 (w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://register.geonorge.no/api/subregister/sosi-kodelister/kartverket/fylkesnummer-alle.json?")
	jsonBytes, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(jsonBytes, &fylker); err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadFile("src/oppgave2/fylker.html")
	streng := string(b)

	t := template.Must(template.New("").Parse(streng))
	if err := t.Execute(w, fylker); err != nil {
		log.Fatal(err)
	}
}

func Hello4 (w http.ResponseWriter, r *http.Request){
	res, _ := http.Get("http://sanntidsappservice-web-prod.azurewebsites.net/busstops?format=json")
	jsonBytes, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(jsonBytes, &holdeplass); err != nil {
		log.Fatal(err)
	}

	b, _ := ioutil.ReadFile("src/oppgave2/holdeplass.html")
	streng := string(b)

	t := template.Must(template.New("").Parse(streng))
	if err := t.Execute(w, holdeplass); err != nil {
		log.Fatal(err)
	}
}

func Hello5(w http.ResponseWriter, r *http.Request){

}
