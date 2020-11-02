package logic

import (
	"context"

	"github.com/elton/bookstore/api/internal/svc"
	"github.com/elton/bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckRequest) (*types.CheckResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CheckResponse{}, nil
}
