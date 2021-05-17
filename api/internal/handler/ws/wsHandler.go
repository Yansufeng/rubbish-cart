package handler

import (
	"net/http"

	"rubbish-cart/api/internal/logic/ws"
	"rubbish-cart/api/internal/svc"
)

func WsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hub := logic.NewHub()
		logic.ServeWs(hub, w, r)
	}
}
