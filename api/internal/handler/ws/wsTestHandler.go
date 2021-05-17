package handler

import (
	"net/http"

	"rubbish-cart/api/internal/svc"
)

func WsTestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "C:\\WorkSpace\\MyGo\\src\\github.com\\zeromicro\\zero-examples\\chat\\home.html")
	}
}
