package logic

import (
	"context"
	"rubbish-cart/rpc/cart/cart"
	"rubbish-cart/rpc/cart/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *cart.Empty) (*cart.GetInfoResp, error) {
	cartList, err := l.svcCtx.Model.GetAll()
	if err != nil {
		return nil, err
	}

	var cartResp []*cart.CartInfo
	for _, item := range cartList {
		cartInfo := cart.CartInfo{
			CartId: item.CartId,
			CartName: item.CartName,
			State: item.State,
		}
		cartResp = append(cartResp, &cartInfo)
	}

	return &cart.GetInfoResp{
		CartList: cartResp,
	}, nil
}
