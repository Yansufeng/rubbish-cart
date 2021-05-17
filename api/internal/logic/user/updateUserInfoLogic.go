package logic

import (
	"context"
	"rubbish-cart/rpc/user/user"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserInfoLogic {
	return UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req types.UpdateUserInfoReq) (*types.UpdateUserInfoResp, error) {
	_, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &user.UpdateUserInfoReq{
		OpenId:  req.OpenId,
		Name:    req.Name,
		Purview: int64(req.Purview),
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateUserInfoResp{
		Resp: types.Resp{
			State: 1,
			Msg: "更新成功！",
		},
	}, nil
}
