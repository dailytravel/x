package migrations

import "github.com/dailytravel/x/cms/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&File{Database: database.Database},
		&Image{Database: database.Database},
		&Post{Database: database.Database},
		&Term{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
