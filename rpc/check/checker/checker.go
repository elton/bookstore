// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

//go:generate mockgen -destination ./checker_mock.go -package checker -source $GOFILE

package checker

import (
	"context"

	check "github.com/elton/bookstore/rpc/check/pb"

	"github.com/tal-tech/go-zero/core/jsonx"
	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Checker interface {
		Check(ctx context.Context, in *CheckRequest) (*CheckResponse, error)
	}

	defaultChecker struct {
		cli zrpc.Client
	}
)

func NewChecker(cli zrpc.Client) Checker {
	return &defaultChecker{
		cli: cli,
	}
}

func (m *defaultChecker) Check(ctx context.Context, in *CheckRequest) (*CheckResponse, error) {
	var request check.CheckRequest
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := check.NewCheckerClient(m.cli.Conn())
	resp, err := client.Check(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret CheckResponse
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}
