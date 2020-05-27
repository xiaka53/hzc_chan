package private

import "api/dao"

type balance interface {
	getBalance(address, token string) (balance float64)
	getBalanceAll(token string) (balance float64)
}

type address interface {
	newAccount() (address, keys, mnemonic string)
	importRawKey(key string) (address string)
	lockAccount(address string) bool
	unlockAccount(address string) bool
	listAccount() (data []string)
}

type transaction interface {
	sendTransaction(from, to, date string, value, gas float64)
	getTransactionByHash(hx string) (hash dao.Hash)
	getTransactionByBlock(block int) (data []dao.Hash)
	getTransactionByAddress(address string) (data []dao.Hash)
}

type miner interface {
	start(threads int) bool
	stop() bool
	setBase(address string) bool
}

type token interface {
	address() (data []dao.Token)
	deploy()
}
