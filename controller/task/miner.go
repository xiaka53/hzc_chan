package task

import (
	"api/dao"
	"api/public"
	"github.com/gin-gonic/gin"
	"math/big"
	"sync"
)

type miner struct {
	address string
	threads int
	mutex   sync.RWMutex
}

var examplesMiner *miner

func MinerInit() {
	data := dao.Miner{Status: 1}
	if err := (&data).Last(); err != nil {
		return
	}
	if data.Threads < 1 {
		return
	}
	examplesMiner = &miner{
		address: data.Address,
		threads: data.Threads,
	}
	return
}

func GetMiner() *miner {
	return examplesMiner
}

func (m *miner) server() (num int) {
	num = m.threads * 10 //TODO 单个线程挖矿数量为10
	if num < 10 {
		return
	}
	balance := &dao.Balance{Address: m.address, Token: "HZC"}
	_ = balance.First()
	balance.Asset.Add(balance.Asset, big.NewInt(int64(num)))
	var c gin.Context
	c.Set("trace", "_new_hash")
	if err := balance.Updates(public.ChanPool.SetCtx(public.GetGinTraceContext(&c))); err != nil {
		return 0
	}
	return
}

func (m *miner) Start(threads int) bool {
	data := dao.Miner{Status: 1}
	if err := (&data).First(); err != nil {
		data.Status = 2
		if err := (&data).Last(); err != nil {
			return false
		}
	}
	data.Status = 1
	data.Threads = threads
	if err := (&data).Updates(); err != nil {
		return false
	}
	if data.Threads < 1 {
		return false
	}
	m.mutex.Lock()
	m.address = data.Address
	m.threads = data.Threads
	m.mutex.Unlock()
	return true
}

func (m *miner) Stop() bool {
	data := dao.Miner{Status: 1}
	if err := (&data).Last(); err != nil {
		return false
	}
	data.Status = 2
	if err := (&data).Updates(); err != nil {
		return false
	}
	m.mutex.Lock()
	m.address = data.Address
	m.threads = data.Threads
	m.mutex.Unlock()
	return true
}

func (m *miner) Set(address string) bool {
	data := dao.Miner{Address: address}
	if err := data.Update(); err != nil {
		return false
	}
	_ = (&data).First()
	data.Status = 1
	if err := (&data).Updates(); err != nil {
		return false
	}
	return true
}
