package migrations

import "github.com/dailytravel/x/account/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Api{Database: database.Database},
		&Client{Database: database.Database},
		&Connection{Database: database.Database},
		&Integration{Database: database.Database},
		&Key{Database: database.Database},
		&Permission{Database: database.Database},
		&Role{Database: database.Database},
		&User{Database: database.Database},
		&Credential{Database: database.Database},
		&Token{Database: database.Database},
		&Workspace{Database: database.Database},
		&Member{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
