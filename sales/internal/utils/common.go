package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dailytravel/x/sales/pkg/auth"
	"github.com/typesense/typesense-go/typesense/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContextKey string

const (
	cardNumberPrefix = "MCRD"
	cardNumberLength = 12
)

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

//String ...
func String(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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

func Date(dateStr *string, targetDate *primitive.DateTime) {
	if dateStr != nil {
		dateTime, err := time.Parse(time.RFC3339, *dateStr)
		if err == nil {
			date := primitive.NewDateTimeFromTime(dateTime)
			*targetDate = date
		}
	}
}

// GenerateMembershipCardNumber generates a unique membership card number.
func GenerateMembershipCardNumber() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random numeric portion of the card number
	randomNumber := rand.Intn(90000000) + 10000000

	// Combine the prefix and the random number to create the card number
	cardNumber := fmt.Sprintf("%s%d", cardNumberPrefix, randomNumber)

	return cardNumber
}

// indexOf finds the index of a string in a slice.
// It returns -1 if the string is not present in the slice.
func IndexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func Filter(input interface{}) interface{} {
	switch v := input.(type) {
	case string:
		// If string matches the length of ObjectID, try converting
		if len(v) == 24 {
			if objID, err := primitive.ObjectIDFromHex(v); err == nil {
				return objID
			}
		}
		return v
	case []interface{}:
		// Handle slice of interfaces, which could be slice of strings
		for i, item := range v {
			v[i] = Filter(item)
		}
		return v
	case map[string]interface{}:
		// Recursive call for nested maps
		for key, value := range v {
			v[key] = Filter(value)
		}
		return v
	default:
		// Return as it is if it's any other type
		return v
	}
}

// Sort converts a map of sort criteria into MongoDB-compatible sort options.
func Sort(sortCriteria map[string]interface{}) *options.FindOptions {
	sortMap := bson.D{}

	for field, order := range sortCriteria {
		orderStr, ok := order.(string)
		if !ok {
			continue
		}

		orderValue := 1 // default to ascending
		if orderStr == "desc" {
			orderValue = -1
		}
		sortMap = append(sortMap, bson.E{Key: field, Value: orderValue})
	}

	return options.Find().SetSort(sortMap)
}

// Project converts a map into MongoDB-compatible projection options.
func Project(projection map[string]interface{}) *options.FindOptions {
	return options.Find().SetProjection(projection)
}
