package main

import (
	"net/http"
	"github.com/gorilla/mux"
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
var m2 []JsonStruct
var m3 JsonStruct2

var myClient = &http.Client{Timeout:10 * time.Second}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", Hello1)
	router.HandleFunc("/2", Hello2)

	srvr := http.Server {
		Addr:         ":8080",
		Handler:      router,
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

	//t := json.Token(jsonBytes)

	if err := json.Unmarshal(jsonBytes, &m2); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "%s \n ", m2[2].Sted)

		//w.Header().Set("Content-Type", "application/json")
		//w.Write(js)

}

func Hello2 (w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	jsonBytes, _ :=ioutil.ReadAll(res.Body)
	//jsonString := string(jsonBytes)
	//jsonReader := strings.NewReader(jsonString)
	//dec := json.NewDecoder(jsonReader)

	if err := json.Unmarshal(jsonBytes, &m3); err != nil {
		log.Fatal(err)
	}


	b, _ := ioutil.ReadFile("oblig3/src/oppgave2/index.html")
	//n := len(b)
	streng := string(b)

	t := template.Must(template.New("").Parse(streng))
	//n := len(m3.Entries)
	if err := t.Execute(w, m3); err != nil {
		log.Fatal(err)
	}

	//fmt.Fprintf(w, "Alle navn i dette Json-datasettet: \n\n")
	//for i := 0; i < len(m3.Entries); i++ {
	//	fmt.Fprintf(w, "%v Her kan du sortere plast. J for ja, N for nei %v\n", m3.Entries[i].Navn, m3.Entries[i].Plast)
	//}
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}