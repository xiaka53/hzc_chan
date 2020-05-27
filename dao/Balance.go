package dao

import (
	"api/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"math/big"
)

type Balance struct {
	Address string  `json:"address" orm:"cloumn(address);primary_key" description:"地址"`
	Token   string  `json:"token" orm:"cloumn(token)" description:"token"`
	Asset   big.Int `json:"asset" orm:"cloumn(asset)" description:"余额"`
	Status  int     `json:"status" orm:"status" description:"账户状态1->正常用户|2->异常用户"`
}

func (b *Balance) TableName() string {
	return "token"
}

func (b *Balance) Create(db *gorm.DB) error {
	var c gin.Context
	c.Set("trace", "_new_balance")
	return db.SetCtx(public.GetGinTraceContext(&c)).Create(b).Error
}

func (b *Balance) First() error {
	return public.ChanPool.Where(b).First(b).Error
}

func (b *Balance) Find() (data []Balance) {
	public.ChanPool.Where(b).Find(data)
	return data
}