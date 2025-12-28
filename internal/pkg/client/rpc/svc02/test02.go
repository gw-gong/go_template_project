package svc02

import (
	"context"
	"errors"
	"fmt"

	"github.com/gw-gong/go-template-project/api/rpc/svc02"

	"github.com/gw-gong/gwkit-go/grpc/consul"
)

type Test02ClientOption consul.HealthyGrpcConnEntry

type Test02Client struct {
	svc02.Test02ServiceClient
}

func NewTest02Client(consulClient consul.ConsulClient, option *Test02ClientOption) (*Test02Client, error) {
	if option == nil {
		return nil, errors.New("option is nil")
	}

	conn, err := consulClient.GetHealthyGrpcConn((*consul.HealthyGrpcConnEntry)(option))
	if err != nil {
		return nil, err
	}

	return &Test02Client{svc02.NewTest02ServiceClient(conn)}, nil
}

func (c *Test02Client) TestFunc(ctx context.Context, field01 string, field02 string) (resField01 string, resField02 string, err error) {
	req := &svc02.Test02Request{
		Field01: field01,
		Field02: field02,
	}

	resp, err := c.Test02ServiceClient.TestFunc(ctx, req)
	if err != nil {
		return "", "", fmt.Errorf("failed to call TestFunc: %w", err)
	}

	return resp.Field01, resp.Field02, nil
}
