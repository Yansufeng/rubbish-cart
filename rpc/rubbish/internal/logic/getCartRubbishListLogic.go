package logic

import (
	"context"

	"rubbish-cart/rpc/rubbish/internal/svc"
	"rubbish-cart/rpc/rubbish/rubbish"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCartRubbishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartRubbishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartRubbishListLogic {
	return &GetCartRubbishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartRubbishListLogic) GetCartRubbishList(in *rubbish.GetCartRubbishListReq) (*rubbish.GetCartRubbishListResp, error) {
	rubbishList, err := l.svcCtx.Model.GetAllByCartId(in.CartId)
	if err != nil {
		return nil, err
	}

	typeMap, err := l.svcCtx.TypeModel.GetAll()
	if err != nil {
		return nil, err
	}

	var res []*rubbish.RubbishInfo
	for _, item := range rubbishList {
		rubbishInfo := rubbish.RubbishInfo{
			CartId: item.CartId,
			TypeId: item.RubbishType,
			Name: typeMap[item.RubbishType],
			Num: item.RubbishAmount,
		}

		res = append(res, &rubbishInfo)
	}

	return &rubbish.GetCartRubbishListResp{
		RubbishList: res,
	}, nil
}
