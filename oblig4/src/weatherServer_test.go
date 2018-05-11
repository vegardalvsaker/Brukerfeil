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
	t.Run("GET", func(t *testing.T) {
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
	})
}

//Denne testen antar at den url-en den får er en fungerende url. Dette blir tatt hånd om før getAndUnmarshall blir kalt.
func TestGetAndUnmarshal (t *testing.T){
	t.Run("positiv", func(t *testing.T) {
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
			t.Error("The function doesn't work properly. The properties are emtpy.")
		}
		fmt.Println(test)
	})
	ødelagteLinker := []string{
		"https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50/acscasda2",
		"https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50/1234567894561285",
		"https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50/(58.8532432,5.409237492)",
		"https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50/xxxxzzzzcczccacaca",
		//"https://api.darksky.net/forecast/a529911b60d81bab2c791732ad9ddf50/¤#&%//%#", //Denne inputen kræsjer testen. Den hadde dog aldri kommet seg så langt at den ville kalt getAndUnmarshal()
	}
	for i, s := range ødelagteLinker{
		navn := "negativ" + strconv.Itoa(i)
		t.Run(navn, func(t *testing.T) {
			var test WeatherData
			rr := httptest.NewRecorder()
			if getAndUnmarshal(s, test, rr){
				t.Errorf("Should've returned false; not true")
			}

		})
	}
}

//Tester at /forecast returnerer StatusKode OK: 200 hvis et api-kall til værdatabasen har blitt gjort i forkant
func TestRunForecast (t *testing.T) {
	kords := map[string]string{
		"58.8532585,5.732945500000028":         "Sandnes",
		"59.32932349999999,18.068580800000063": "Stockholm",
		"34.0522342,-118.2436849":              "Los Angeles",
		"-33.9248685,18.424055299999964":       "Cape Town",
		"37.566535,126.97796919999996":         "Seoul",
		"55.755826,37.617299900000035":        "Moskva",
	}
	corrupt := []string{
		"askjhda k21237123()",
		"(-1111,3333)",
		"Oslo",
		"===9((&%#'``#¤%agsdg",
		"",
		"569939,-12",  //10 characters
		"569939,-123", //11 characters
		"569939,-1234", //12 characters
		"1616618816516851",
	}

	//Positiv testing
	for kord, sted := range kords {
		t.Run(sted, func(t *testing.T) {
			rr := httptest.NewRecorder()
			//Endrer denne variabelen fordi runForecast henter koordinatene som ligger i denne globale variabelen
			latLng = kord
			fmt.Println(latLng)

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
	//Negativ testing
	//Tester for koordinater som ikke er som de skal, som kun kan skje hvis noen endrer kildekoden i index.html
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
				t.Errorf("Expected %v; got %v", http.StatusBadRequest, res.StatusCode)
				}
			})
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

func TestJoinUrlAndCoord(t *testing.T) {
	kords := map[string]string{
		"58.8532585,5.732945500000028":         "Sandnes",
		"59.32932349999999,18.068580800000063": "Stockholm",
		"34.0522342,-118.2436849":              "Los Angeles",
		"-33.9248685,18.424055299999964":       "Cape Town",
		"37.566535,126.97796919999996":         "Seoul",
		"55.755826,37.617299900000035":        "Moskva",
	}
	corrupt := []string{
		"askjhda k21237123()",
		"(-1111,3333)",
		"Oslo",
		"===9((&%#'` `#¤%agsdg",
		"",
		"569939, -12",  //10 characters
		"569939,-123", //11 characters
		"569939,-1234", //12 characters
		"1616618816516851",
	}
	for i, s := range kords {
		navn := "positiv" + "/" + s
		t.Run(navn, func(t *testing.T) {
			latLng = i
			shouldBeThree := 0
			res := joinUrlAndCoord()
			if strings.Contains(res, weatherUrlWithKey) {
				shouldBeThree++
			}
			if strings.Contains(res, langAndUnits) {
				shouldBeThree++
			}
			if strings.Contains(res, i) {
				shouldBeThree++
			}
			if shouldBeThree != 3 {
				t.Errorf("joinUrlAndCoord joinet ikke URLen som ønsket; slik ble den %v \n", res)
			}
		})
	}
	for i, s := range corrupt {
			navn := "negativ"+ strconv.Itoa(i)
			t.Run(navn, func(t *testing.T) {
				latLng = s
				shouldBeThree := 0
				res := joinUrlAndCoord()
				if strings.Contains(res, weatherUrlWithKey) {
					shouldBeThree++
				}
				if strings.Contains(res, langAndUnits) {
					shouldBeThree++
				}
				if strings.Contains(res, s) {
					shouldBeThree++
				}
				if shouldBeThree != 3 {
					t.Errorf("joinUrlAndCoord joinet ikke URLen som ønsket; slik ble den %v \n", res)
				}
			})
	}
}


/*
//Prøvde å lage en test for funksjonen getForm, men sliter med å passe en body til requesten
//Mer om dette i Systemarkitekturen under avsnittet om testing
func TestGetForm (t *testing.T) {
	t.Run("negativ", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "localhost:8080/", nil)
		getForm(req)
		//Output:
		//HTTP request is GET, waiting until it's POST before getting user input
	})
	t.Run("positiv", func(t *testing.T) {
		bs, _ := ioutil.ReadFile("test.html")
		reader := bytes.NewReader(bs)
		req, _ := http.NewRequest("POST", "localhost:8080", reader )


		getForm(req)
		fmt.Println(weath.Sted)
		fmt.Println(by.Sted)
		fmt.Println(latLng)
		if weath.Sted != "Sandnes" {
			t.Errorf("weath.Sted should be 'Sandnes'; not %v", weath.Sted)
			}
		if by.Sted != "Sandnes" {
			t.Errorf("by.Sted should be 'Sandnes'; not %v", by.Sted)
		}
		if latLng != "58.8532585,5.732945500000028" {
				t.Errorf("latLng should be 58.8532585,5.732945500000028; not %v", latLng)
					}
		})

}*/