package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/models"
	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/server/tools"

	"github.com/lithammer/shortuuid"
)

func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requireId := shortuuid.New()
		ctx := context.WithValue(r.Context(), models.RequireIdKey, requireId)
		r = r.WithContext(ctx)

		startTime := time.Now()
		tools.AccessLogStart(r.URL.Path, r.RemoteAddr, r.Method, requireId, startTime)
		next.ServeHTTP(w, r)
		tools.AccessLogEnd(r.URL.Path, r.RemoteAddr, r.Method, requireId, startTime)
	})
}
