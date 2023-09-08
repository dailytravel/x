package migrations

import "github.com/dailytravel/x/base/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Goal{Database: database.Database},
		&Board{Database: database.Database},
		&List{Database: database.Database},
		&Task{Database: database.Database},
		&Phase{Database: database.Database},
		&Portfolio{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
