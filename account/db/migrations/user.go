package migrations

import (
	"context"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Database *mongo.Database
	Model    *model.User
}

func (m *User) Migrate() error {
	emails := []map[string]interface{}{
		{
			"name":  "Admin",
			"email": "admin@uid.one",
			"roles": []string{"admin"},
		},
		{
			"name":  "SM",
			"email": "info@dailytravel.vn",
			"roles": []string{"editor"},
		},
		{
			"name":  "Sales 1",
			"email": "sales1@dailytravel.vn",
			"roles": []string{"editor"},
		},
		{
			"name":  "Sales 2",
			"email": "sales2@dailytravel.vn",
			"roles": []string{"editor"},
		},
		{
			"name":  "Sales 3",
			"email": "sales3@dailytravel.vn",
			"roles": []string{"editor"},
		},
	}

	col := m.Database.Collection(m.Model.Collection())
	indexes, err := col.Indexes().List(context.Background())
	if err != nil {
		return err
	}

	indexNames := make(map[string]bool)
	for indexes.Next(context.Background()) {
		var index bson.M
		if err := indexes.Decode(&index); err != nil {
			return err
		}

		indexNames[index["name"].(string)] = true
	}

	for _, index := range m.Model.Index() {
		keys := index.Keys
		if keys != nil {
			indexName := ""
			for _, key := range keys.(bson.D) {
				indexName = key.Key
				break
			}

			if !indexNames[indexName] {
				if _, err := col.Indexes().CreateOne(context.Background(), index); err != nil {
					return err
				}
			}
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("abc@123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	for _, email := range emails {
		// Check if the user exists
		filter := bson.M{"email": email["email"].(string)}
		if err := col.FindOne(context.Background(), filter).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				item := &model.User{
					Name:  email["name"].(string),
					Email: email["email"].(string),
				}

				rolesSlice := email["roles"].([]string)
				roles := make([]string, len(rolesSlice))
				for i, role := range rolesSlice {
					r := role // Create a copy of the string
					roles[i] = r
				}
				item.Roles = roles
				item.Password = string(hashedPassword)

				if _, err := col.InsertOne(context.Background(), item); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
