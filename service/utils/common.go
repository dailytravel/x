package utils

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContextKey string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Bytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

//Number ...
func Number(n int) string {
	var letters = []rune("0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Base64(n int, padded bool) (string, error) {
	bytes, err := Bytes(n)
	if err != nil {
		return "", err
	}
	result := base64.StdEncoding.EncodeToString(bytes)
	result = strings.Replace(result, "\n", "", -1)
	if !padded {
		result = strings.Replace(result, "=", "", -1)
	}
	return result, nil
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func Query(args map[string]interface{}) interface{} {
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
func Options(args map[string]interface{}) *options.FindOptions {
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