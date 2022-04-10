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
		"postgres://uhjxadtqnmvedz:827b87d9c520ea5e6c07d04522a1efa923be6740fe2332174bc7511532e52e0a@ec2-63-35-156-160.eu-west-1.compute.amazonaws.com:5432/d5qq14sdoip2f")

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
