package dao

import (
	"api/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"math/big"
)

type Token struct {
	Token            string   `json:"token" orm:"cloumn(token);primary_key" description:"token"`
	Name             string   `json:"name" orm:"cloumn(name)" description:"token名称"`
	Symbol           string   `json:"symbol" orm:"cloumn(symbol)" description:"token符号"`
	Logo             string   `json:"logo" orm:"cloumn(logo)" description:"代币logo"`
	IssueNumber      *big.Int `json:"issue_number" orm:"cloumn(issue_number)" description:"发行数量"`
	AdditionalIssue  int      `json:"additional_issue" orm:"cloumn(additional_issue)" description:"是否增发1->不增发|2->增发"`
	AddIssue         *big.Int `json:"add_issue" orm:"cloumn(add_issue)" description:"增发数量"`
	AuthorityAddress string   `json:"authority_address" orm:"cloumn(authority_address)" description:"权限地址"`
	IsClose          int      `json:"is_close" orm:"cloumn(is_close)" description:"关闭交易权限1->没权限|2->有权限"`
	Status           int      `json:"status" orm:"cloumn(status)" description:"是否可以交易1->可以交易|2->不能交易"`
}

func (t *Token) TableName() string {
	return "token"
}

func (t *Token) Create(db *gorm.DB) error {
	var c gin.Context
	c.Set("trace", "_new_token")
	return db.SetCtx(public.GetGinTraceContext(&c)).Create(t).Error
}

func (t *Token) First() error {
	return public.ChanPool.Where(t).First(t).Error
}
