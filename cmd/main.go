package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/Eiphoria/GoReversi/internal/config"
	"github.com/Eiphoria/GoReversi/internal/repository"
	migrations "github.com/Eiphoria/GoReversi/internal/repository/miggrations"
	"github.com/Eiphoria/GoReversi/internal/server"
	"github.com/Eiphoria/GoReversi/internal/service"
	"github.com/Eiphoria/GoReversi/pkg/logger"
	"github.com/Eiphoria/GoReversi/pkg/migrator"
)

func main() {
	cfg := config.New()

	err := migrator.Migratos(migrations.Migrations, cfg.DBConf.ConnectionURL)
	if err != nil {
		panic(err)
	}

	repo, err := repository.New(cfg.DBConf)
	if err != nil {
		panic(err)
	}

	svc := service.New(repo, cfg.HashSalt)

	s := server.New(svc, logger.New())
	if err := s.Run(); err != nil {
		panic(err)
	}
}
