package dto

import (
	"api/public"
	"errors"
	"github.com/gin-gonic/gin"
)

type SendTransactionInout struct {
	From  string  `form:"from" json:"from" validate:"len=43,sql"`
	To    string  `form:"to" json:"to" validate:"len=43,sql"`
	Date  string  `form:"date" json:"date" validate:"omitempty,len=43,sql"`
	Input string  `form:"input" json:"input" validate:"sql,omitempty"`
	Value float64 `form:"value" json:"value" validate:"numeric"`
	Gas   float64 `form:"gas" json:"gas" validate:"numeric"`
}

func (o *SendTransactionInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type GetTransactionByHashInout struct {
	Hash string `form:"hash" json:"hash" validate:"len=67,sql"`
}

func (o *GetTransactionByHashInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type GetTransactionByBlockInout struct {
	Block int `form:"block" json:"block" validate:"numeric"`
}

func (o *GetTransactionByBlockInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type GetTransactionByAddressInout struct {
	Address string `form:"address" json:"address" validate:"len=43,sql"`
}

func (o *GetTransactionByAddressInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}
