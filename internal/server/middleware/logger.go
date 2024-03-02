package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type responseWriter struct {
	body *bytes.Buffer
	http.ResponseWriter
}

func (r *responseWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// NewLogging httpリクエストとレスポンスのログを出力する
func NewLogging(zapLogger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req := ctx.Request()
			res := ctx.Response()
			traceID := getTraceID(ctx)
			reqHeader := suppressHTTPHeader(req.Header)

			bufBody, _ := io.ReadAll(req.Body)
			reqBody := map[string]any{}
			_ = json.Unmarshal(bufBody, &reqBody)

			req.Body = io.NopCloser(bytes.NewBuffer(bufBody))

			zapLogger.Info("request",
				zap.String("trace_id", traceID),
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.String("scheme", req.URL.Scheme),
				zap.String("host", req.Host),
				zap.String("endpoint", req.RequestURI),
				zap.String("query", req.URL.RawQuery),
				zap.String("remote-address", req.RemoteAddr),
				zap.String("user-agent", req.UserAgent()),
				zap.Any("header", reqHeader),
				zap.Any("body", reqBody),
			)

			res.Writer = &responseWriter{
				body:           bytes.NewBuffer([]byte{}),
				ResponseWriter: res.Writer,
			}

			if err := next(ctx); err != nil {
				ctx.Error(err)
			}

			resBody := res.Writer.(*responseWriter).body.String()

			zapLogger.Info("response",
				zap.String("trace_id", traceID),
				zap.String("method", req.Method),
				zap.Int("status", res.Status),
				zap.Any("header", res.Header()),
				zap.Any("body", resBody),
			)

			// エラーレスポンスの場合はsentryに送信する
			switch res.Status {
			case http.StatusInternalServerError:
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("trace_id", traceID) // trace_idをSentryのタグとして設定
					sentry.CaptureMessage("Internal Server Error: " + resBody)
				})
			case http.StatusBadRequest:
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("trace_id", traceID) // trace_idをSentryのタグとして設定
					sentry.CaptureMessage("Bad Request: " + resBody)
				})
			case http.StatusBadGateway:
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("trace_id", traceID) // trace_idをSentryのタグとして設定
					sentry.CaptureMessage("Bad Gateway: " + resBody)
				})
			}

			return nil
		}
	}
}

// Authorizationヘッダの値をマスクする
func suppressHTTPHeader(header http.Header) http.Header {
	result := make(http.Header, len(header))
	for k, v := range header {
		if k == "Authorization" {
			for _, auth := range header.Values(k) {
				masked := strings.Replace(auth, auth[strings.Index(auth, " ")+1:], "***", -1)
				result.Add(k, masked)
			}
		} else {
			result[k] = v
		}
	}
	return result
}
