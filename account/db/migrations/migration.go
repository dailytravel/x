package migrations

import "github.com/dailytravel/x/account/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&User{Database: db.Database},
		&Role{Database: db.Database},
		&Permission{Database: db.Database},
		&Client{Database: db.Database},
		&Key{Database: db.Database},
		&Notification{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
