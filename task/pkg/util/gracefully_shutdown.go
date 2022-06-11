package util

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func GracefullyShutdown(server *http.Server) {
	// 创建系统信号接收器接收关闭信号
	done := make(chan os.Signal, 1)
	/**
	os.Interrupt           -> ctrl+c 的信号
	syscall.SIGINT|SIGTERM -> kill 进程时传递给进程的信号
	*/
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	LogrusObj.Println("closing http server gracefully ...")

	if err := server.Shutdown(context.Background()); err != nil {
		LogrusObj.Fatalln("closing http server gracefully failed: ", err)
	}

}
