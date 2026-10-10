package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/xMurungi/elimu-core-api/internal/httpx"
)

type Pinger interface {
	Ping(ctx context.Context) error
}

// Healthz: "is the process alive?" No dependencies checked.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// Readyz: "can this instance serve traffic?" Checks the database.
func Readyz(db Pinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := db.Ping(ctx); err != nil {
			slog.ErrorContext(ctx, "readyz: db ping failed", "err", err)
			http.Error(w, "not ready", http.StatusServiceUnavailable)
			return
		}
		_, _ = w.Write([]byte("ok"))
	}
}

type versionBody struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(
		w,
		http.StatusOK,
		versionBody{Service: "core-api", Version: "0.1.0"})
}
