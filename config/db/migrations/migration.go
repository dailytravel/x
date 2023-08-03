package migrations

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
