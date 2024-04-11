package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Eiphoria/GoReversi/internal/config"
	"github.com/Eiphoria/GoReversi/internal/repository"
	migrations "github.com/Eiphoria/GoReversi/internal/repository/miggrations"
	"github.com/Eiphoria/GoReversi/internal/service"
	"github.com/Eiphoria/GoReversi/pkg/assert"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func TestHealth(t *testing.T) {
	s := New(nil)
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
	}

	d, err := iofs.New(migrations.Migrations, ".")
	if err != nil {
		log.Fatal(err, "1")
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, cfg.DBConf.ConnectionURL) //New("file:///C:/Games/gulagovna/GoReversi/internal/repository/migrations", "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err, "2")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err, "3")
	}
	repo, err := repository.New(cfg.DBConf)
	if err != nil {
		log.Fatal(err, "4")
	}
	svc := service.New(repo)
	s := New(svc)
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
