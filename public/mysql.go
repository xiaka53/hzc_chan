package public

import (
	"github.com/e421083458/gorm"
	"github.com/xiaka53/DeployAndLog/lib"
)

var (
	ChanPool *gorm.DB
)

//数据库初始化
func InitMysql() (err error) {
	var (
		dbpool *gorm.DB
	)
	if dbpool, err = lib.GetGormPool("chan"); err != nil {
		return
	}
	ChanPool = dbpool
	return
}
