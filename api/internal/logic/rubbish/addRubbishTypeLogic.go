package logic

import (
	"context"
	"rubbish-cart/rpc/rubbish/rubbish"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddRubbishTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRubbishTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddRubbishTypeLogic {
	return AddRubbishTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRubbishTypeLogic) AddRubbishType(req types.AddRubbishTypeReq) (*types.AddRubbishTypeResp, error) {
	res, err := l.svcCtx.RubbishRpc.AddRubbishType(l.ctx, &rubbish.AddRubbishTypeReq{
		Name: req.Name,
	})
	if err != nil {
		return &types.AddRubbishTypeResp{
			Id: 0,
			Resp: types.Resp{
				State: 0,
				Msg: err.Error(),
			},
		}, nil
	}

	return &types.AddRubbishTypeResp{
		Id: int(res.Id),
		Resp: types.Resp{
			State: 1,
			Msg: "增加垃圾类别成功！",
		},
	}, nil
}
