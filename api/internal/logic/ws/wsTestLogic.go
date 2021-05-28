package logic

import (
	"context"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WsTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) WsTestLogic {
	return WsTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsTestLogic) WsTest(req types.WsReq) (*types.WsResp, error) {
	// todo: add your logic here and delete this line

	return &types.WsResp{}, nil
}
