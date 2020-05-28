package api

import (
	_interface "api/controller/interface"
	"api/controller/interface_v1"
	"api/dto"
	"api/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

type account struct {
	_interface.Account
}

func AccountRouter(r *gin.RouterGroup) {
	var a account
	a.Account = interface_v1.GetAccount()
	r.GET("new", a.getAddress)
	r.GET("import_rawKey", a.importRawKey)
	r.GET("export_rawKey", a.exportRawKey)
	r.GET("export_mnemonic", a.exportMnemonic)
	r.GET("lock", a.lockAccount)
	r.GET("unlock", a.unlockAccount)
	r.GET("list", a.listAccount)
}

// @Summary 生成地址
// @Tags account
// @Id 001
// @Produce  json
// @Success 200 {string} string
// @Router /chan/account/getAddress [get]
func (a *account) getAddress(c *gin.Context) {
	address, _, _ := a.Account.NewAccount()
	middleware.ResponseSuccess(c, address)
}

// @Summary 导入地址
// @Tags account
// @Id 002
// @Produce  json
// @Param keys query string true "密钥｜助记词"
// @Success 200 {string} string
// @Router /chan/account/importRawKey [post]
func (a *account) importRawKey(c *gin.Context) {
	var (
		param   dto.ImportRawKeyInout
		address string
		err     error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.PARAMETER, err)
		return
	}
	address = a.Account.ImportRawKey(param.Keys)
	middleware.ResponseSuccess(c, address)
}

// @Summary 导出私钥
// @Tags account
// @Id 003
// @Produce  json
// @Param address query string true "地址"
// @Success 200 {string} string
// @Router /chan/account/exportRawKey [post]
func (a *account) exportRawKey(c *gin.Context) {
	var (
		param dto.ExportRawKeyInout
		keys  string
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.PARAMETER, err)
		return
	}
	if param.Address[:3] != "hzc" {
		middleware.ResponseError(c, middleware.ERROR, errors.New("error address!"))
		return
	}
	keys = a.Account.ExportRawKey(param.Address)
	if len(keys) != 35 {
		middleware.ResponseError(c, middleware.ERROR, errors.New("no address!"))
		return
	}
	middleware.ResponseSuccess(c, keys)
}

// @Summary 导出助记词
// @Tags account
// @Id 004
// @Produce  json
// @Param address query string true "地址"
// @Success 200 {string} string
// @Router /chan/account/exportMnemonic [post]
func (a *account) exportMnemonic(c *gin.Context) {
	var (
		param    dto.ExportMnemonicInout
		mnemonic string
		err      error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.PARAMETER, err)
		return
	}
	if param.Address[:3] != "hzc" {
		middleware.ResponseError(c, middleware.ERROR, errors.New("error address!"))
		return
	}
	mnemonic = a.Account.ExportMnemonic(param.Address)
	if len(mnemonic) < 3 {
		middleware.ResponseError(c, middleware.ERROR, errors.New("no address!"))
		return
	}
	middleware.ResponseSuccess(c, mnemonic)
}

// @Summary 锁定账户（锁定后无法转账，但是可以收帐）
// @Tags account
// @Id 005
// @Produce  json
// @Param address query string true "地址"
// @Success 200 {string} string
// @Router /chan/account/lockAccount [post]
func (a *account) lockAccount(c *gin.Context) {
	var (
		param dto.LockAccountInout
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.PARAMETER, err)
		return
	}
	if param.Address[:3] != "hzc" {
		middleware.ResponseError(c, middleware.ERROR, errors.New("error address!"))
		return
	}
	if !a.Account.LockAccount(param.Address) {
		middleware.ResponseError(c, middleware.ERROR, errors.New("no address!"))
		return
	}
	middleware.ResponseSuccess(c, "")
}

// @Summary 解锁账户
// @Tags account
// @Id 006
// @Produce  json
// @Param address query string true "地址"
// @Success 200 {string} string
// @Router /chan/account/unlockAccount [post]
func (a *account) unlockAccount(c *gin.Context) {
	var (
		param dto.UnlockAccounttInout
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.PARAMETER, err)
		return
	}
	if param.Address[:3] != "hzc" {
		middleware.ResponseError(c, middleware.ERROR, errors.New("error address!"))
		return
	}
	if !a.Account.UnlockAccount(param.Address) {
		middleware.ResponseError(c, middleware.ERROR, errors.New("no address!"))
		return
	}
	middleware.ResponseSuccess(c, "")
}

// @Summary 获取所有账户
// @Tags account
// @Id 007
// @Produce  json
// @Success 200 {string} string
// @Router /chan/account/listAccount [get]
func (a *account) listAccount(c *gin.Context) {
	data := a.Account.ListAccount()
	middleware.ResponseSuccess(c, data)
}
