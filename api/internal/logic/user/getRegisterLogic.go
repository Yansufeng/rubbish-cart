package logic

import (
	"context"
	"rubbish-cart/rpc/user/user"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRegisterLogic {
	return GetRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegisterLogic) GetRegister(req types.GetRegisterReq) (*types.GetRegisterResp, error) {
	res, err := l.svcCtx.UserRpc.GetRegister(l.ctx, &user.GetRegisterReq{
		Openid: req.OpenId,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetRegisterResp{
		Purview: int(res.Purview),
		Resp: types.Resp{
			State: 1,
			Msg: "获取权限成功！",
		},
	}, nil
}