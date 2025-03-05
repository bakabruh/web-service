package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "2983DNZ",
		Username:  "alex",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
}

func (m *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (m *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (m *mockDB) SetupDatabase() error {
	return nil
}
