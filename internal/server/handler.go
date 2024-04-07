package server

import (
	"encoding/json"
	"errors"
	"fmt"
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
		panic(err)
	}
}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}
	defer r.Body.Close()
	/*валидироват юсернаме\пароль: 3 символа минимум,
	при ошибке валидации вернут ErrInvalidData
	в хандлер проверить на ошибку ErrInvalidData если она вернуть статус код 400 в ином случае 500
	если ошибок нету вернуть 201 ок
	*/
	err := s.service.CreateUser(r.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidData) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("bebrow: ", err)
		return
	}
	fmt.Println(req)
	w.WriteHeader(http.StatusCreated)

}
