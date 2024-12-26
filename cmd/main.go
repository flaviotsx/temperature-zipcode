package main

import (
	"log"
	"net/http"

	"github.com/flaviotsx/temperature-zipcode/internal/repository"
	"github.com/flaviotsx/temperature-zipcode/internal/usecase"
	"github.com/flaviotsx/temperature-zipcode/internal/web"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", string(err.Error()))
	}

	zipcodeRepository := repository.NewZipCodeRepository()
	weatherApiRepository := repository.NewWeatherApiRepository()
	getTemperatureByZipCodeUseCase := usecase.NewGetTemperatureByZipcodeUseCase(zipcodeRepository, weatherApiRepository)
	getTemperatureByZipCodeHandler := web.NewGetTemperatureByZipCodeHandler(getTemperatureByZipCodeUseCase)

	r := mux.NewRouter()

	r.HandleFunc("/{zipcode}/temperature", getTemperatureByZipCodeHandler.GetTemperatureByZipcodeHandler)
	http.ListenAndServe(":8080", r)
}
