package handler

import (
	"net/http"

	"github.com/elton/bookstore/api/internal/logic"
	"github.com/elton/bookstore/api/internal/svc"
	"github.com/elton/bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func addHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRequst
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), ctx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
