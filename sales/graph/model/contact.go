package model

import (
	"context"
	"time"

	"github.com/dailytravel/x/sales/db"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contact struct {
	Model        `bson:",inline"`
	User         primitive.ObjectID  `bson:"user" json:"user"`
	Company      *primitive.ObjectID `bson:"company,omitempty" json:"company,omitempty"`
	Reference    string              `json:"reference" bson:"reference"`
	FirstName    string              `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName     string              `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Birthday     primitive.DateTime  `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Gender       Gender              `json:"gender,omitempty" bson:"gender,omitempty"`
	JobTitle     string              `json:"job_title,omitempty" bson:"job_title,omitempty"`
	Email        Email               `json:"email,omitempty" bson:"email,omitempty"`
	Phone        Phone               `json:"phone,omitempty" bson:"phone,omitempty"`
	Picture      string              `json:"picture,omitempty" bson:"picture,omitempty"`
	Street       string              `json:"street,omitempty" bson:"street,omitempty"`
	City         string              `json:"city,omitempty" bson:"city,omitempty"`
	Zip          string              `json:"zip,omitempty" bson:"zip,omitempty"`
	State        string              `json:"state,omitempty" bson:"state,omitempty"`
	Country      string              `json:"country,omitempty" bson:"country,omitempty"`
	Website      string              `json:"website,omitempty" bson:"website,omitempty"`
	Photos       []*string           `json:"photos,omitempty" bson:"photos,omitempty"`
	Source       string              `json:"source,omitempty" bson:"source,omitempty"`
	Revenue      float64             `json:"revenue,omitempty" bson:"revenue,omitempty"`
	Timezone     string              `json:"timezone,omitempty" bson:"timezone,omitempty"`
	Language     string              `json:"language,omitempty" bson:"language,omitempty"`
	Subscribed   bool                `json:"subscribed,omitempty" bson:"subscribed,omitempty"`
	Notes        string              `json:"notes,omitempty" bson:"notes,omitempty"`
	Stage        string              `json:"stage,omitempty" bson:"stage,omitempty"`
	Labels       []string            `json:"labels,omitempty" bson:"labels,omitempty"`
	Status       string              `json:"status" bson:"status"`
	Reviewable   bool                `json:"reviewable,omitempty" bson:"reviewable,omitempty"`
	LastActivity primitive.Timestamp `json:"last_activity,omitempty" bson:"last_activity,omitempty"`
}

func (Contact) IsEntity() {}

func (i *Contact) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.UpdatedBy == nil {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Contact
	return bson.Marshal((*t)(i))
}

func (i *Contact) Collection() string {
	return "contacts"
}

func (i *Contact) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "user", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "company", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "first_name", Value: "text"}, {Key: "last_name", Value: "text"}}, Options: options.Index().SetWeights(bson.M{"first_name": 2, "last_name": 1})},
		{Keys: bson.D{{Key: "source", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "phone", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "birthday", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "last_activity", Value: 1}}, Options: options.Index()},
	}
}

func (i *Contact) Schema() interface{} {
	return &api.CollectionSchema{
		Name: i.Collection(),
		Fields: []api.Field{
			{Name: "user", Type: "string"},
			{Name: "company", Type: "string", Optional: pointer.True()},
			{Name: "first_name", Type: "string", Optional: pointer.True()},
			{Name: "last_name", Type: "string", Optional: pointer.True()},
			{Name: "job_title", Type: "string", Optional: pointer.True()},
			{Name: "gender", Type: "string", Optional: pointer.True()},
			{Name: "photos", Type: "string[]", Optional: pointer.True()},
			{Name: "country", Type: "string", Optional: pointer.True(), Facet: pointer.True()},
			{Name: "timezone", Type: "string", Optional: pointer.True()},
			{Name: "language", Type: "string", Optional: pointer.True()},
			{Name: "source", Type: "string", Optional: pointer.True()},
			{Name: "revenue", Type: "float"},
			{Name: "notes", Type: "string", Optional: pointer.True()},
			{Name: "status", Type: "string", Facet: pointer.True()},
			{Name: "created_by", Type: "string", Optional: pointer.True()},
			{Name: "updated_by", Type: "string", Optional: pointer.True()},
			{Name: "created_at", Type: "int32"},
			{Name: "updated_at", Type: "int32"},
			{Name: "last_activity", Type: "int32"},
			{Name: "birthday", Type: "int32", Facet: pointer.True()},
			{Name: "email", Type: "object", Optional: pointer.True()},
			{Name: "phone", Type: "object", Optional: pointer.True()},
			{Name: "metadata", Type: "object", Optional: pointer.True()},
			{Name: "followers", Type: "string[]", Optional: pointer.True()},
			{Name: "labels", Type: "string[]", Optional: pointer.True()},
		},
		DefaultSortingField: pointer.String("created_at"),
		TokenSeparators:     &[]string{"(", ")", "-"},
		EnableNestedFields:  pointer.True(),
	}
}

func (i *Contact) Document() map[string]interface{} {
	// followers := i.Followers()

	document := map[string]interface{}{
		"id":            i.ID,
		"user":          i.User,
		"company":       i.Company,
		"first_name":    i.FirstName,
		"last_name":     i.LastName,
		"job_title":     i.JobTitle,
		"gender":        i.Gender.String(),
		"source":        i.Source,
		"country":       i.Country,
		"city":          i.City,
		"state":         i.State,
		"street":        i.Street,
		"zip":           i.Zip,
		"timezone":      i.Timezone,
		"language":      i.Language,
		"status":        i.Status,
		"photos":        i.Photos,
		"email":         i.Email,
		"phone":         i.Phone,
		"metadata":      i.Metadata,
		"revenue":       i.Revenue,
		"created_at":    i.CreatedAt.T,
		"updated_at":    i.UpdatedAt.T,
		"last_activity": i.LastActivity.T,
		"birthday":      i.Birthday.Time().Unix(),
		// "followers":     followers,
		"labels": i.Labels,
	}

	return document
}

func (i *Contact) Insert(collection typesense.CollectionInterface) error {
	document := i.Document()

	// Retrieve Typesense collection schema and create it if it doesn't exist
	if _, err := collection.Retrieve(); err != nil {
		// Create collection
		if _, err := db.Client.Collections().Create(i.Schema().(*api.CollectionSchema)); err != nil {
			return err
		}
	}

	if _, err := collection.Documents().Create(document); err != nil {
		return err
	}

	return nil
}

func (i *Contact) Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error {
	documentID := documentKey["_id"].(primitive.ObjectID).Hex()

	// Create a map to hold the updated fields
	updatePayload := make(map[string]interface{})

	// Check if 'deleted_at' field is in updatedFields and its value is of type primitive.Timestamp
	_, deletedAtExist := updatedFields["deleted_at"].(primitive.Timestamp)
	if deletedAtExist {
		if err := i.Delete(collection, documentKey); err != nil {
			return err
		}
		return nil
	}

	// Loop through updatedFields
	for field, value := range updatedFields {
		switch field {
		case "created_at", "updated_at", "last_activity":
			if timestamp, ok := value.(primitive.Timestamp); ok {
				updatePayload[field] = timestamp.T
			}
		default:
			updatePayload[field] = value
		}
	}

	// Loop through removedFields
	for _, field := range removedFields {
		updatePayload[field.(string)] = nil
	}

	// Update the document with the updatePayload
	if _, err := collection.Document(documentID).Update(updatePayload); err != nil {
		// If the update fails, attempt to retrieve the item from the dataModel
		var item *Contact
		if err := db.Database.Collection(i.Collection()).FindOne(context.Background(), documentKey).Decode(&item); err != nil {
			return err
		}

		// Insert the item if it doesn't exist in the collection
		if err := i.Insert(collection); err != nil {
			return err
		}
	}

	return nil
}

func (i *Contact) Delete(collection typesense.CollectionInterface, documentKey primitive.M) error {
	id := documentKey["_id"].(primitive.ObjectID).Hex()

	// Delete document from Typesense
	if _, err := collection.Document(id).Delete(); err != nil {
		return err
	}

	return nil
}

func (i *Contact) Setuser(id *string) {
	if id != nil {
		i.User, _ = primitive.ObjectIDFromHex(*id)
	}
}

func (i *Contact) SetReference(r *string) {
	if r != nil {
		i.Reference = *r
	}
}

func (i *Contact) SetFirstName(f *string) {
	if f != nil {
		i.FirstName = *f
	}
}

func (i *Contact) SetLastName(l *string) {
	if l != nil {
		i.LastName = *l
	}
}

func (i *Contact) SetJobTitle(j *string) {
	if j != nil {
		i.JobTitle = *j
	}
}

func (c *Contact) SetBirthday(b *string) {
	if b != nil {
		birthdayTime, err := time.Parse(time.RFC3339, *b)
		if err == nil {
			c.Birthday = primitive.NewDateTimeFromTime(birthdayTime)
		}
	}
}

func (c *Contact) SetGender(g *string) {
	c.Gender = Gender(*g)
}

func (c *Contact) SetEmail(e Email) {
	c.Email = e
}

func (c *Contact) SetPhone(p Phone) {
	c.Phone = p
}

func (c *Contact) SetStreet(s *string) {
	if s != nil {
		c.Street = *s
	}
}

func (c *Contact) SetCity(ct *string) {
	if ct != nil {
		c.City = *ct
	}
}

func (c *Contact) SetZip(z *string) {
	if z != nil {
		c.Zip = *z
	}
}

func (c *Contact) SetState(st *string) {
	if st != nil {
		c.State = *st
	}
}

func (c *Contact) SetCountry(co *string) {
	if co != nil {
		c.Country = *co
	}
}

func (c *Contact) SetWebsite(w *string) {
	if w != nil {
		c.Website = *w
	}
}

func (c *Contact) SetSource(s *string) {
	if s != nil {
		c.Source = *s
	}
}

func (c *Contact) SetPicture(picture *string) {
	if picture != nil {
		c.Picture = *picture
	}
}

func (c *Contact) SetRevenue(r float64) {
	c.Revenue = r
}

func (c *Contact) SetTimezone(tz *string) {
	if tz != nil {
		c.Timezone = *tz
	}
}

func (c *Contact) SetLanguage(lang *string) {
	if lang != nil {
		c.Language = *lang
	}
}

func (c *Contact) SetSubscribed(sub bool) {
	c.Subscribed = sub
}

func (c *Contact) SetNotes(notes *string) {
	if notes != nil {
		c.Notes = *notes
	}
}

func (c *Contact) SetStage(stage *string) {
	if stage != nil {
		c.Stage = *stage
	}
}

func (c *Contact) SetStatus(status *string) {
	if status != nil {
		c.Status = *status
	}
}

func (c *Contact) SetLastActivity(t *primitive.Timestamp) {
	if t != nil {
		c.LastActivity = *t
	}
}

func (c *Contact) SetMetadata(m map[string]interface{}) {
	if m != nil {
		if c.Metadata == nil {
			c.Metadata = make(map[string]interface{})
		}

		for k, v := range m {
			c.Metadata[k] = v
		}
	}
}

func (c *Contact) SetLabels(labels []*string) {
	c.Labels = make([]string, len(labels))
	for i, v := range labels {
		if v != nil {
			c.Labels[i] = *v
		}
	}
}
