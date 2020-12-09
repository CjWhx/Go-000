package main

import (
	"context"
	"fmt"
	"github.com/golang/sync/errgroup"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/
var wg sync.WaitGroup
var stop = make(chan struct{}, 0)
var sigs = make(chan os.Signal, 1)
var over = make(chan struct{}, 0)

func main() {

	done := make(chan error, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		group, _ := errgroup.WithContext(context.Background())
		group.Go(func() error {
			err := serve("127.0.0.1:9090", myHandler{}, stop)
			return err
		})
		err1 := group.Wait()
		done <- err1
	}()

	go func() {
		defer wg.Done()
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigs:
		case <-over:
		}
		fmt.Println("signal goroutine: get exit signal")
		done <- errors.New("接收到信号，程序需要中断")
	}()

	for i := 0; i < cap(done); i++ {
		err := <-done
		if err != nil {
			fmt.Println("main goroutine: 接收的错误:", err)
			close(over)
			break
		}
	}
	wg.Wait()
	fmt.Println("程序结束运行...")

}

func serve(addr string, handler http.Handler, stop chan struct{}) error {
	defer fmt.Println("http server closed")

	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		select {
		case <-stop:
		case <-over:
		}
		fmt.Println("server goroutine: close service ")
		s.Shutdown(context.Background())

	}()

	return s.ListenAndServe()
}

type myHandler struct {
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi, let's golang"))

}
