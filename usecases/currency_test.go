package usecases

import (
	"errors"
	"testing"

	"github.com/felipepnascimento/challenge-bravo-flp/mocks"
	"github.com/felipepnascimento/challenge-bravo-flp/models"

	"github.com/stretchr/testify/suite"
)

type currencyUsecaseSuite struct {
	suite.Suite
	repository *mocks.CurrencyRepository
	usecase    CurrencyUsecase
}

func (suite *currencyUsecaseSuite) SetupSuite() {
	repository := new(mocks.CurrencyRepository)
	usecase := InitializeCurrencyUsecase(repository)
	suite.repository = repository
	suite.usecase = usecase
}

func (suite *currencyUsecaseSuite) TestCreateCurrencyWithNilValues() {
	err := suite.usecase.CreateCurrency(nil)
	suite.Equal("currency is nil", err.Error())
}

func (suite *currencyUsecaseSuite) TestCreateCurrencyWithInvalidKey() {
	currency := models.Currency{
		Description:   "description",
	}

	err := suite.usecase.CreateCurrency(&currency)
	suite.Equal("key and description cannot be empty", err.Error())
}

func (suite *currencyUsecaseSuite) TestCreateCurrencyWithInvalidDescription() {
	currency := models.Currency{
		Key:           "key",
	}

	err := suite.usecase.CreateCurrency(&currency)
	suite.Equal("key and description cannot be empty", err.Error())
}

func (suite *currencyUsecaseSuite) TestCreateCurrency() {
	currency := models.Currency{
		Key:           "key",
		Description:   "description",
	}
	suite.repository.On("CreateCurrency", &currency).Return(nil)

	err := suite.usecase.CreateCurrency(&currency)
	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *currencyUsecaseSuite) TestGetAllCurrencies() {
	currencies := []models.Currency{
		{
			Key:           "key",
			Description:   "description",
		},
		{
			Key:           "key",
			Description:   "description",
		},
	}
	suite.repository.On("GetAllCurrencies").Return(&currencies, nil)

	result, err := suite.usecase.GetAllCurrencies()
	suite.NoError(err)
	suite.Equal(len(currencies), len(*result))
	suite.Equal(currencies, *result)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *currencyUsecaseSuite) TestGetCurrencyByIdNotFound() {
	id := 1
	suite.repository.On("GetCurrencyById", id).Return(nil, errors.New("currency is not found"))
	result, err := suite.usecase.GetCurrencyById(id)

	suite.Nil(result)
	suite.Equal("currency is not found", err.Error())
	suite.repository.AssertExpectations(suite.T())
}

func (suite *currencyUsecaseSuite) TestGetCurrencyById() {
	id := 2
	currency := models.Currency{
		Key:           "key",
		Description:   "description",
	}

	suite.repository.On("GetCurrencyById", id).Return(&currency, nil)

	result, err := suite.usecase.GetCurrencyById(id)
	suite.Nil(err)
	suite.Equal(*result, currency)
}

func (suite *currencyUsecaseSuite) TestDeleteCurrency() {
	id := 1
	suite.repository.On("DeleteCurrency", id).Return(nil)

	err := suite.usecase.DeleteCurrency(id)
	suite.Nil(err)
}

func TestCurrencyUsecase(t *testing.T) {
	suite.Run(t, new(currencyUsecaseSuite))
}
