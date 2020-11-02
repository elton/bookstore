package logic

import (
	"context"

	"github.com/elton/bookstore/api/internal/svc"
	"github.com/elton/bookstore/api/internal/types"
	"github.com/elton/bookstore/rpc/check/checker"

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
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckRequest{
		Book: req.Book,
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckResponse{
		Found: resp.Found,
		Price: resp.Price,
	}, nil
}
