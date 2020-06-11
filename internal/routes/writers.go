package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/requestid"
)

func Write(_ context.Context, w http.ResponseWriter, code int, out interface{}) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(out)
}

func WriteError(ctx context.Context, w http.ResponseWriter, out error) {
	w.Header().Add("Content-Type", "application/json")

	id, _ := requestid.GetRequestId(ctx)

	code := errors.Code(out)
	if code == 0 {
		code = 500
	}

	w.WriteHeader(code)

	if out == nil {
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{
		"id":      id,
		"message": out.Error(),
	})
}
