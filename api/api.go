package api

import (
	"cf-proposal/common/types"
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

const basePath types.Path = "/"
const createpath types.Path = "/create"
const deletepath types.Path = "/delete/{id}"
const statisticspath types.Path = "/{id}"

// TODO: Figure out how to either generalize middleware handlers; time permitting.
func RedirectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "shortUrl", chi.URLParam(r, "shortUrl"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
