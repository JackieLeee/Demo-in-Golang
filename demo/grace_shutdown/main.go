package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个 HTTP 服务器
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}
	var isShutDown = make(chan struct{})
	server.RegisterOnShutdown(
		func() {
			isShutDown <- struct{}{}
		},
	)

	// 启动 HTTP 服务器
	go func() {
		log.Println("Starting server...")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	wait4ShunDown(server, isShutDown)
}

func wait4ShunDown(server *http.Server, isShutDown chan struct{}) {
	// 创建一个信号通道，用于接收操作系统的中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// 等待中断信号
	<-sigChan

	// 创建一个定时器，用于优雅关闭服务器
	shutdownTimeout := 5 * time.Second
	timeout := time.NewTimer(shutdownTimeout)

	// 关闭服务器
	log.Println("Shutting down server...")
	err := server.Shutdown(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 等待服务器关闭或超时
	select {
	case <-timeout.C:
		log.Println("Shutdown timeout. Forcefully exiting...")
		os.Exit(1)
	case <-isShutDown:
		log.Println("Server gracefully stopped.")
	}
}

// 处理 HTTP 请求的处理程序
func handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello, World!")
}
