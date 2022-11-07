package repositories

import (
	"testing"

	"github.com/felipepnascimento/challenge-bravo-flp/config"
	"github.com/felipepnascimento/challenge-bravo-flp/entities"
	"github.com/felipepnascimento/challenge-bravo-flp/utils"

	"github.com/stretchr/testify/suite"
)

type currencyRepositorySuite struct {
	suite.Suite
	repository      CurrencyRepository
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *currencyRepositorySuite) SetupSuite() {

	configs := config.GetConfig()
	db := config.ConnectDB(configs)
	repository := InitializeCurrencyRepository(db)

	suite.repository = repository
	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
}

func (suite *currencyRepositorySuite) TearDownTest() {
	defer suite.cleanupExecutor.TruncateTable([]string{"currencies"})
}

func (suite *currencyRepositorySuite) TestCreateCurrency() {
	currency := entities.Currency{
		Key:           "key",
		Description:   "description",
		QuotationType: "quotationType",
	}

	err := suite.repository.CreateCurrency(&currency)
	suite.NoError(err)
}

func (suite *currencyRepositorySuite) TestCreateCurrencyWithDuplicatedKey() {
	currency := entities.Currency{
		Key:           "key",
		Description:   "description",
		QuotationType: "quotationType",
	}

	err := suite.repository.CreateCurrency(&currency)
	suite.NoError(err)

	err = suite.repository.CreateCurrency(&currency)
	suite.Equal(err.Error(), "ERROR: duplicate key value violates unique constraint \"currencies_pkey\" (SQLSTATE 23505)")
}

func (suite *currencyRepositorySuite) TestGetAllCurrencies() {
	currency := entities.Currency{
		Key:           "key",
		Description:   "description",
		QuotationType: "quotationType",
	}

	currencyOne := entities.Currency{
		Key:           "key_one",
		Description:   "description",
		QuotationType: "quotationType",
	}

	err := suite.repository.CreateCurrency(&currency)
	suite.NoError(err)
	err = suite.repository.CreateCurrency(&currencyOne)
	suite.NoError(err)

	currencies, err := suite.repository.GetAllCurrencies()
	suite.NoError(err)
	suite.Equal(len(*currencies), 2, "insert 2 records before get all data, so it should contain three currencies")
}

func (suite *currencyRepositorySuite) TestGetCurrencyByIDNotFound() {
	id := 1

	_, err := suite.repository.GetCurrencyByID(id)
	suite.Error(err)
	suite.Equal(err.Error(), "record not found")
}

func (suite *currencyRepositorySuite) TestGetCurrencyByID() {
	id := 1
	currency := entities.Currency{
		Key:           "key",
		Description:   "description",
		QuotationType: "quotationType",
	}

	err := suite.repository.CreateCurrency(&currency)
	suite.NoError(err)

	result, err := suite.repository.GetCurrencyByID(id)
	suite.NoError(err)
	suite.Equal(currency.Key, (*result).Key)
	suite.Equal(currency.Description, (*result).Description)
	suite.Equal(currency.QuotationType, (*result).QuotationType)
}

func (suite *currencyRepositorySuite) TestDeleteCurrency() {
	id := 1
	currency := entities.Currency{
		Key:           "key",
		Description:   "description",
		QuotationType: "quotationType",
	}

	err := suite.repository.CreateCurrency(&currency)
	suite.NoError(err)

	_, err = suite.repository.GetCurrencyByID(id)
	suite.NoError(err)

	err = suite.repository.DeleteCurrency(id)
	suite.NoError(err)

	_, err = suite.repository.GetCurrencyByID(id)
	suite.Equal(err.Error(), "record not found")
}

func TestCurrencyRepository(t *testing.T) {
	suite.Run(t, new(currencyRepositorySuite))
}