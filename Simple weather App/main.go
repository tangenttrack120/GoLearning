package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	key := os.Getenv("WEATHER_API_KEY")
	unit := "metric"
	CallWeatherUrl := "http://api.openweathermap.org/data/2.5/weather?"
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a city name.")
		fmt.Println("Usage: go run main.go <city>")
		os.Exit(1) // Gracefully exit the program with an error status
	}
	City := os.Args[1] // check if there are command line arguments, Args[0] is the path or the name of the program
	GetWeatherURL := fmt.Sprintf("q=%s&appid=%s&units=%s", City, key, unit)

	// I feel this information is enough
	type Weather struct {
		City string `json:"name"`
		// 'sys' is a nested object in the JSON
		Sys  struct {
			Country string `json:"country"`
		} `json:"sys"`
		// 'main' is a nested object
		Main struct {
			Temperature float64 `json:"temp"`
		} `json:"main"`
		// 'weather' is an array of objects
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	}

	Response, err := http.Get(CallWeatherUrl + GetWeatherURL)
	if err != nil {
		panic(err)
	}
	defer Response.Body.Close()
	var weatherData Weather
	err = json.NewDecoder(Response.Body).Decode(&weatherData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The weather in %s, %s is %s with a temperature of %f\n", weatherData.City, weatherData.Sys.Country, weatherData.Weather[0].Description, weatherData.Main.Temperature)



}