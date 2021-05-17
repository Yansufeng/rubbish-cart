// Code generated by goctl. DO NOT EDIT!
// Source: rubbish.proto

//go:generate mockgen -destination ./rubbish_mock.go -package rubbishclient -source $GOFILE

package rubbishclient

import (
	"context"

	"rubbish-cart/rpc/rubbish/rubbish"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	AddRubbishTypeResp     = rubbish.AddRubbishTypeResp
	UpdateCartRubbishReq   = rubbish.UpdateCartRubbishReq
	UpdateCartRubbishResp  = rubbish.UpdateCartRubbishResp
	RubbishInfo            = rubbish.RubbishInfo
	GetCartRubbishListReq  = rubbish.GetCartRubbishListReq
	GetCartRubbishListResp = rubbish.GetCartRubbishListResp
	AddRubbishTypeReq      = rubbish.AddRubbishTypeReq

	Rubbish interface {
		GetCartRubbishList(ctx context.Context, in *GetCartRubbishListReq) (*GetCartRubbishListResp, error)
		AddRubbishType(ctx context.Context, in *AddRubbishTypeReq) (*AddRubbishTypeResp, error)
		UpdateCartRubbish(ctx context.Context, in *UpdateCartRubbishReq) (*UpdateCartRubbishResp, error)
	}

	defaultRubbish struct {
		cli zrpc.Client
	}
)

func NewRubbish(cli zrpc.Client) Rubbish {
	return &defaultRubbish{
		cli: cli,
	}
}

func (m *defaultRubbish) GetCartRubbishList(ctx context.Context, in *GetCartRubbishListReq) (*GetCartRubbishListResp, error) {
	client := rubbish.NewRubbishClient(m.cli.Conn())
	return client.GetCartRubbishList(ctx, in)
}

func (m *defaultRubbish) AddRubbishType(ctx context.Context, in *AddRubbishTypeReq) (*AddRubbishTypeResp, error) {
	client := rubbish.NewRubbishClient(m.cli.Conn())
	return client.AddRubbishType(ctx, in)
}

func (m *defaultRubbish) UpdateCartRubbish(ctx context.Context, in *UpdateCartRubbishReq) (*UpdateCartRubbishResp, error) {
	client := rubbish.NewRubbishClient(m.cli.Conn())
	return client.UpdateCartRubbish(ctx, in)
}