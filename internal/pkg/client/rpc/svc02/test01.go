package svc02

import (
	"context"
	"errors"
	"fmt"

	"github.com/gw-gong/go-template-project/api/rpc/svc02"

	"github.com/gw-gong/gwkit-go/grpc/consul"
)

type Test01ClientOption consul.HealthyGrpcConnOption

type Test01Client struct {
	svc02.Test01ServiceClient
}

func NewTest01Client(option *Test01ClientOption) (*Test01Client, error) {
	if option == nil {
		return nil, errors.New("option is nil")
	}

	conn, err := consul.NewHealthyGrpcConn((*consul.HealthyGrpcConnOption)(option))
	if err != nil {
		return nil, err
	}

	return &Test01Client{svc02.NewTest01ServiceClient(conn)}, nil
}

func (c *Test01Client) TestFunc(ctx context.Context, field01 string, field02 string) (resField01 string, resField02 string, err error) {
	req := &svc02.Test01Request{
		Field01: field01,
		Field02: field02,
	}

	resp, err := c.Test01ServiceClient.TestFunc(ctx, req)
	if err != nil {
		return "", "", fmt.Errorf("failed to call TestFunc: %w", err)
	}

	return resp.Field01, resp.Field02, nil
}
