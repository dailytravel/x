package migrations

import "github.com/dailytravel/x/sales/pkg/database"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Balance{Database: database.Database},
		&Benefit{Database: database.Database},
		&Company{Database: database.Database},
		&Contact{Database: database.Database},
		&Coupon{Database: database.Database},
		&Product{Database: database.Database},
		&Package{Database: database.Database},
		&Membership{Database: database.Database},
		&Order{Database: database.Database},
		&Variant{Database: database.Database},
		&Promotion{Database: database.Database},
		&Quote{Database: database.Database},
		&Reward{Database: database.Database},
		&Tier{Database: database.Database},
		&Wishlist{Database: database.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
