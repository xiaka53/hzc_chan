package dao

import (
	"api/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"math/big"
)

type Hash struct {
	Hash            string  `json:"hash" orm:"cloumn(hash);primary_key" description:"哈希"`
	Nonce           int     `json:"nonce" orm:"cloumn(nonce)" description:"本次交易之前发送方已经生成的交易数量"`
	BlockHash       string  `json:"blockHash" orm:"cloumn(blockHash)" description:"交易所在块的哈希，对于挂起块，该值为null"`
	BlockNumber     int     `json:"blockNumber" orm:"cloumn(blockNumber)" description:"交易所在块的编号，对于挂起块，该值为null"`
	From            string  `json:"from" orm:"cloumn(from)" description:"交易发送方地址"`
	To              string  `json:"to" orm:"cloumn(to)" description:"交易接收方地址，对于合约创建交易，该值为null"`
	Value           big.Int `json:"value" orm:"cloumn(value)" description:"发送的HZC数量 单位：wei"`
	Gas             big.Int `json:"gas" orm:"cloumn(gas)" description:"燃料费  单位：wei"`
	Input           string  `json:"input" orm:"cloumn(input)" description:"随交易发送的数据"`
	ContractAddress string  `json:"contractAddress" orm:"cloumn(contractAddress)" description:"如果是创建合约，这个是合约地址"`
	Status          int     `json:"status" orm:"cloumn(status)" description:"交易状态0->打包中|1->交易成功|2->交易回滚"`
	Index           big.Int `json:"index" orm:"cloumn(index)" description:"代币交易时，代币金额  单位：wei"`
}

func (h *Hash) TableName() string {
	return "hash"
}

func (h *Hash) Create(db *gorm.DB) error {
	var c gin.Context
	c.Set("trace", "_new_hash")
	return db.SetCtx(public.GetGinTraceContext(&c)).Create(h).Error
}

func (h *Hash) Updates(hash []string) {
	var c gin.Context
	c.Set("trace", "_update_hash")
	public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Table(h.TableName()).Where("hash in (?)", hash).Updates(map[string]interface{}{"blockHash": h.BlockHash, "blockNumber": h.BlockNumber})
	public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Table(h.TableName()).Where("hash in (?) and status=0", hash).Update("status", 1)
}
