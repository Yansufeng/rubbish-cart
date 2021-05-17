package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"rubbish-cart/rpc/cart/model"

	"rubbish-cart/rpc/cart/cart"
	"rubbish-cart/rpc/cart/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartRegisterLogic {
	return &CartRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartRegisterLogic) CartRegister(in *cart.CartRegisterReq) (*cart.CartRegisterResp, error) {
	cartRet, err := l.svcCtx.Model.FindOneByName(in.CartName)
	fmt.Println("31--FindOneByName: ", cartRet, err)
	if err != nil && err != sqlx.ErrNotFound {
		return &cart.CartRegisterResp{
			CartId: 0,
			Ok: false,
		}, err
	}

	if err != sqlx.ErrNotFound {
		return &cart.CartRegisterResp{
			CartId: 0,
			Ok: false,
		}, nil
	}

	cacheCartCartNamePrefix := "cache#cart#cartName#" + in.CartName
	err = l.svcCtx.Redis.DelOneByKey(cacheCartCartNamePrefix)
	fmt.Println("48--DelOneByKey's err: ", err)
	ret, err := l.svcCtx.Model.Insert(model.Cart{
		CartName: in.CartName,
		State: "正常",
	})

	if err != nil {
		return &cart.CartRegisterResp{
			CartId: 0,
			Ok: false,
		}, err
	}

	id, _ := ret.LastInsertId()
	return &cart.CartRegisterResp{
		CartId: id,
		Ok: true,
	}, nil
}
