package logic

import (
	"context"
	"rubbish-cart/rpc/user/user"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserListLogic {
	return GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList() (*types.GetUserListResp, error) {
	userRet, err := l.svcCtx.UserRpc.GetUserList(l.ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}

	var userResp []types.UserInfo
	for _, item := range userRet.UserinfoList {
		userInfo := types.UserInfo{
			OpenId:  item.OpenId,
			Name:    item.Name,
			Purview: int(item.Purview),
		}

		userResp = append(userResp, userInfo)
	}

	return &types.GetUserListResp{
		UserList: userResp,
		Resp: types.Resp{
			State: 1,
			Msg: "获取成功!",
		},
	}, nil
}
