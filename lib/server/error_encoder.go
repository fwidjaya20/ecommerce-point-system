package server

import (
	"context"
	"encoding/json"
	libError "github.com/fwidjaya20/ecommerce-point-system/lib/error"
	"net/http"
)

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	code := http.StatusInternalServerError
	message := "Something Went Wrong"
	messageCode := "server_side_error"

	if sc, ok := err.(*libError.Error); ok {
		code = sc.StatusCode
		message = sc.Message
		messageCode = sc.MessageCode
	}

	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"trace":  err.Error(),
		"error": message,
		"code": messageCode,
	})
}