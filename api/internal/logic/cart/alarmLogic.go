package logic

import (
	"context"
	"rubbish-cart/rpc/cart/cart"

	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlarmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlarmLogic {
	return AlarmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmLogic) Alarm(req types.AlarmReq) (*types.AlarmResp, error) {
	alarmRet, err := l.svcCtx.CartRpc.Alarm(l.ctx, &cart.AlarmReq{
		CartId: int64(req.CartId),
		Msg: req.Msg,
	})
	if err != nil {
		return &types.AlarmResp{
			Resp: types.Resp{
				State: 0,
				Msg: "发布警报信息失败!",
			},
		}, err
	}

	if !alarmRet.Ok {
		return &types.AlarmResp{
			Resp: types.Resp{
				State: 0,
				Msg: "发布警报信息失败!",
			},
		}, nil
	}

	return &types.AlarmResp{
		Resp: types.Resp{
			State: 1,
			Msg: "发布警报信息成功!",
		},
	}, nil
}
