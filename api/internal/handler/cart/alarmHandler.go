package handler

import (
	"net/http"

	"rubbish-cart/api/internal/logic/cart"
	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AlarmHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AlarmReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAlarmLogic(r.Context(), ctx)
		resp, err := l.Alarm(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
