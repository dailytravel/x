package model

import (
	"encoding/json"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ClientCollection     = "clients"
	InvitationCollection = "invitations"
	PermissionCollection = "permissions"
	RoleCollection       = "roles"
	KeyCollection        = "keys"
	UserCollection       = "users"
)

type Model struct {
	ID          primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Metadata    primitive.M         `json:"metadata,omitempty" bson:"metadata,omitempty"`
	CreatedAt   primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	PublishedAt primitive.Timestamp `json:"published_at,omitempty" bson:"published_at,omitempty"`
	DeletedAt   primitive.Timestamp `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	CreatedBy   primitive.ObjectID  `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedBy   primitive.ObjectID  `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	DeletedBy   primitive.ObjectID  `json:"deleted_by,omitempty" bson:"deleted_by,omitempty"`
}

func (i *Model) Query(args map[string]interface{}) interface{} {
	query := bson.M{"deleted_at": bson.M{"$exists": false}}
	q := ""
	if val, ok := args["q"].(string); ok && val != "" {
		q = val
	}

	if q != "" {
		orQuery := []bson.M{{"$text": bson.M{"$search": q}}}
		query["$or"] = orQuery
	}

	if queryBy, ok := args["query"].(string); ok && queryBy != "" {
		queryFields := strings.Split(queryBy, ",")
		for _, field := range queryFields {
			if field == "query" {
				continue
			}
			val, ok := args[field].(string)
			if ok && val != "" {
				query[field] = bson.M{"$regex": primitive.Regex{Pattern: val, Options: "i"}}
			}
		}
	}

	if filter, ok := args["filter"].(map[string]interface{}); ok && filter != nil {
		for k, v := range filter {
			if k == "_id" || k == "object._id" {
				_id, _ := primitive.ObjectIDFromHex(v.(string))
				query[k] = _id
			}
			query[k] = v
		}
	}

	return query
}

// Options returns a MongoDB options object with skip, limit, and sort applied.
func (i *Model) Options(args map[string]interface{}) *options.FindOptions {
	options := options.Find()

	page := int64(1)
	if p, ok := args["page"].(json.Number); ok {
		page, _ = p.Int64()
		if page < 1 {
			page = 1 // handle negative or zero page values
		}
	}

	limit := int64(0) // Set limit to 0 to indicate no limit
	if l, ok := args["limit"].(json.Number); ok {
		limit, _ = l.Int64()
		if limit < 1 {
			limit = 1 // handle negative or zero limit values
		}
	}

	options.SetSkip((page - 1) * limit)
	options.SetLimit(limit)

	if sortBy, ok := args["sort"].(map[string]interface{}); ok {
		for k, v := range sortBy {
			// convert v to lowercase as "asc" or "desc" to 1 or -1
			if strings.ToLower(v.(string)) == "asc" {
				options.SetSort(bson.M{k: 1})
			} else {
				options.SetSort(bson.M{k: -1})
			}
		}
	}

	return options
}
