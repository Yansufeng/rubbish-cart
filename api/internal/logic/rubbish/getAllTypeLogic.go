package logic

import (
	"context"
	"rubbish-cart/rpc/rubbish/rubbish"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllTypeLogic {
	return GetAllTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllTypeLogic) GetAllType() (*types.GetAllResp, error) {
	typeReq, err := l.svcCtx.RubbishRpc.GetTypeAll(l.ctx, &rubbish.Empty{})
	if err != nil {
		return &types.GetAllResp{
			TypeList: nil,
			Resp: types.Resp{
				State: 0,
				Msg:   err.Error(),
			},
		}, nil
	}

	var res []types.RubbishType
	for _, item := range typeReq.TypeList {
		typeInfo := types.RubbishType{
			TypeId: int(item.TypeId),
			Name:   item.Name,
			IconId: int(item.IconId),
		}
		res = append(res, typeInfo)
	}

	return &types.GetAllResp{
		TypeList: res,
		Resp: types.Resp{
			State: 1,
			Msg:   "获取成功!",
		},
	}, nil
}
