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
	return public.ChanPool.Updates(m).Error
}

func (m *Miner) Update() error {
	return public.ChanPool.Table(m.TableName()).Update("status", 2).Error
}

func (m *Miner) Create() error {
	c := gin.Context{}
	c.Set("trace", "_new_blog")
	return public.ChanPool.SetCtx(&c).Create(m).Error
}
