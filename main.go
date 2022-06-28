package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
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

		result := WeatherData{
			Water:  water,
			Wind:   wind,
			Status: weatherStatus,
		}

		tmpl, err := template.ParseFiles("views/weather_status.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, result)
		return

	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
