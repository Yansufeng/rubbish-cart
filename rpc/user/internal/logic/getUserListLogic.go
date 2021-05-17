package logic

import (
	"context"

	"rubbish-cart/rpc/user/internal/svc"
	"rubbish-cart/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.Empty) (*user.GetUserListResp, error) {
	userList, err := l.svcCtx.Model.GetAll()
	if err != nil {
		return nil, err
	}

	var userResp []*user.UserInfo
	for _, item := range userList {
		userInfo := user.UserInfo{
			OpenId: item.OpenId,
			Name: item.UserName,
			Purview: item.Purview,
		}
		userResp = append(userResp, &userInfo)
	}

	return &user.GetUserListResp{
		UserinfoList: userResp,
	}, nil
}
