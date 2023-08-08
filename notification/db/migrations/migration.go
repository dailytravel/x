package migrations

import "github.com/dailytravel/x/notification/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Conversation{Database: db.Database},
		&Notification{Database: db.Database},
		&Message{Database: db.Database},
		&Recipient{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
