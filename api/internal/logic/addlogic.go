package logic

import (
	"context"

	"github.com/elton/bookstore/api/internal/svc"
	"github.com/elton/bookstore/api/internal/types"
	"github.com/elton/bookstore/rpc/add/adder"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddRequest) (*types.AddResponse, error) {
	resp, err := l.svcCtx.Adder.Add(l.ctx, &adder.AddRequest{
		Book:  req.Book,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResponse{
		Ok: resp.Ok,
	}, nil
}
