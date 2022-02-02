package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

func Compose() error {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://ttuzpbnnzexpnh:3aae25a4574b6a3a47781a0afe4c306a0461d9e7010cad46d69a9ba3933ffea6@ec2-52-18-116-67.eu-west-1.compute.amazonaws.com:5432/df82hje0i8jjqa")

	if err != nil {
		return errors.New("не удалось создать инстанс миграций")
	}

	version, _, _ := m.Version()

	if version != 1 {
		if err := m.Up(); err != nil {
			return err
		}
	}
	return nil
}
