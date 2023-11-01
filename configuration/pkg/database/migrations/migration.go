package migrations

import "github.com/dailytravel/x/configuration/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Currency{Database: database.Database},
		&Locale{Database: database.Database},
		&Global{Database: database.Database},
		&Place{Database: database.Database},
		&Timezone{Database: database.Database},
		&Webhook{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
