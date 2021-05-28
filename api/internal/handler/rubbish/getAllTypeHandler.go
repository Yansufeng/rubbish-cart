package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"rubbish-cart/api/internal/logic/rubbish"
	"rubbish-cart/api/internal/svc"
)

func GetAllTypeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetAllTypeLogic(r.Context(), ctx)
		resp, err := l.GetAllType()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
