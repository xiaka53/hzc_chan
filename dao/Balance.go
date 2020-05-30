package dao

import (
	"api/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type Balance struct {
	Address string  `json:"address" orm:"cloumn(address);primary_key" description:"地址"`
	Token   string  `json:"token" orm:"cloumn(token)" description:"token"`
	Asset   float64 `json:"asset" orm:"cloumn(asset)" description:"余额"`
	Status  int     `json:"status" orm:"status" description:"账户状态1->正常用户|2->异常用户"`
}

func (b *Balance) TableName() string {
	return "balance"
}

func (b *Balance) Create(db *gorm.DB) error {
	return db.Create(b).Error
}

func (b *Balance) HzcCreate() error {
	var c gin.Context
	c.Set("trace", "_new_hzc_account")
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Create(b).Error
}

func (b *Balance) Updates(db *gorm.DB) error {
	return db.Table(b.TableName()).Where("address=? and token=?", b.Address, b.Token).Updates(b).Error
}

func (b *Balance) Close() error {
	return public.ChanPool.Table(b.TableName()).Where(b).Update("status", 2).Error
}

func (b *Balance) Open() error {
	return public.ChanPool.Table(b.TableName()).Where(b).Update("status", 1).Error
}

func (b *Balance) First() error {
	return public.ChanPool.Where(b).First(b).Error
}

func (b *Balance) Find() (data []Balance) {
	public.ChanPool.Where(b).Find(data)
	return data
}

func (b *Balance) Sum() error {
	return public.ChanPool.Where(b).Select("sum(asset) asset").First(b).Error
}

func (b *Balance) RecordNotFound() bool {
	return public.ChanPool.Where(b).RecordNotFound()
}
