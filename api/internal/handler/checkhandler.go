package handler

import (
	"net/http"

	"github.com/elton/bookstore/api/internal/logic"
	"github.com/elton/bookstore/api/internal/svc"
	"github.com/elton/bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func checkHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCheckLogic(r.Context(), ctx)
		resp, err := l.Check(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
