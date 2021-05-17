package logic

import (
	"context"
	"rubbish-cart/rpc/user/user"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetPurviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPurviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPurviewLogic {
	return GetPurviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPurviewLogic) GetPurview(req types.GetPurviewReq) (*types.GetPurviewResp, error) {
	res, err := l.svcCtx.UserRpc.GetPurview(l.ctx, &user.GetPurviewReq{
		Openid: req.OpenId,
	})
	if err != nil {
		return &types.GetPurviewResp{
			Purview: 0,
			Resp: types.Resp{
				State: 0,
				Msg: "获取权限失败！",
			},
		}, err
	}

	return &types.GetPurviewResp{
		Purview: int(res.Purview),
		Resp: types.Resp{
			State: 1,
			Msg: "获取权限成功！",
		},
	}, nil
}
