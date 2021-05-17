package logic

import (
	"context"
	"rubbish-cart/rpc/cart/cart"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetInfoLogic {
	return GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo() (*types.GetInfoResp, error) {
	cartRet, err := l.svcCtx.CartRpc.GetInfo(l.ctx, &cart.Empty{})
	if err != nil {
		return &types.GetInfoResp{
			CartInfoList: nil,
			Resp: types.Resp{
				State: 0,
				Msg: err.Error(),
			},
		}, nil
	}

	var cartResp []types.CartInfo
	for _, item := range cartRet.CartList {
		cartInfo := types.CartInfo{
			CartId: int(item.CartId),
			CartName: item.CartName,
			State:   item.State,
		}

		cartResp = append(cartResp, cartInfo)
	}

	return &types.GetInfoResp{
		CartInfoList: cartResp,
		Resp: types.Resp{
			State: 1,
			Msg: "获取成功！",
		},
	}, nil
}
