package main

import (
	"fmt"
	"log"

	"sky-take-out-go/config"
	"sky-take-out-go/dao"
	"sky-take-out-go/router"
	"sky-take-out-go/utils"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := config.Load("config.yaml"); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	db, err := gorm.Open(mysql.Open(config.Cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}
	dao.DB = db
	dao.InitRedis(config.Cfg.Redis)
	utils.InitRmq()
	defer utils.CloseConnect()

	r := router.Setup()

	addr := fmt.Sprintf(":%d", config.Cfg.Server.Port)
	logger.Info("server starting", zap.String("addr", addr))
	if err := r.Run(addr); err != nil {
		logger.Fatal("server failed", zap.Error(err))
	}
}
