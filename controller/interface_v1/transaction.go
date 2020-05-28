package interface_v1

import (
	"api/controller/task"
	"api/dao"
	"math"
	"math/big"
)

type apiTransaction struct {
}

func GetTransaction() *apiTransaction {
	return new(apiTransaction)
}

func (a *apiTransaction) SendTransaction(from, to, date, input string, value, gas float64) {
	var (
		fromAddress dao.Balance
		toAddress   dao.Balance
		hash        dao.Hash
	)
	fromAddress.Address = from
	if (&fromAddress).RecordNotFound() {
		return
	}
	toAddress.Address = to
	if (&toAddress).RecordNotFound() {
		return
	}
	if len(date) > 0 {
		if (&dao.Token{Token: date, Status: 1}).RecordNotFound() {
			return
		}
		hash.Index = big.NewInt(int64(value * math.Pow(10, 18)))
	} else {
		hash.Value = big.NewInt(int64(value * math.Pow(10, 18)))
	}
	hash = dao.Hash{
		From:   from,
		To:     to + date,
		Gas:    big.NewInt(int64(gas * math.Pow(10, 18))),
		Input:  input,
		Status: 0,
	}
	task.GetHash().Add(hash)
	return
}

func (a *apiTransaction) GetTransactionByHash(hx string) (hash dao.Hash) {
	hash.Hash = hx
	_ = (&hash).First()
	return
}

func (a *apiTransaction) GetTransactionByBlock(block int) (data []dao.Hash) {
	data = (&dao.Hash{BlockNumber: block}).Find()
	return
}

func (a *apiTransaction) GetTransactionByAddress(address string) (data []dao.Hash) {
	data = (&dao.Hash{From: address}).FindByAddress()
	return
}
