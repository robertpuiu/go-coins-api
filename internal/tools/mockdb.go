package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"robert": {
		AuthToken: "944PUIU",
		Username:  "robert",
	},
	"vlad": {
		AuthToken: "001STEF",
		Username:  "vlad",
	},
	"alex": {
		AuthToken: "004LEAF",
		Username:  "alex",
	},
}

var mockCoinsDetails = map[string]CoinDetails{
	"robert": {
		Coins:    9441,
		Username: "robert",
	},
	"vlad": {
		Coins:    861,
		Username: "vlad",
	},
	"alex": {
		Coins:    1285,
		Username: "alex",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {

	time.Sleep(time.Second * 1) // simulate DB call

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinsDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
