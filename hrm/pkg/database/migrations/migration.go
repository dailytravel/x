package migrations

import "github.com/dailytravel/x/hrm/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Application{Database: database.Database},
		&Attendance{Database: database.Database},
		&Job{Database: database.Database},
		&Leave{Database: database.Database},
		&Organization{Database: database.Database},
		&Resume{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
