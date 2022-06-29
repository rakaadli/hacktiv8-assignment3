package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

const (
	PORT = ":3000"
	MIN  = 1
	MAX  = 100
)

type WeatherData struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/weather", getWeather)
	fmt.Println("please open http://localhost:3000/weather")
	http.ListenAndServe(PORT, nil)
}

func getWeather(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		//Read Data
		file, _ := ioutil.ReadFile("template/data.json")
		hasil := WeatherData{}
		_ = json.Unmarshal(file, &hasil)
		//fmt.Println(hasil)

		water := rand.Intn(MAX-MIN) + MIN
		wind := rand.Intn(MAX-MIN) + MIN
		weatherStatus := "STATUS DEFAULT"

		if water < 5 {
			weatherStatus = "AMAN"
		} else if water >= 6 && water == 8 {
			weatherStatus = "SIAGA"
		} else if water > 8 {
			weatherStatus = "BAHAYA"
		} else if wind < 6 {
			weatherStatus = "AMAN"
		} else if wind >= 7 && wind == 15 {
			weatherStatus = "SIAGA"
		} else if wind > 15 {
			weatherStatus = "BAHAYA"
		}

		updatedData := WeatherData{
			Water:  water,
			Wind:   wind,
			Status: weatherStatus,
		}

		//update data
		json, _ := json.Marshal(&updatedData)

		// fmt.Println(string(json))

		//write data
		_ = ioutil.WriteFile("template/data.json", json, os.ModePerm)

		tmpl, err := template.ParseFiles("views/weather_status.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, updatedData)
		return

	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
