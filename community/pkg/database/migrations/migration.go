package migrations

import "github.com/dailytravel/x/community/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Comment{Database: database.Database},
		&Conversation{Database: database.Database},
		&Share{Database: database.Database},
		&Message{Database: database.Database},
		&Notification{Database: database.Database},
		&Reaction{Database: database.Database},
		&Recipient{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
