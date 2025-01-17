package usecases

import (
	"errors"
	"testing"

	entities "github.com/felipepnascimento/challenge-bravo-flp/entities"
	"github.com/felipepnascimento/challenge-bravo-flp/mocks"
	"github.com/felipepnascimento/challenge-bravo-flp/models"
	"github.com/stretchr/testify/suite"
)

type enchangeUsecaseSuite struct {
	suite.Suite
	service *mocks.ExchangeRateService
	usecase ExchangeRateUsecase
}

func (suite *enchangeUsecaseSuite) SetupSuite() {
	service := new(mocks.ExchangeRateService)
	usecase := InitializeExchangeRateUsecase(service)
	suite.service = service
	suite.usecase = usecase
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRateWithNilFromCurrency() {
	toCurrency := models.Currency{
		Key:         "BLR",
		ExchangeApi: true,
	}
	_, err := suite.usecase.GetCurrencyRate(nil, &toCurrency)
	suite.Equal("from currency cannot be nil", err.Error())
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRateWithNilToCurrency() {
	fromCurrency := models.Currency{
		Key:         "USD",
		ExchangeApi: true,
	}

	_, err := suite.usecase.GetCurrencyRate(&fromCurrency, nil)
	suite.Equal("to currency cannot be nil", err.Error())
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRateWithError() {
	fromCurrency := models.Currency{
		Key:         "USD",
		ExchangeApi: true,
	}
	toCurrency := models.Currency{
		Key:         "BLR",
		ExchangeApi: true,
	}
	suite.service.On("GetLatestRate", fromCurrency.Key, toCurrency.Key).Return(nil, errors.New("Some generic error"))

	_, err := suite.usecase.GetCurrencyRate(&fromCurrency, &toCurrency)

	suite.Equal("Some generic error", err.Error())
	suite.service.AssertExpectations(suite.T())
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRateWithNotFoundRate() {
	fromCurrency := models.Currency{
		Key:         "NOT-EXISTS",
		ExchangeApi: true,
	}
	toCurrency := models.Currency{
		Key:         "NOT-EXISTS",
		ExchangeApi: true,
	}
	rate := float32(1.1)
	exchangeResult := entities.ExchangeResult{
		Success: true,
		Base:    "USD",
		Rates:   map[string]float32{"BRL": rate},
	}
	suite.service.On("GetLatestRate", fromCurrency.Key, toCurrency.Key).Return(&exchangeResult, nil)

	_, err := suite.usecase.GetCurrencyRate(&fromCurrency, &toCurrency)

	suite.Equal("Can not find target currency rate", err.Error())
	suite.service.AssertExpectations(suite.T())
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRate() {
	fromCurrency := models.Currency{
		Key:         "USD",
		ExchangeApi: true,
	}
	toCurrency := models.Currency{
		Key:         "BRL",
		ExchangeApi: true,
	}
	rate := float32(1.1)
	exchangeResult := entities.ExchangeResult{
		Success: true,
		Base:    "USD",
		Rates:   map[string]float32{"BRL": rate},
	}
	suite.service.On("GetLatestRate", fromCurrency.Key, toCurrency.Key).Return(&exchangeResult, nil)

	result, err := suite.usecase.GetCurrencyRate(&fromCurrency, &toCurrency)

	suite.NoError(err)
	suite.Equal(rate, result)
	suite.service.AssertExpectations(suite.T())
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRateWithHURBCurrencyAndInvalidToCurrency() {
	fromCurrency := models.Currency{
		Key:         "HURB",
		ExchangeApi: false,
	}
	toCurrency := models.Currency{
		Key:         "BRL",
		ExchangeApi: false,
	}

	_, err := suite.usecase.GetCurrencyRate(&fromCurrency, &toCurrency)

	suite.Equal("To currency must be converted by ExchangeApi", err.Error())
	suite.service.AssertExpectations(suite.T())
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRateWithHURBCurrencyAndInvalidFromCurrency() {
	fromCurrency := models.Currency{
		Key:         "HURB",
		ExchangeApi: false,
	}
	toCurrency := models.Currency{
		Key:         "BRL",
		ExchangeApi: true,
	}

	_, err := suite.usecase.GetCurrencyRate(&fromCurrency, &toCurrency)

	suite.Equal("from currency must have custom currency and custom amount", err.Error())
	suite.service.AssertExpectations(suite.T())
}

func (suite *enchangeUsecaseSuite) TestGetCurrencyRateWithHURBCurrencyAndValidToCurrency() {
	customAmount := float32(54)
	fromCurrency := models.Currency{
		Key:            "HURB",
		ExchangeApi:    false,
		CustomAmount:   customAmount,
		CustomCurrency: "BRL",
	}
	toCurrency := models.Currency{
		Key:         "USD",
		ExchangeApi: true,
	}

	rate := float32(2)
	exchangeResult := entities.ExchangeResult{
		Success: true,
		Base:    "BRL",
		Rates:   map[string]float32{"USD": rate},
	}
	suite.service.On("GetLatestRate", fromCurrency.CustomCurrency, toCurrency.Key).Return(&exchangeResult, nil)

	result, err := suite.usecase.GetCurrencyRate(&fromCurrency, &toCurrency)

	suite.NoError(err)
	suite.Equal((rate * customAmount), result)
	suite.service.AssertExpectations(suite.T())
}

func TestExchangeRateUsecase(t *testing.T) {
	suite.Run(t, new(enchangeUsecaseSuite))
}
