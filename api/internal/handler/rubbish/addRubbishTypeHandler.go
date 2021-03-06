package handler

import (
	"net/http"

	"rubbish-cart/api/internal/logic/rubbish"
	"rubbish-cart/api/internal/svc"
	"rubbish-cart/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AddRubbishTypeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRubbishTypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddRubbishTypeLogic(r.Context(), ctx)
		resp, err := l.AddRubbishType(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
