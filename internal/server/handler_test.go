package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Eiphoria/GoReversi/internal/config"
	"github.com/Eiphoria/GoReversi/internal/repository"
	migrations "github.com/Eiphoria/GoReversi/internal/repository/miggrations"
	"github.com/Eiphoria/GoReversi/internal/service"
	"github.com/Eiphoria/GoReversi/pkg/assert"
	"github.com/Eiphoria/GoReversi/pkg/logger"
	"github.com/Eiphoria/GoReversi/pkg/migrator"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func TestHealth(t *testing.T) {
	s := New(nil, nil)
	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
		w := httptest.NewRecorder()
		s.serv.Handler.ServeHTTP(w, req)
		if got := w.Result().StatusCode; got != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, got)
		}

		b, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("unxepected error")
			return
		}

		if string(b) != "OK" {
			t.Errorf("expected OK, got %s", string(b))
		}

	})
	t.Run("not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/health", nil)
		w := httptest.NewRecorder()
		s.serv.Handler.ServeHTTP(w, req)
		assert.Equal(t, http.StatusMethodNotAllowed, w.Result().StatusCode)

	})

}

func TestRegister(t *testing.T) {
	cfg := config.Config{
		DBConf: config.DBConfig{
			ConnectionURL: "postgres://postgres:postgres@127.0.0.1:8624/postgres?sslmode=disable",
		},
		HashSalt: "d74f9086f3c7482b5eb22c26c9793396",
	}

	err := migrator.Migratos(migrations.Migrations, cfg.DBConf.ConnectionURL)
	if err != nil {
		t.Fatal("migratos error: %w", err)
	}

	repo, err := repository.New(cfg.DBConf)
	if err != nil {
		t.Error("repository new error: %w", err)
	}

	svc := service.New(repo, cfg.HashSalt)
	s := New(svc, logger.New())
	t.Run("not allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/reg", nil)
		w := httptest.NewRecorder()
		s.serv.Handler.ServeHTTP(w, req)
		assert.Equal(t, http.StatusMethodNotAllowed, w.Result().StatusCode)

	})
	t.Run("bad request", func(t *testing.T) {
		reqBody := createUserRequest{Username: "te", Password: "testpasswordsssssssss"}
		reqJSON, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/reg", bytes.NewReader(reqJSON))
		w := httptest.NewRecorder()
		s.serv.Handler.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("created", func(t *testing.T) {
		reqBody := createUserRequest{Username: "testoser", Password: "tephs1!sword"}
		reqJSON, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/reg", bytes.NewReader(reqJSON))
		w := httptest.NewRecorder()
		s.serv.Handler.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
	})
}
