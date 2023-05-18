package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/CocaineCong/grpc-todolist/app/user/internal/handler"
	"github.com/CocaineCong/grpc-todolist/app/user/internal/repository/db/dao"
	"github.com/CocaineCong/grpc-todolist/config"
	"github.com/CocaineCong/grpc-todolist/idl/user"
	"github.com/CocaineCong/grpc-todolist/pkg/discovery"
)

func main() {
	config.InitConfig()
	dao.InitDB()
	// etcd 地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()
	// 绑定service
	user.RegisterUserServiceServer(server, handler.NewUserService())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
