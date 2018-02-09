package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"path/filepath"
	"os"
	"github.com/ADreamean/ddz-backend/login"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"github.com/ADreamean/ddz-backend/until/redis"
	"github.com/ADreamean/ddz-backend/room"
)

const VERSION = "0.1.0"

func main() {
	a := kingpin.New(filepath.Base(os.Args[0]), "斗地主服务器")
	a.Version(VERSION)

	cfg := struct {
		login        login.Option
		redisAddress string
	}{}

	a.HelpFlag.Short('h')

	a.Flag("login.listen_address", "登陆系统监听地址").
		Default("0.0.0.0:9090").StringVar(&cfg.login.ListenAddress)

	a.Flag("redis.address", "redis连接地址").
		Default(":6379").StringVar(&cfg.redisAddress)

	_, err := a.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		a.Usage(os.Args[1:])
		os.Exit(2)
	}
	redis.Init(cfg.redisAddress)
	grsv := grpc.NewServer()
	login.RegisterGRPC(grsv)
	room.RegisterGRPC(grsv)
	lis, err := net.Listen("tcp", cfg.login.ListenAddress)
	grsv.Serve(lis)
}
