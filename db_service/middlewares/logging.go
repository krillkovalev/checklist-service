// To do доделать logging middleware

package middlewares

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"context"
)

func Logger(l zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		rec := httptest.NewRecorder()

		ctx := r.Context()

		path := r.URL.EscapedPath()

		reqData, _ := httputil.DumpRequest(r, false)

		logger := l.Log().Timestamp().Str("path", path).Bytes("request_data", reqData)

		defer func (begin time.Time) {
			status := ww.Status()

			tookMs := time.Since(begin).Milliseconds()
			logger.Int64("took", tookMs).Int("status_code", status).Msgf("[%d] %s http request for %s took %dms", status, r.Method, path, tookMs)
		}(time.Now())

		ctx = context.WithValue(ctx, "logger", logger)
		next.ServeHTTP(rec, r.WithContext(ctx))

		for k, v := range rec.Header() {
			ww.Header()[k] = v
		}
		ww.WriteHeader(rec.Code)
		rec.Body.WriteTo(ww)
		})
	}
}	