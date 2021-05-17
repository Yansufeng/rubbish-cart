package logic

import (
	"context"
	"fmt"

	"rubbish-cart/rpc/rubbish/internal/svc"
	"rubbish-cart/rpc/rubbish/rubbish"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateCartRubbishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartRubbishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartRubbishLogic {
	return &UpdateCartRubbishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartRubbishLogic) UpdateCartRubbish(in *rubbish.UpdateCartRubbishReq) (*rubbish.UpdateCartRubbishResp, error) {
	res, err := l.svcCtx.Model.FindOneById(in.CartId, in.RubbishInfo.TypeId)
	fmt.Println("---RPC UpdateCartRubbish() RES : ", res)
	fmt.Println("---RPC UpdateCartRubbish() ERROR : ", err)

	return &rubbish.UpdateCartRubbishResp{}, nil
}
