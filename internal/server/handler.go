package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Eiphoria/GoReversi/internal/service"
)

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := w.Write([]byte("OK"))
	if err != nil {
		s.logger.Error("failed to write response", slog.Any("err", err))
		return
	}
}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Error("json newdecoder decode ", slog.Any("err", err))
		return
	}

	defer r.Body.Close()

	err := s.service.CreateUser(r.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidData) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.logger.Info("created user succesfull")
	w.WriteHeader(http.StatusCreated)

}
