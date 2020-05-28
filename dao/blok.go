package dao

import (
	"api/public"
	"github.com/gin-gonic/gin"
	"time"
)

type Blok struct {
	Number     int    `json:"number" orm:"cloumn(number);primary_key" description:"块高"`
	Hash       string `json:"hash" orm:"cloumn(hash)" description:"块hash"`
	Size       int    `json:"size" orm:"cloumn(size)" description:"块大小"`
	Createtime int    `json:"createtime" orm:"cloumn(createtime)" description:"出块时间"`
}

func (b *Blok) TableName() string {
	return "block"
}

func (b *Blok) _Create() error {
	c := gin.Context{}
	c.Set("trace", "_new_blog")
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Create(b).Error
}

func (b *Blok) Create() {
	b.Createtime = int(time.Now().Unix())
	if err := b._Create(); err != nil {
		b.Hash = "hzcblock" + public.RandString(59, public.LOWER_CASE, public.NUMBER)
		b.Create()
	}
	return
}

func (b *Blok) Last() error {
	return public.ChanPool.Table(b.TableName()).Last(b).Error
}
