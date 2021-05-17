package logic

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rubbish-cart/rpc/user/internal/config"

	"rubbish-cart/rpc/user/internal/svc"
	"rubbish-cart/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOpenIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOpenIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOpenIdLogic {
	return &GetOpenIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOpenIdLogic) GetOpenId(in *user.GetOpenIdReq) (*user.GetOpenIdResp, error) {
	Appid := config.Appid
	Appsecret := config.Appsecret
	res, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + Appid + "&secret=" + Appsecret + "&js_code=" + in.Code + "&grant_type=authorization_code")
	if err != nil {
		return nil, err
	}
	res1, err := ioutil.ReadAll(res.Body)

	var openid string
	var f interface{}
	err = json.Unmarshal(res1, &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		if k == "openid" {
			openid = v.(string)
		}
	}
	return &user.GetOpenIdResp{
		OpenId: openid,
	}, nil
}
