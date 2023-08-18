package migrations

import "github.com/dailytravel/x/base/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Goal{Database: db.Database},
		&Board{Database: db.Database},
		&List{Database: db.Database},
		&Task{Database: db.Database},
		&Portfolio{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
