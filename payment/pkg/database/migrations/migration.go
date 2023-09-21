package migrations

import "github.com/dailytravel/x/payment/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Card{Database: database.Database},
		&Transaction{Database: database.Database},
		&Wallet{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
