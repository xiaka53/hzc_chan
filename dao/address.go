package dao

import (
	"api/public"
	"github.com/gin-gonic/gin"
)

type Address struct {
	Address    string `json:"number" orm:"cloumn(number);primary_key" description:"地址"`
	Status     int    `json:"status" orm:"cloumn(status)" description:"状态1->未使用｜2->已使用"`
	Keys       string `json:"keys" orm:"cloumn(keys)" description:"私钥"`
	Mnemonic   string `json:"mnemonic" orm:"cloumn(mnemonic)" description:"助记词"`
	Createtime int    `json:"createtime" orm:"cloumn(createtime)" description:"出块时间"`
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) Create() error {
	c := gin.Context{}
	c.Set("trace", "_new_address")
	a.Status = 1
	//a.Createtime = int(time.Now().Unix())
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Create(a).Error
}

func (a *Address) First() error {
	return public.ChanPool.Where(a).First(a).Error
}

func (a *Address) Find() (data []Address) {
	public.ChanPool.Where(a).Find(&data)
	return data
}

func (a *Address) Updates() error {
	c := gin.Context{}
	c.Set("trace", "_out_address")
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Table(a.TableName()).Where("address=?", a.Address).Updates(a).Error
}
