package web

import "github.com/flaviotsx/temperature-zipcode/internal/entity"

type GetTemperatureByZipCodeUseCaseInterface interface {
	Execute(zipcode string) (*entity.GetTemperatureByZipcodeResponse, *entity.ErrorResponse)
}
