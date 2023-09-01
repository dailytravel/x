package migrations

import "github.com/dailytravel/x/sales/db"

type Migrator interface {
	Migrate() error
}

func AutoMigrate() error {
	migrators := []Migrator{
		&Balance{Database: db.Database},
		&Benefit{Database: db.Database},
		&Company{Database: db.Database},
		&Contact{Database: db.Database},
		&Coupon{Database: db.Database},
		&Membership{Database: db.Database},
		&Order{Database: db.Database},
		&Point{Database: db.Database},
		&Variant{Database: db.Database},
		&Promotion{Database: db.Database},
		&Quote{Database: db.Database},
		&Reward{Database: db.Database},
		&Tier{Database: db.Database},
		&Wishlist{Database: db.Database},
	}

	for _, migration := range migrators {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}

	return nil
}
