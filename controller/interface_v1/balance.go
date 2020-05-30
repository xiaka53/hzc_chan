package interface_v1

import (
	"api/dao"
)

type apiBalance struct {
}

func GetBalance() *apiBalance {
	return new(apiBalance)
}

func (a *apiBalance) GetBalance(address, token string) (balance float64) {
	balanceData := dao.Balance{Address: address, Token: token}
	if err := (&balanceData).First(); err != nil {
		return
	}
	balance = float64(balanceData.Asset)
	return
}

func (a *apiBalance) GetBalanceAll(token string) (balance float64) {
	balanceData := dao.Balance{Token: token}
	if err := (&balanceData).Sum(); err != nil {
		return
	}
	balance = float64(balanceData.Asset)
	return
}
