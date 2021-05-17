package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"rubbish-cart/api/internal/logic/cart"
	"rubbish-cart/api/internal/svc"
)

func GetInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetInfoLogic(r.Context(), ctx)
		resp, err := l.GetInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
