package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"anika": {
		AuthToken: "1234",
		Username:  "anika",
	},
	"shoccho": {
		AuthToken: "abcd",
		Username:  "Shoccho",
	},
}
var mockCoinDetails = map[string]CoinDetails{
	"anika": {
		Coins:    1234,
		Username: "anika",
	},
	"shoccho": {
		Coins:    999,
		Username: "Shoccho",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData //todo: check why pointers
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData //todo: check why pointers

}

func (d *mockDB) SetupDatabase() error {
	return nil
}

var MockDB = mockDB{}
