package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dailytravel/x/account/pkg/auth"
	"github.com/typesense/typesense-go/typesense/api"
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
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	encoded := base64.RawURLEncoding.EncodeToString(bytes)
	if !padded {
		encoded = strings.TrimRight(encoded, "=")
	}
	return encoded, nil
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

func Params(args map[string]interface{}) *api.SearchCollectionParams {
	searchParameters := &api.SearchCollectionParams{
		Q:        getStringArg(args, "q"),
		QueryBy:  getStringArg(args, "query_by"),
		FilterBy: getStringPointerArg(args, "filter_by"),
		SortBy:   getStringPointerArg(args, "sort_by"),
		Page:     getIntPointerArg(args, "page"),
		PerPage:  getIntPointerArg(args, "per_page"),
	}

	return searchParameters
}

func getStringArg(args map[string]interface{}, key string) string {
	if val, ok := args[key].(string); ok && val != "" {
		return val
	}
	return ""
}

func getStringPointerArg(args map[string]interface{}, key string) *string {
	if val, ok := args[key].(string); ok && val != "" {
		return &val
	}
	return nil
}

func getIntPointerArg(args map[string]interface{}, key string) *int {
	if val, ok := args[key].(json.Number); ok {
		if intValue, err := val.Int64(); err == nil {
			intVal := int(intValue)
			return &intVal
		}
	}
	return nil
}

func UID(ctx context.Context) (*primitive.ObjectID, error) {
	claims := auth.Auth(ctx)
	if claims == nil {
		return nil, fmt.Errorf("not authenticated")
	}

	uid, err := primitive.ObjectIDFromHex(claims["sub"].(string))
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}

	return &uid, nil
}

func Contains(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func DecodeFromBase64(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
