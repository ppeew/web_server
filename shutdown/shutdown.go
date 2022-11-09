package shutdown

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"web_server/server"
)

type Hook func(c context.Context) error

func WaitForShutdown(hooks ...Hook) {
	signalCh := make(chan os.Signal, 1)
	shutdownSignals := []os.Signal{
		os.Interrupt, os.Kill, syscall.SIGKILL,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
		syscall.SIGABRT, syscall.SIGTERM,
	}
	signal.Notify(signalCh, shutdownSignals...)
	select {
	case sig := <-signalCh:
		fmt.Printf("recive %s signal\n", sig)
		time.AfterFunc(time.Minute, func() {
			fmt.Println("exit because timeout 1 minute")
			os.Exit(1)
		})
		//执行业务退出处理
		for _, hook := range hooks {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			err := hook(ctx)
			if err != nil {
				fmt.Printf("hook problem:%s\n", err)
			}
			cancel()
		}
		//成功退出
		os.Exit(0)
	}
}

var _ Hook = BuildCloseServersHook()

func BuildCloseServersHook(servers ...server.Server) Hook {
	return func(c context.Context) error {
		wg := sync.WaitGroup{}
		doneCh := make(chan struct{}, 1)
		for _, s := range servers {
			wg.Add(1)
			//执行server退出
			go func(ser server.Server) {
				err := ser.Shutdown(c)
				if err != nil {
					fmt.Printf("server close failed : %s", err)
				}
				wg.Done()
			}(s)
		}
		go func() {
			//检查是不是全部完成了
			wg.Wait()
			doneCh <- struct{}{}
		}()
		select {
		case <-doneCh:
			fmt.Println("all server closed")
			return nil
		case <-c.Done():
			return errors.New("closing servers timeout")
		}

	}
}
