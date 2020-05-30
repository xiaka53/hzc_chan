package dto

import (
	"api/public"
	"errors"
	"github.com/gin-gonic/gin"
)

type StartInout struct {
	Threads int `form:"threads" json:"threads" validate:"min=1"`
}

func (o *StartInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}

type SetInout struct {
	Address string `form:"address" json:"address" validate:"len=43,sql"`
}

func (o *SetInout) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	if err = public.Validate.Struct(o); err != nil {
		return errors.New("param error!")
	}
	return
}
