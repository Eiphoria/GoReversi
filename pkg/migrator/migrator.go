package migrator

import (
	"embed"
	"errors"
	"fmt"

	"github.com/Eiphoria/GoReversi/internal/repository/miggrations"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func Migratos(migrations embed.FS, dbConn string) error {

	d, err := iofs.New(miggrations.Migrations, ".")
	if err != nil {
		return fmt.Errorf("iofs new error: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dbConn) //New("file:///C:/Games/gulagovna/GoReversi/internal/repository/migrations", "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		return fmt.Errorf("migrate withsrcinst error: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("m up error: %w", err)
	}

	return nil
}
