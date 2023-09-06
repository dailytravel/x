package migrations

import "github.com/dailytravel/x/finance/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Expense{Database: database.Database},
		&Invoice{Database: database.Database},
		&Payment{Database: database.Database},
		&Transaction{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
