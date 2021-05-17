package logic

import (
	"context"
	"rubbish-cart/rpc/cart/cartclient"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartRegisterLogic {
	return CartRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartRegisterLogic) CartRegister(req types.CartRegisterReq) (*types.CartRegisterResp, error) {
	resp, err := l.svcCtx.CartRpc.CartRegister(l.ctx, &cartclient.CartRegisterReq{
		CartName: req.CartName,
	})
	if err != nil {
		return &types.CartRegisterResp{
			CartId: 0,
			Resp: types.Resp{
				State: 0,
				Msg: err.Error(),
			},
		}, err
	}
	if resp == nil {
		return &types.CartRegisterResp{
			CartId: 0,
			Resp: types.Resp{
				State: 0,
				Msg: "注册失败！",
			},
		}, err
	}
	if resp.Ok {
		return &types.CartRegisterResp{
			CartId: int(resp.CartId),
			Resp: types.Resp{
				State: 1,
				Msg: "注册成功!",
			},
		}, nil
	}

	return &types.CartRegisterResp{
		CartId: 0,
		Resp: types.Resp{
			State: 2,
			Msg: "命名重复",
		},
	}, err
}
