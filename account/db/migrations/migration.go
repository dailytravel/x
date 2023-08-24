package migrations

import "github.com/dailytravel/x/account/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Api{Database: db.Database},
		&Client{Database: db.Database},
		&Connection{Database: db.Database},
		&Integration{Database: db.Database},
		&Key{Database: db.Database},
		&Permission{Database: db.Database},
		&Role{Database: db.Database},
		&Token{Database: db.Database},
		&User{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
