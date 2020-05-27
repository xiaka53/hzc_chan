package task

import (
	"api/dao"
	"api/public"
	"strings"
	"sync"
)

const EARLYWARNING int = 30

type account struct {
	Chan  chan dao.Address
	Sync  sync.RWMutex
	Count int
}

func AccountInit() {
	examples := &account{
		Chan:  make(chan dao.Address),
		Sync:  sync.RWMutex{},
		Count: 0,
	}
	data := (&dao.Address{Status: 1}).Find()
	for _, v := range data {
		examples.add(v)
	}
	examples.inspection()
}

func (a *account) inspection() {
	if a.Count < EARLYWARNING {
		a.newAccount()
	}
	return
}

func (a *account) newAccount() {
	account := "hzc" + public.RandString(40, public.LOWER_CASE, public.NUMBER)
	keys := "hzc" + public.RandString(32, public.LOWER_CASE, public.NUMBER)
	mnemonic := (&dao.Words{}).Rand(12)
	address := dao.Address{Address: account, Keys: keys, Mnemonic: strings.Join(mnemonic, ",")}
	if err := (&address).Create(); err == nil {
		a.add(address)
	}
	a.inspection()
	return
}

func (a *account) add(address dao.Address) {
	go func() {
		for {
			a.Sync.RLock()
			a.Count += 1
			a.Sync.RUnlock()
			a.Chan <- address
		}
	}()
	return
}

func (a *account) out() dao.Address {
	address := <-a.Chan
	a.Sync.RLock()
	a.Count -= 1
	a.Sync.RUnlock()
	go a.inspection()
	return address
}
