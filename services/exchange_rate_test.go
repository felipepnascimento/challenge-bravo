package services

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/felipepnascimento/challenge-bravo-flp/entities"
	"github.com/felipepnascimento/challenge-bravo-flp/mocks"
	"github.com/stretchr/testify/suite"
)

type exchangeRateServiceSuite struct {
	suite.Suite
	httpClient *mocks.HTTPClient
	service    ExchangeRateService
}

func (suite *exchangeRateServiceSuite) SetupSuite() {
	httpClient := new(mocks.HTTPClient)
	service := InitializeExchangeRateService(httpClient)
	suite.httpClient = httpClient
	suite.service = service
}

func (suite *exchangeRateServiceSuite) TestGetLatestRateWitError() {
	toCurrency := ""
	req := MockGet(toCurrency)
	resp := http.Response{}

	suite.httpClient.On("Do", req).Return(&resp, errors.New("Some generic error"))

	_, err := suite.service.GetLatestRate(toCurrency)
	suite.Equal(err.Error(), "An error occurred to makes the request")
	suite.httpClient.AssertExpectations(suite.T())
}

func (suite *exchangeRateServiceSuite) TestGetLatestRate() {
	toCurrency := "BRL"
	req := MockGet(toCurrency)
	respBody := ioutil.NopCloser(bytes.NewBufferString(`{"Success": true, "Base": "USD", "Date": "2022-11-08", "Rates": {"BRL": 1.1}}`))
	resp := http.Response{
		Body: respBody,
	}

	suite.httpClient.On("Do", req).Return(&resp, nil)

	result, err := suite.service.GetLatestRate(toCurrency)

	expectedResult := entities.ExchangeResult{
		Success: true,
		Base:    "USD",
		Date:    "2022-11-08",
		Rates:   map[string]float32{"BRL": 1.1},
	}

	suite.NoError(err)
	suite.Equal(result, &expectedResult)
	suite.httpClient.AssertExpectations(suite.T())
}

func MockGet(toCurrency string) *http.Request {
	BASE_CURRENCY := "USD"
	BASE_URL := "https://api.exchangerate.host/latest?base=%s&symbols=%s"

	baseUrl := fmt.Sprintf(BASE_URL, BASE_CURRENCY, toCurrency)
	req, _ := http.NewRequest("GET", baseUrl, nil)

	return req
}

func TestExchangeService(t *testing.T) {
	suite.Run(t, new(exchangeRateServiceSuite))
}