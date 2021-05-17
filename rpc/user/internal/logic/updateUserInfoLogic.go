package logic

import (
	"context"
	"rubbish-cart/rpc/user/model"

	"rubbish-cart/rpc/user/internal/svc"
	"rubbish-cart/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoReq) (*user.UpdateUserInfoResp, error) {
	err := l.svcCtx.Model.Update(model.User{
		OpenId: in.OpenId,
		UserName: in.Name,
		Purview: in.Purview,
	})
	if err != nil {
		return &user.UpdateUserInfoResp{
			Ok: false,
		}, err
	}

	return &user.UpdateUserInfoResp{
		Ok: true,
	}, nil
}
