package logic

import (
	"context"
	"rubbish-cart/rpc/rubbish/model"

	"rubbish-cart/rpc/rubbish/internal/svc"
	"rubbish-cart/rpc/rubbish/rubbish"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddRubbishTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRubbishTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRubbishTypeLogic {
	return &AddRubbishTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRubbishTypeLogic) AddRubbishType(in *rubbish.AddRubbishTypeReq) (*rubbish.AddRubbishTypeResp, error) {
	res, err := l.svcCtx.TypeModel.Insert(model.Rubbishtype{
		RubbishName: in.Name,
		IconId:      4,
	})
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &rubbish.AddRubbishTypeResp{
		Id: id,
	}, nil
}
