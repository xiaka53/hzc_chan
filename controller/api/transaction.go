package api

import (
	_interface "api/controller/interface"
	"api/controller/interface_v1"
	"api/dao"
	"api/dto"
	"api/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

type transaction struct {
	_interface.Transaction
}

func TransactionRouter(r *gin.RouterGroup) {
	var t transaction
	t.Transaction = interface_v1.GetTransaction()
	r.GET("send", t.send_transaction)
	r.GET("get_ransactionByHash", t.get_ransactionByHash)
	r.GET("get_transactionByBlock", t.get_transactionByBlock)
	r.GET("get_transactionByAddress", t.get_transactionByAddress)
}

// @Summary 发送交易
// @Tags transaction
// @Id 013
// @Produce  json
// @Param from query string true "转出方"
// @Param to query string true "接收方"
// @Param date query string true "token"
// @Param input query string true "转账备注"
// @Param value query number true "转出金额"
// @Param gas query number true "手续费"
// @Success 200 {bool} bool
// @Router /chan/transaction/send [get]
func (t *transaction) send_transaction(c *gin.Context) {
	var (
		param dto.SendTransactionInout
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ERROR, err)
		return
	}
	if param.Gas < 0.00001 {
		middleware.ResponseError(c, middleware.ERROR, errors.New("Gas is low!"))
		return
	}
	if (&dao.Balance{Address: param.From, Status: 1}).RecordNotFound() {
		middleware.ResponseError(c, middleware.ERROR, errors.New("from:no address"))
		return
	}
	if (&dao.Balance{Address: param.To, Status: 1}).RecordNotFound() {
		middleware.ResponseError(c, middleware.ERROR, errors.New("to:no address"))
		return
	}
	t.Transaction.SendTransaction(param.From, param.To, param.Date, param.Input, param.Value, param.Gas)
	middleware.ResponseSuccess(c, true)
}

// @Summary genjuhash获取交易信息
// @Tags transaction
// @Id 014
// @Produce  json
// @Param hash query string true "hash"
// @Success 200 {obiect} dao.Hash
// @Router /chan/transaction/get_ransactionByHash [get]
func (t *transaction) get_ransactionByHash(c *gin.Context) {
	var (
		param dto.GetTransactionByHashInout
		hash  dao.Hash
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ERROR, err)
		return
	}
	hash = t.Transaction.GetTransactionByHash(param.Hash)
	middleware.ResponseSuccess(c, hash)
}

// @Summary 根据块高获取交易信息
// @Tags transaction
// @Id 015
// @Produce  json
// @Param block query number true "块高"
// @Success 200 {obiect} dao.Hash
// @Router /chan/transaction/get_transactionByBlock [get]
func (t *transaction) get_transactionByBlock(c *gin.Context) {
	var (
		param dto.GetTransactionByBlockInout
		hash  []dao.Hash
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ERROR, err)
		return
	}
	hash = t.Transaction.GetTransactionByBlock(param.Block)
	middleware.ResponseSuccess(c, hash)
}

// @Summary 根据地址获取交易信息
// @Tags transaction
// @Id 016
// @Produce  json
// @Param address query string true "地址"
// @Success 200 {obiect} dao.Hash
// @Router /chan/transaction/get_transactionByAddress [get]
func (t *transaction) get_transactionByAddress(c *gin.Context) {
	var (
		param dto.GetTransactionByAddressInout
		hash  []dao.Hash
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ERROR, err)
		return
	}
	hash = t.Transaction.GetTransactionByAddress(param.Address)
	middleware.ResponseSuccess(c, hash)
}
