package logic

import (
	"context"

	"rubbish-cart/rpc/cart/cart"
	"rubbish-cart/rpc/cart/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateLogic {
	return &CartUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartUpdateLogic) CartUpdate(in *cart.CartUpdateReq) (*cart.CartUpdateResp, error) {
	cartRet, err := l.svcCtx.Model.FindOne(in.CartId)
	if err != nil {
		return &cart.CartUpdateResp{
			Ok: false,
		}, err
	}

	cartRet.State = in.State

	err = l.svcCtx.Model.Update(*cartRet)
	if err != nil {
		return &cart.CartUpdateResp{
			Ok: false,
		}, err
	}

	return &cart.CartUpdateResp{
		Ok: true,
	}, nil
}
