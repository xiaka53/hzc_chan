package _interface

import "api/dao"

//余额
type Balance interface {
	GetBalance(address, token string) (balance float64)
	GetBalanceAll(token string) (balance float64)
}

//账号
type Account interface {
	NewAccount() (address, keys, mnemonic string)
	ImportRawKey(key string) (address string)
	LockAccount(address string) bool
	UnlockAccount(address string) bool
	ListAccount() (data []string)
}

//交易
type Transaction interface {
	SendTransaction(from, to, date string, value, gas float64) (tx string)
	GetTransactionByHash(hx string) (hash dao.Hash)
	GetTransactionByBlock(block int) (data []dao.Hash)
	GetTransactionByAddress(address string) (data []dao.Hash)
}

//挖矿
type Miner interface {
	Start(threads int) bool
	Stop() bool
	SetBase(address string) bool
}

//代币
type Token interface {
	Address() (data []dao.Token)
	Deploy()
}
