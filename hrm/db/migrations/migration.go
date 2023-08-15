package migrations

import "github.com/dailytravel/x/hrm/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Attendance{Database: db.Database},
		&Leave{Database: db.Database},
		&Payroll{Database: db.Database},
		&Resume{Database: db.Database},
		&Salary{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
