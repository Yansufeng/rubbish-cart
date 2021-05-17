package handler

import (
	"net/http"

	"rubbish-cart/api/internal/logic/rubbish"
	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateCartRubbishHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCartRubbishReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateCartRubbishLogic(r.Context(), ctx)
		resp, err := l.UpdateCartRubbish(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
