package logic

import (
	"context"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) WsLogic {
	return WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws(req types.WsReq) (*types.WsResp, error) {
	// todo: add your logic here and delete this line

	return &types.WsResp{}, nil
}
