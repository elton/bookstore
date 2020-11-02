package logic

import (
	"context"

	"github.com/elton/bookstore/rpc/check/internal/svc"
	check "github.com/elton/bookstore/rpc/check/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckRequest) (*check.CheckResponse, error) {
	resp, err := l.svcCtx.Model.FindOne(in.Book)
	if err != nil {
		return nil, err
	}

	return &check.CheckResponse{
		Found: true,
		Price: resp.Price,
	}, nil
}
