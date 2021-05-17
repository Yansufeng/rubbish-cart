package logic

import (
	"context"
	"rubbish-cart/rpc/cart/cart"
	"rubbish-cart/rpc/cart/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlarmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmLogic {
	return &AlarmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlarmLogic) Alarm(in *cart.AlarmReq) (*cart.AlarmResp, error) {
	cartRet, err := l.svcCtx.Model.FindOne(in.CartId)
	if err != nil {
		return &cart.AlarmResp{
			Ok: false,
		}, err
	}

	cartRet.State = in.Msg

	err = l.svcCtx.Model.Update(*cartRet)
	if err != nil {
		return &cart.AlarmResp{
			Ok: false,
		}, err
	}

	return &cart.AlarmResp{
		Ok: true,
	}, nil
}
