package migrations

import "github.com/dailytravel/x/cms/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Term{Database: db.Database},
		&Content{Database: db.Database},
		&Country{Database: db.Database},
		&Currency{Database: db.Database},
		&Locale{Database: db.Database},
		&Term{Database: db.Database},
		&Timezone{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
