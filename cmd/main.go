package main

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/Eiphoria/GoReversi/internal/config"
	"github.com/Eiphoria/GoReversi/internal/repository"
	"github.com/Eiphoria/GoReversi/internal/repository/migrations"
	"github.com/Eiphoria/GoReversi/internal/server"
	"github.com/Eiphoria/GoReversi/internal/service"
)

func main() {
	cfg := config.New()

	d, err := iofs.New(migrations.Migrations, ".")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, cfg.DBConf.ConnectionURL)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	repo, err := repository.New(cfg.DBConf)
	if err != nil {
		panic(err)
	}

	svc := service.New(repo)
	s := server.New(svc)
	if err := s.Run(); err != nil {
		panic(err)
	}
}
