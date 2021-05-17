package logic

import (
	"context"
	"rubbish-cart/rpc/user/model"

	"rubbish-cart/rpc/user/internal/svc"
	"rubbish-cart/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegisterLogic {
	return &GetRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRegisterLogic) GetRegister(in *user.GetRegisterReq) (*user.GetRegisterResp, error) {
	_, err := l.svcCtx.Model.Insert(model.User{
		OpenId: in.Openid,
		UserName: in.Name,
		Purview: 2,
	})
	if err == nil {
		return &user.GetRegisterResp{
			Purview: 2,
		}, nil
	}

	if err.Error() == "Error 1062: Duplicate entry '" + in.Openid + "' for key 'PRIMARY'" {
		res, err := l.svcCtx.Model.FindOne(in.Openid)
		if err != nil {
			return nil, err
		}

		return &user.GetRegisterResp{
			Purview: res.Purview,
		}, nil
	}

	return nil, err
}
