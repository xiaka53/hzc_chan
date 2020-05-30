package dao

import (
	"api/public"
	"github.com/gin-gonic/gin"
)

type Miner struct {
	Id      string `json:"id" orm:"cloumn(id);primary_key" description:"ID"`
	Address string `json:"address" orm:"cloumn(address)" description:"挖矿地址"`
	Threads int    `json:"threads" orm:"cloumn(threads)" description:"线程数"`
	Status  int    `json:"status" orm:"cloumn(status)" description:"挖矿状态1->挖矿中｜2->停止挖矿"`
}

func (m *Miner) TableName() string {
	return "miner"
}

func (m *Miner) First() error {
	return public.ChanPool.Where(m).Last(m).Error
}

func (m *Miner) Last() error {
	return public.ChanPool.Where(m).Last(m).Error
}

func (m *Miner) Updates() error {
	c := gin.Context{}
	c.Set("trace", "_updates_miner")
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Table(m.TableName()).Updates(m).Error
}

func (m *Miner) Update() error {
	c := gin.Context{}
	c.Set("trace", "_update_miner_all")
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Table(m.TableName()).Where("status=1").Update("status", 2).Error
}

func (m *Miner) Create() error {
	c := gin.Context{}
	c.Set("trace", "_new_miner")
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Create(m).Error
}
