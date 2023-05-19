package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/CocaineCong/grpc-todolist/app/gateway/routes"
	"github.com/CocaineCong/grpc-todolist/config"
	taskPb "github.com/CocaineCong/grpc-todolist/idl/task/pb"
	userPb "github.com/CocaineCong/grpc-todolist/idl/user/pb"
	"github.com/CocaineCong/grpc-todolist/pkg/discovery"
	"github.com/CocaineCong/grpc-todolist/pkg/util/shutdown"
)

func main() {
	config.InitConfig()
	go startListen() // 转载路由
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :3000")
}

func startListen() {
	// etcd注册
	etcdAddress := []string{config.Conf.Etcd.Address}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// 服务名
	userServiceName := config.Conf.Domain["user"].Name
	taskServiceName := config.Conf.Domain["task"].Name

	// RPC 连接
	connUser, err := RPCConnect(ctx, userServiceName, etcdRegister)
	if err != nil {
		return
	}
	userService := userPb.NewUserServiceClient(connUser)

	connTask, err := RPCConnect(ctx, taskServiceName, etcdRegister)
	if err != nil {
		return
	}
	taskService := taskPb.NewTaskServiceClient(connTask)

	// 加入熔断 TODO main太臃肿了
	// wrapper.NewServiceWrapper(userServiceName)
	// wrapper.NewServiceWrapper(taskServiceName)

	ginRouter := routes.NewRouter(userService, taskService)
	server := &http.Server{
		Addr:           config.Conf.Server.Port,
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}
	go func() {
		// 优雅关闭
		shutdown.GracefullyShutdown(server)
	}()
}

func RPCConnect(ctx context.Context, serviceName string, etcdRegister *discovery.Resolver) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	addr := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}
