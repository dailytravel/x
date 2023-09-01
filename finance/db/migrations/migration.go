package migrations

import "github.com/dailytravel/x/finance/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Expense{Database: db.Database},
		&Invoice{Database: db.Database},
		&Payment{Database: db.Database},
		&Transaction{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
