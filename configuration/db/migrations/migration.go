package migrations

import "github.com/dailytravel/x/configuration/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Country{Database: db.Database},
		&Currency{Database: db.Database},
		&Locale{Database: db.Database},
		&Option{Database: db.Database},
		&Template{Database: db.Database},
		&Timezone{Database: db.Database},
		&Webhook{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
