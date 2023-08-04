package migrations

import (
	"github.com/dailytravel/x/account/pkg/mongo"
)

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&User{DB: mongo.DB},
		&Key{DB: mongo.DB},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
