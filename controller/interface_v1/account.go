package interface_v1

import (
	"api/controller/task"
	"api/dao"
)

type apiAccount struct {
}

func GetAccount() *apiAccount {
	return new(apiAccount)
}

func (a *apiAccount) NewAccount() (address, keys, mnemonic string) {
	data := task.GetAccount().Out()
	address = data.Address
	keys = data.Keys
	mnemonic = data.Mnemonic
	return
}

func (a *apiAccount) ImportRawKey(key string) (address string) {
	var account dao.Address
	if len(key) == 35 {
		account.Keys = key
	} else {
		account.Mnemonic = key
	}
	_ = (&account).First()
	address = account.Address
	return
}

func (a *apiAccount) LockAccount(address string) bool {
	var account dao.Address
	account.Address = address
	if err := (&account).First(); err != nil {
		return false
	}
	var balance dao.Balance
	balance.Address = account.Address
	if err := (&balance).Close(); err != nil {
		return false
	}
	return true
}

func (a *apiAccount) UnlockAccount(address string) bool {
	var account dao.Address
	account.Address = address
	if err := (&account).First(); err != nil {
		return false
	}
	var balance dao.Balance
	balance.Address = account.Address
	if err := (&balance).Open(); err != nil {
		return false
	}
	return true
}

func (a *apiAccount) ListAccount() (data []string) {
	address := (&dao.Address{Status: 2}).Find()
	for _, v := range address {
		data = append(data, v.Address)
	}
	return
}
