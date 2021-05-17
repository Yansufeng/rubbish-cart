package logic

import (
	"context"
	"rubbish-cart/rpc/user/user"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOpenIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOpenIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOpenIdLogic {
	return GetOpenIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOpenIdLogic) GetOpenId(req types.GetOpenIdReq) (*types.GetOpenIdResp, error) {
	res, err := l.svcCtx.UserRpc.GetOpenId(l.ctx, &user.GetOpenIdReq{
		Code: req.Code,
	})
	if err != nil {
		return nil, err
	}

	if res == nil {
		return &types.GetOpenIdResp{
			OpenId: "",
			Resp: types.Resp{
				State: 0,
				Msg: "获取OpenId失败！",
			},
		}, nil
	}

	return &types.GetOpenIdResp{
		OpenId: res.OpenId,
		Resp: types.Resp{
			State: 1,
			Msg: "获取OpenId成功！",
		},
	}, nil
}
