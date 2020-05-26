package main

import (
	"api/public"
	"api/router"
	"github.com/xiaka53/DeployAndLog/lib"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title HZC赫兹项目
// @version 1.0
// @description 暂无备注
// @contact.name 陶然
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	if err := lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"}); err != nil {
		log.Fatal(err)
	}
	if err := public.InitMysql(); err != nil {
		log.Fatal(err)
	}
	defer lib.Destroy()
	public.InitValidate()
	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
