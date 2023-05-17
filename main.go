package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	conf "oos/config"
	ctl "oos/controller"
	"oos/model"
	rt "oos/router"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"context"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	//model 모듈 선언
	if mod, err := model.NewModel(); err != nil {
		panic(err)
	} else if controller, err := ctl.NewCTL(mod); err != nil { //controller 모듈 설정
		panic(err)
	} else if router, err := rt.NewRouter(controller); err != nil { //router 모듈 설정
		panic(err)
	} else {
		config := conf.GetConfig("./config/config.toml")

		mapi := &http.Server{
			Addr:           config.Server.Port,
			Handler:        router.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		
		
		g.Go(func() error {
				return mapi.ListenAndServe()
		})	

		//우아한 종료
		stopSig := make(chan os.Signal) //chan 선언
		// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
		signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM) 
		<-stopSig //메세지 등록

		// 해당 context 타임아웃 설정, 5초후 server stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
			case <-ctx.Done():
				fmt.Println("timeout 5 seconds.")
		}
		fmt.Println("Server stop")
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}