package logic

import (
	"context"
	"rubbish-cart/rpc/rubbish/rubbish"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateCartRubbishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartRubbishLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateCartRubbishLogic {
	return UpdateCartRubbishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartRubbishLogic) UpdateCartRubbish(req types.UpdateCartRubbishReq) (*types.UpdateCartRubbishResp, error) {
	for _, item := range req.RubbishList {
		res, err := l.svcCtx.RubbishRpc.UpdateCartRubbish(l.ctx, &rubbish.UpdateCartRubbishReq{
			CartId: int64(req.CartId),
			RubbishInfo: &rubbish.RubbishInfo{
				CartId: int64(item.CartId),
				TypeId: int64(item.TypeId),
				Name:   item.Name,
				Num:    int64(item.Num),
			},
		})
		if err != nil {
			return &types.UpdateCartRubbishResp{
				Resp: types.Resp{
					State: 0,
					Msg:   err.Error(),
				},
			}, nil
		}
		if !res.Ok {
			return &types.UpdateCartRubbishResp{
				Resp: types.Resp{
					State: 0,
					Msg:   "更新失败！",
				},
			}, nil
		}
	}

	return &types.UpdateCartRubbishResp{
		Resp: types.Resp{
			State: 1,
			Msg:   "更新成功！",
		},
	}, nil
}
