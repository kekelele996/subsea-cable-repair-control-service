package api

import (
	"encoding/json"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
	"net/http"
)

type Server struct{ s *service.System }

func NewServer(s *service.System) *Server { return &Server{s: s} }
func (x *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/plans", x.create)
	m.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	return m
}
func (x *Server) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var input struct {
		ID    string   `json:"id"`
		Cable string   `json:"cable"`
		Spans []string `json:"spans"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p := x.s.CreatePlan(input.ID, input.Cable, input.Spans)
	_ = json.NewEncoder(w).Encode(p)
}
