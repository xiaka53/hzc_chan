package dto

import (
	"api/public"
	"errors"
	"github.com/gin-gonic/gin"
)

type BalanceInout struct {
	Address string `form:"address" json:"address" validate:"len=43,sql"`
	Token   string `form:"token" json:"token" validate:"sql,omitempty"`
}

func (o *BalanceInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type AllInout struct {
	Token string `form:"token" json:"token" validate:"sql,omitempty"`
}

func (o *AllInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}
