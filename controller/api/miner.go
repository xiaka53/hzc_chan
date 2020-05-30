package api

import (
	_interface "api/controller/interface"
	"api/controller/interface_v1"
	"api/dto"
	"api/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

type miner struct {
	_interface.Miner
}

func MinerRouter(r *gin.RouterGroup) {
	var m miner
	m.Miner = interface_v1.GetMiner()
	r.POST("start", m.start)
	r.GET("stop", m.stop)
	r.POST("set", m.set)
}

// @Summary 开启挖矿
// @Tags miner
// @Id 010
// @Produce  json
// @Param threads query string true "开启挖矿线程数"
// @Success 200 {bool} bool
// @Router /chan/miner/start [post]
func (m *miner) start(c *gin.Context) {
	var (
		param dto.StartInout
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ERROR, err)
		return
	}
	if !m.Miner.Start(param.Threads) {
		middleware.ResponseError(c, middleware.ERROR, errors.New("Error miner start!"))
		return
	}
	middleware.ResponseSuccess(c, true)
}

// @Summary 关闭挖矿
// @Tags miner
// @Id 011
// @Produce  json
// @Success 200 {bool} bool
// @Router /chan/miner/stop [get]
func (m *miner) stop(c *gin.Context) {
	if !m.Miner.Stop() {
		middleware.ResponseError(c, middleware.ERROR, errors.New("Error miner stop!"))
		return
	}
	middleware.ResponseSuccess(c, true)
}

// @Summary 设置挖矿地址
// @Tags miner
// @Id 012
// @Produce  json
// @Param address query string true "地址"
// @Success 200 {bool} bool
// @Router /chan/miner/set [post]
func (m *miner) set(c *gin.Context) {
	var (
		param dto.SetInout
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ERROR, err)
		return
	}
	if !m.Miner.SetBase(param.Address) {
		middleware.ResponseError(c, middleware.ERROR, errors.New("Error set miner address!"))
		return
	}
	middleware.ResponseSuccess(c, true)
}
