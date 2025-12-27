package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gw-gong/go-template-project/api/rpc/svc02"
	"github.com/gw-gong/go-template-project/internal/app/svc02/svc/test01"
	"github.com/gw-gong/go-template-project/internal/app/svc02/svc/test02"
	"github.com/gw-gong/go-template-project/internal/config/svc02/localcfg"

	"github.com/gw-gong/gwkit-go/grpc/interceptor/server/unary"
	"github.com/gw-gong/gwkit-go/log"
	"github.com/gw-gong/gwkit-go/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type RpcServer struct {
	cfg       *localcfg.Config
	test01Svc *test01.Test01Svc
	test02Svc *test02.Test02Svc
}

func (s *RpcServer) Run(ctx context.Context) {
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
