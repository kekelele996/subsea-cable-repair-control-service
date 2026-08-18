package api

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"net/http"
)

func writeDecision(w http.ResponseWriter, d domain.CompletionDecision) {
	if !d.Allowed {
		http.Error(w, d.Reason, http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
