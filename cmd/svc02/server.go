package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gw-gong/boilerplate-go/api/rpc/svc02"
	"github.com/gw-gong/boilerplate-go/internal/app/svc02/svc/test01"
	"github.com/gw-gong/boilerplate-go/internal/app/svc02/svc/test02"
	"github.com/gw-gong/boilerplate-go/internal/config/svc02/localcfg"
	"github.com/gw-gong/boilerplate-go/internal/pkg/util/consul"

	gwkitconsul "github.com/gw-gong/gwkit-go/grpc/consul"
	"github.com/gw-gong/gwkit-go/grpc/interceptor/server/unary"
	"github.com/gw-gong/gwkit-go/hotcfg"
	"github.com/gw-gong/gwkit-go/log"
	"github.com/gw-gong/gwkit-go/setting"
	"github.com/gw-gong/gwkit-go/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type RpcServer struct {
	cfg          *localcfg.Config
	hlm          hotcfg.HotLoaderManager
	consulClient gwkitconsul.ConsulClient
	test01Svc    *test01.Test01Svc
	test02Svc    *test02.Test02Svc
}

func (s *RpcServer) SetupAndRun(ctx context.Context) {
	setting.SetEnv(s.cfg.Env)

	// init global logger
	syncFn, err := log.InitGlobalLogger(s.cfg.Logger)
	util.ExitOnErr(ctx, err)
	defer syncFn()

	// start hot reload
	util.ExitOnErr(ctx, s.hlm.RegisterHotLoader(s.cfg))
	util.ExitOnErr(ctx, s.hlm.Watch())

	// register services
	deregister, err := consul.RegisterServices(s.consulClient, s.cfg.RpcServer.RegisterEntries, s.cfg.RpcServer.Port)
	util.ExitOnErr(ctx, err)
	defer deregister()

	// init grpc server
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			unary.PanicRecoverInterceptor(),
			unary.ParseMetaToCtx(),
		),
	)

	// register health check service
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	// register biz services
	svc02.RegisterTest01ServiceServer(grpcServer, s.test01Svc)
	svc02.RegisterTest02ServiceServer(grpcServer, s.test02Svc)

	// start rpc server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.RpcServer.Port))
	util.ExitOnErr(ctx, err)
	defer listener.Close()

	log.Infoc(ctx, "rpc server running", log.Int("port", s.cfg.RpcServer.Port))
	err = grpcServer.Serve(listener)
	util.ExitOnErr(ctx, err)
}
