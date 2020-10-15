package main

import (
	"gomicrodemo/share/config"
	"gomicrodemo/share/pb"
	"gomicrodemo/share/utils/log"
	"gomicrodemo/user-srv/db"
	"gomicrodemo/user-srv/handler"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

func main() {
	logger := log.Init("user")

	// 使用etcd注册

	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name(config.Namespace+"user"),
		micro.Version("latest"),
	)

	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Info("micro.Action test ...")
			// 先注册db
			db.Init(config.MysqlDSN)
			pb.RegisterUserServiceHandler(service.Server(), handler.NewUserHandler(), server.InternalHandler(true))
			return nil
		}),
		micro.AfterStop(func() error {
			logger.Info("micro.AfterStop test ...")
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Info("micro.AfterStart test ...")
			return nil
		}),
	)

	logger.Info("启动user-srv服务 ...")

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("user-srv服务启动失败 ...")
	}
}
