package api

import (
	_interface "api/controller/interface"
	"api/controller/interface_v1"
	"api/dto"
	"api/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

type balance struct {
	_interface.Balance
}

func balanceRouter(r *gin.RouterGroup) {
	var b balance
	b.Balance = interface_v1.GetBalance()

}

// @Summary 获取余额
// @Tags blance
// @Id 008
// @Produce  json
// @Param address query string true "地址"
// @Param token query string false "token币种"
// @Success 200 {string} string
// @Router /chan/blance/balance [get]
func (b *balance) balance(c *gin.Context) {
	var (
		param  dto.BalanceInout
		amount float64
		err    error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.PARAMETER, err)
		return
	}
	if param.Address[:3] != "hzc" {
		middleware.ResponseError(c, middleware.ERROR, errors.New("error address!"))
		return
	}
	if len(param.Token) < 1 {
		param.Token = "HZC"
	}
	amount = b.Balance.GetBalance(param.Address, param.Token)
	middleware.ResponseSuccess(c, amount)
}

// @Summary 获取余额总数
// @Tags blance
// @Id 009
// @Produce  json
// @Param token query string false "token币种"
// @Success 200 {string} string
// @Router /chan/blance/balance [get]
func (b *balance) all(c *gin.Context) {
	var (
		param  dto.AllInout
		amount float64
		err    error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.PARAMETER, err)
		return
	}
	if len(param.Token) < 1 {
		param.Token = "HZC"
	}
	amount = b.Balance.GetBalanceAll(param.Token)
	middleware.ResponseSuccess(c, amount)
}
