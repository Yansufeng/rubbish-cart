package logic

import (
	"context"

	"rubbish-cart/rpc/rubbish/internal/svc"
	"rubbish-cart/rpc/rubbish/rubbish"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTypeAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTypeAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeAllLogic {
	return &GetTypeAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTypeAllLogic) GetTypeAll(in *rubbish.Empty) (*rubbish.GetTypeAllResp, error) {
	typeList, err := l.svcCtx.TypeModel.GetAllToList()
	if err != nil {
		return nil, err
	}

	var typeResp []*rubbish.RubbishType
	for _, item := range typeList {
		typeInfo := rubbish.RubbishType{
			TypeId: item.RubbishType,
			Name:   item.RubbishName,
			IconId: item.IconId,
		}
		typeResp = append(typeResp, &typeInfo)
	}

	return &rubbish.GetTypeAllResp{
		TypeList: typeResp,
	}, nil
}
