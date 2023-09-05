package migrations

import "github.com/dailytravel/x/community/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Comment{Database: db.Database},
		&Conversation{Database: db.Database},
		&Share{Database: db.Database},
		&Message{Database: db.Database},
		&Notification{Database: db.Database},
		&Reaction{Database: db.Database},
		&Recipient{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
