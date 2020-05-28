package dto

import (
	"api/public"
	"errors"
	"github.com/gin-gonic/gin"
)

type ImportRawKeyInout struct {
	Keys string `form:"keys" json:"keys" validate:"sql"`
}

func (o *ImportRawKeyInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type ExportRawKeyInout struct {
	Address string `form:"address" json:"address" validate:"len=43,sql"`
}

func (o *ExportRawKeyInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type ExportMnemonicInout struct {
	Address string `form:"address" json:"address" validate:"len=43,sql"`
}

func (o *ExportMnemonicInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type LockAccountInout struct {
	Address string `form:"address" json:"address" validate:"len=43,sql"`
}

func (o *LockAccountInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type UnlockAccounttInout struct {
	Address string `form:"address" json:"address" validate:"len=43,sql"`
}

func (o *UnlockAccounttInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}
