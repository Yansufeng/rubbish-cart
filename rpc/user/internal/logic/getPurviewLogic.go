package logic

import (
	"context"

	"rubbish-cart/rpc/user/internal/svc"
	"rubbish-cart/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetPurviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPurviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPurviewLogic {
	return &GetPurviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPurviewLogic) GetPurview(in *user.GetPurviewReq) (*user.GetPurviewResp, error) {
	res, err := l.svcCtx.Model.FindOne(in.Openid)
	if err != nil {
		return nil, err
	}

	return &user.GetPurviewResp{
		Purview: res.Purview,
	}, nil
}
