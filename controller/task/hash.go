package task

import (
	"api/dao"
	"api/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type hash struct {
	Chan chan dao.Hash
}

var examplesHash *hash

func HashInit() {
	examplesHash = &hash{Chan: make(chan dao.Hash)}
	examplesHash.server()
}

func GetHash() *hash {
	return examplesHash
}

func (h *hash) Add(hs dao.Hash) {
	go func() {
		h.Chan <- hs
	}()
}

func (h *hash) server() {
	for {
		hashData := <-h.Chan
		var (
			fromBalance      dao.Balance
			fromTokenBalance dao.Balance
			toBalance        dao.Balance
			db               *gorm.DB
			c                gin.Context
			err              error
		)
		db = public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Begin()
		if hashData.Status == "2" {
			goto CON
		}
		fromBalance.Address = hashData.From
		fromBalance.Token = "HZC"
		if err := (&fromBalance).First(); err != nil {
			hashData.Status = "2"
			goto CON
		}
		if fromBalance.Status == 2 {
			hashData.Status = "2"
			goto CON
		}
		fromBalance.Asset, _ = decimal.NewFromFloat(fromBalance.Asset).Sub(decimal.NewFromFloat(hashData.Gas)).Float64()
		if fromBalance.Asset < 0 {
			goto CON
		}
		fromTokenBalance.Address = hashData.From
		c.Set("trace", "_new_transfer")
		if hashData.Index > 0 {
			if hashData.From == hashData.To { //创建代币
				fromTokenBalance.Token = hashData.ContractAddress
				fromTokenBalance.Asset = hashData.Index
				if err = (&fromTokenBalance).Create(db); err != nil {
					goto ERR
				}
			} else { //代币交易
				fromTokenBalance.Token = hashData.ContractAddress
				if err = (&fromTokenBalance).First(); err != nil {
					goto ERR
				}
				if fromTokenBalance.Status == 2 {
					goto ERR
				}
				fromTokenBalance.Asset, _ = decimal.NewFromFloat(fromTokenBalance.Asset).Sub(decimal.NewFromFloat(hashData.Index)).Float64()
				if err = (&fromTokenBalance).Updates(db); err != nil {
					goto ERR
				}
				toBalance.Address = hashData.To
				toBalance.Token = fromTokenBalance.Token
				_ = (&toBalance).First()
				toBalance.Asset, _ = decimal.NewFromFloat(toBalance.Asset).Add(decimal.NewFromFloat(hashData.Index)).Float64()
				if err = (&toBalance).Updates(db); err != nil {
					goto ERR
				}
			}
		} else {
			fromBalance.Asset, _ = decimal.NewFromFloat(fromBalance.Asset).Sub(decimal.NewFromFloat(hashData.Value)).Float64()
			toBalance.Address = hashData.To
			toBalance.Token = fromBalance.Token
			_ = (&toBalance).First()
			toBalance.Asset, _ = decimal.NewFromFloat(toBalance.Asset).Add(decimal.NewFromFloat(hashData.Value)).Float64()
			if err = (&toBalance).Updates(db); err != nil {
				goto ERR
			}
		}
		if err = (&fromBalance).Updates(db); err != nil {
			goto ERR
		}
		hashData.Status = "0"
		goto CON
	ERR:
		hashData.Status = "2"
		db.Rollback()
	CON:
		if err = (&hashData).Update(db); err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
		getBlock().add(hashData.Hash)
		continue
	}
}
