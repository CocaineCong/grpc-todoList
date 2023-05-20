package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CocaineCong/grpc-todolist/app/gateway/routes"
	"github.com/CocaineCong/grpc-todolist/app/gateway/rpc"
	"github.com/CocaineCong/grpc-todolist/config"
	"github.com/CocaineCong/grpc-todolist/pkg/util/shutdown"
)

func main() {
	config.InitConfig()
	rpc.Init()

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
	// 加入熔断 TODO main太臃肿了
	// wrapper.NewServiceWrapper(userServiceName)
	// wrapper.NewServiceWrapper(taskServiceName)

	r := routes.NewRouter()
	server := &http.Server{
		Addr:           config.Conf.Server.Port,
		Handler:        r,
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
