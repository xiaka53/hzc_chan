package _interface

import "api/dao"

//余额
type balance interface {
	getBalance(address, token string) (balance float64)
	getBalanceAll(token string) (balance float64)
}

//账号
type account interface {
	newAccount() (address, keys, mnemonic string)
	importRawKey(key string) (address string)
	lockAccount(address string) bool
	unlockAccount(address string) bool
	listAccount() (data []string)
}

//交易
type transaction interface {
	sendTransaction(from, to, date string, value, gas float64)
	getTransactionByHash(hx string) (hash dao.Hash)
	getTransactionByBlock(block int) (data []dao.Hash)
	getTransactionByAddress(address string) (data []dao.Hash)
}

//挖矿
type miner interface {
	start(threads int) bool
	stop() bool
	setBase(address string) bool
}

//代币
type token interface {
	address() (data []dao.Token)
	deploy()
}
