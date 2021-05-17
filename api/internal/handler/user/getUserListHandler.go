package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"rubbish-cart/api/internal/logic/user"
	"rubbish-cart/api/internal/svc"
)

func GetUserListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetUserListLogic(r.Context(), ctx)
		resp, err := l.GetUserList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
