package dao

import (
	"api/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type Hash struct {
	Hash            string  `json:"hash" orm:"cloumn(hash);primary_key" description:"哈希"`
	Nonce           int     `json:"nonce" orm:"cloumn(nonce)" description:"本次交易之前发送方已经生成的交易数量"`
	BlockHash       string  `json:"block_hash" orm:"cloumn(block_hash)" description:"交易所在块的哈希，对于挂起块，该值为null"`
	BlockNumber     int     `json:"block_number" orm:"cloumn(block_number)" description:"交易所在块的编号，对于挂起块，该值为null"`
	From            string  `json:"from" orm:"cloumn(from)" description:"交易发送方地址"`
	To              string  `json:"to" orm:"cloumn(to)" description:"交易接收方地址，对于合约创建交易，该值为null"`
	Value           float64 `json:"value" orm:"cloumn(value)" description:"发送的HZC数量 单位：wei"`
	Gas             float64 `json:"gas" orm:"cloumn(gas)" description:"燃料费  单位：wei"`
	Input           string  `json:"input" orm:"cloumn(input)" description:"随交易发送的数据"`
	ContractAddress string  `json:"contract_address" orm:"cloumn(contract_address)" description:"如果是创建合约，这个是合约地址"`
	Status          string  `json:"status" orm:"cloumn(status)" description:"交易状态0->打包中|1->交易成功|2->交易回滚"`
	Index           float64 `json:"index" orm:"cloumn(index)" description:"代币交易时，代币金额  单位：wei"`
	Createtime      int     `json:"createtime" orm:"cloumn(createtime)" description:"交易时间"`
}

func (h *Hash) TableName() string {
	return "hash"
}

func (h *Hash) _Create() error {
	var c gin.Context
	c.Set("trace", "_new_hash")
	h.Createtime = int(time.Now().Unix())
	h.Hash = "hzchash" + public.RandString(60, public.LOWER_CASE, public.NUMBER)
	return public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Create(h).Error
}

func (h *Hash) Create() {
	public.ChanPool.Where("from=? or to=?", h.From, h.From).Count(h.Nonce)
	if err := h._Create(); err != nil {
		h.Create()
	}
	return
}

func (h *Hash) Update(db *gorm.DB) error {
	return db.Table(h.TableName()).Where("hash=?", h.Hash).Update("status", h.Status).Error
}

func (h *Hash) Updates(hash []string) {
	if len(hash) < 1 {
		return
	}
	var c gin.Context
	c.Set("trace", "_update_hash")
	public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Table(h.TableName()).Where("hash in (?)", hash).Updates(map[string]interface{}{"block_hash": h.BlockHash, "block_number": h.BlockNumber})
	public.ChanPool.SetCtx(public.GetGinTraceContext(&c)).Table(h.TableName()).Where("hash in (?) and status='0'", hash).Update("status", "1")
}

func (h *Hash) First() error {
	return public.ChanPool.Where(h).First(h).Error
}

func (h *Hash) Find() []Hash {
	var data []Hash
	public.ChanPool.Where(h).Order("createtime desc").Find(&data)
	return data
}

func (h *Hash) FindByAddress() []Hash {
	var data []Hash
	public.ChanPool.Where("`from`=? or `to`=?", h.From, h.From).Order("createtime desc").Find(&data)
	return data
}

func (h *Hash) FindByInput() []Hash {
	var data []Hash
	public.ChanPool.Where("`from`=? or `to`=?", h.From, h.From).Where("input=?", h.Input).Order("createtime desc").Find(&data)
	return data
}
