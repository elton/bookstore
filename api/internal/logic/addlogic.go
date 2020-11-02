package logic

import (
	"context"

	"github.com/elton/bookstore/api/internal/svc"
	"github.com/elton/bookstore/api/internal/types"

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

func (l *AddLogic) Add(req types.AddRequst) (*types.AddResponse, error) {
	// todo: add your logic here and delete this line

	return &types.AddResponse{}, nil
}
