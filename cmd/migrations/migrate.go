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
		"postgres://aaxagnxiuvdgfe:0ca29444e6e2a9dd1565bc6384921da5d5f4511587fe4498f9219e4a87ff3400@ec2-176-34-211-0.eu-west-1.compute.amazonaws.com:5432/df6u184vr7eg3s")

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
