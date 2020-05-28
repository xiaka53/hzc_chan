package interface_v1

import (
	"api/dao"
	"math"
)

type apiBalance struct {
}

func GetBalance() *apiBalance {
	return new(apiBalance)
}

func (a *apiAccount) GetBalance(address, token string) (balance float64) {
	balanceData := dao.Balance{Address: address, Token: token}
	if err := (&balanceData).First(); err != nil {
		return
	}
	balance = float64(balanceData.Asset.Int64()) / math.Pow(10, 18)
	return
}

func (a *apiAccount) GetBalanceAll(token string) (balance float64) {
	balanceData := dao.Balance{Token: token}
	if err := (&balanceData).Sum(); err != nil {
		return
	}
	balance = float64(balanceData.Asset.Int64()) / math.Pow(10, 18)
	return
}