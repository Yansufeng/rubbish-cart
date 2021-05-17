package logic

import (
	"context"
	"rubbish-cart/rpc/rubbish/rubbish"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCartRubbishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartRubbishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCartRubbishListLogic {
	return GetCartRubbishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartRubbishListLogic) GetCartRubbishList(req types.GetCartRubbishListReq) (*types.GetCartRubbishListResp, error) {
	rubbishReq, err := l.svcCtx.RubbishRpc.GetCartRubbishList(l.ctx, &rubbish.GetCartRubbishListReq{
		CartId: int64(req.CartId),
	})
	if err != nil {
		return &types.GetCartRubbishListResp{
			RubbishList: nil,
			Resp: types.Resp{
				State: 0,
				Msg: err.Error(),
			},
		}, nil
	}

	var res []types.RubbishInfoResp
	for _, item := range rubbishReq.RubbishList {
		rubbishInfo := types.RubbishInfoResp{
			CartId: int(item.CartId),
			TypeId: int(item.TypeId),
			Name:   item.Name,
			Num:    int(item.Num),
		}
		res = append(res, rubbishInfo)
	}

	return &types.GetCartRubbishListResp{
		RubbishList: res,
		Resp: types.Resp{
			State: 1,
			Msg: "获取成功！",
		},
	}, nil
}
