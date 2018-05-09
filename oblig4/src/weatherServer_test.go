package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
	"strconv"
)

const darksky  = "https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50/58.8532585,5.732945500000028?lang=nb&units=si"


func TestIndex ( t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Error("could not create request %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(index)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", err, http.StatusOK)
	}

}

func TestGetAndUnmarshal (t *testing.T){
	var test WeatherData
	//Henter jsondataen i s
	res, err := http.Get(darksky)
	defer res.Body.Close()
	errorHandling(err)
	//Leser kroppen til jsondataen
	jsonBytes, err2 := ioutil.ReadAll(res.Body)
	errorHandling(err2)
	//Legger jsondaten inn i variabelen v
	errorHandling(json.Unmarshal(jsonBytes, &test))
	if test.Currently.Summary == "" {
		t.Error("Does not work")
	}
	fmt.Println(test)
}

//Tester at /forecast returnerer StatusKode OK: 200 hvis et api-kall til værdatabasen har blitt gjort i forkant
func TestRunForecast (t *testing.T) {
	kords := map[string]string{
		"58.8532585,5.732945500000028":         "Sandnes",
		"59.32932349999999,18.068580800000063": "Stockholm",
		"34.0522342,-118.2436849":              "Los Angeles",
		"-33.9248685,18.424055299999964":       "Cape Town",
		"37.566535,126.97796919999996":         "Seoul",
		"55.755826, 37.617299900000035":        "Moskva",
	}
	corrupt := []string{
		"askjhda k21237123()",
		"90120123983180120381083102",
		"(-1111,3333)",
		"Oslo",
		"===9((&%#'``#¤%agsdg",
		"",
		"569939,-12",  //10 characters
		"569939,-123", //11 characters
	}


	for kord, sted := range kords {

		t.Run(sted, func(t *testing.T) {
			rr := httptest.NewRecorder()
			//Endrer denne variabelen fordi runForecast henter koordinatene som ligger i denne globale variabelen
			latLng = kord

			req, err := http.NewRequest("GET", "localhost:8080/forecast", nil)
			if err != nil {
				t.Error("could not create request %v", err)
			}

			runForecast(rr, req)
			res := rr.Result()
			fmt.Println(res.StatusCode)
			if res.StatusCode != http.StatusOK {
				t.Errorf("Expected %v; got %v", http.StatusOK, res.StatusCode)
			}
		})
	}
	for i, kords := range corrupt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			rr := httptest.NewRecorder()
			//Endrer denne variabelen fordi runForecast henter koordinatene som ligger i denne globale variabelen
			latLng = kords

			req, err := http.NewRequest("GET", "localhost:8080/forecast", nil)
			if err != nil {
				t.Error("could not create request %v", err)
			}

			runForecast(rr, req)
			res := rr.Result()
			fmt.Println(res.StatusCode)
			if res.StatusCode != http.StatusBadRequest {
				t.Errorf("Expected %v; got %v", http.StatusOK, res.StatusCode)
				}
			})
		}

}

//Negativ test for runForecast-funksjonen
func TestRunForecastNoData (t *testing.T) {
	rr := httptest.NewRecorder()
	latLng = ""

	req, err := http.NewRequest("GET", "localhost:8080/forecast", nil)
	if err != nil {
		t.Error("could not create request %v", err)
	}

	runForecast(rr, req)
	res := rr.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code should be %v; not %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestLatLngFormat (t *testing.T) {
	t.Run("positiv", func(t *testing.T) {
		fmt := latLngFormat("(58.8532585, 5.732945500000028)")
		if strings.ContainsAny(fmt, "( )") {
			t.Errorf("latLngFormat greide ikke å formatere koordinatene. Resultatet ble %v", fmt)
		}
	})
	t.Run("negativ", func(t *testing.T){
		expected := "asdljhafkjfh91281028"
		fmt := latLngFormat(expected)
		if fmt != "asdljhafkjfh91281028" {
			t.Errorf("latLngFormat skulle returnere %v; fikk %v", expected, fmt)
		}
	})
	t.Run("negativ2", func(t *testing.T){
		expected := "(asdljhafkjf h91281028)"
		fmt := latLngFormat(expected)
		if fmt != "asdljhafkjfh91281028" {
			t.Errorf("latLngFormat skulle returnere %v; fikk %v", expected, fmt)
		}
	})
}
