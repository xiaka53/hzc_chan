package task

import (
	"api/dao"
	"api/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
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

func (h *hash) add(hs dao.Hash) {
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
		if hashData.Status == 2 {
			goto CON
		}
		fromBalance.Address = hashData.From
		fromBalance.Token = "HZC"
		if err := (&fromBalance).First(); err != nil {
			hashData.Status = 2
			goto CON
		}
		fromBalance.Asset.Sub(fromBalance.Asset, hashData.Gas)
		if fromBalance.Asset.Int64() < 0 {
			goto CON
		}
		fromTokenBalance.Address = hashData.From
		c.Set("trace", "_new_transfer")
		db = public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Begin()
		if hashData.Index.Int64() > 0 {
			if hashData.From == hashData.To { //创建代币
				fromTokenBalance.Token = hashData.ContractAddress
				fromTokenBalance.Asset = hashData.Index
				if err = (fromTokenBalance).Create(db); err != nil {
					goto ERR
				}
			} else { //代币交易
				fromTokenBalance.Token = hashData.To[:43]
				if err = (&fromTokenBalance).First(); err != nil {
					goto ERR
				}
				fromTokenBalance.Asset.Sub(fromTokenBalance.Asset, hashData.Index)
				if err = (&fromTokenBalance).Updates(db); err != nil {
					goto ERR
				}
				toBalance.Address = hashData.To[43:]
				toBalance.Token = fromTokenBalance.Token
				if err = (&toBalance).First(); err != nil {
					goto ERR
				}
				toBalance.Asset.Add(toBalance.Asset, hashData.Index)
				if err = (&toBalance).Updates(db); err != nil {
					goto ERR
				}
			}
		} else {
			fromBalance.Asset.Sub(fromBalance.Asset, hashData.Value)
			toBalance.Address = hashData.To
			toBalance.Token = fromBalance.Token
			if err = (&toBalance).First(); err != nil {
				goto ERR
			}
			toBalance.Asset.Add(toBalance.Asset, hashData.Index)
			if err = (&toBalance).Updates(db); err != nil {
				goto ERR
			}
		}
		if err = (&fromBalance).Updates(db); err != nil {
			goto ERR
		}
		db.Commit()
		goto CON
	ERR:
		hashData.Status = 2
		db.Rollback()
	CON:
		(&hashData).Create()
		getBlock().add(hashData.Hash)
		continue
	}
}
