package main

import (
	"fmt"
	"github.com/ZuoFuhong/go_wallet_manage_svr/app/interfaces"
	"github.com/ZuoFuhong/go_wallet_manage_svr/pkg/config"
	glog "github.com/ZuoFuhong/go_wallet_manage_svr/pkg/log"
	_ "github.com/ZuoFuhong/grpc-middleware/encoding/json"
	ngm "github.com/ZuoFuhong/grpc-middleware/tracing"
	"github.com/ZuoFuhong/grpc-naming-monica/registry"
	pb "github.com/ZuoFuhong/grpc-standard-pb/go_wallet_manage_svr"
	gm "github.com/grpc-ecosystem/go-grpc-middleware"
	gr "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

	"log"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	serviceImpl := interfaces.InitializeService()
	s := grpc.NewServer(grpc.UnaryInterceptor(gm.ChainUnaryServer(ngm.UnaryServerInterceptor(), gr.UnaryServerInterceptor())))
	pb.RegisterGoWalletManageSvrServer(s, serviceImpl)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("load config fail: " + err.Error())
	}
	config.SetGlobalConfig(cfg)

	// 服务注册
	if err := registry.NewRegistry(&registry.Config{
		Token:       cfg.Monica.Token,
		Namespace:   cfg.Monica.Namespace,
		ServiceName: cfg.Monica.ServiceName,
		IP:          cfg.Server.Addr,
		Port:        cfg.Server.Port,
	}).Register(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.Addr, cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
	glog.Debugf("Serving %s on %s", cfg.Server.Name, lis.Addr().String())
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
