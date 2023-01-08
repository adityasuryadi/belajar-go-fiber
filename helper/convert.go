package helpers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PrimitveObjToString(value interface{}) string {
	return value.(primitive.ObjectID).Hex()
}

func PrimitveObjToTIme(value interface{}) time.Time {
	return time.Unix(int64(value.(primitive.Timestamp).T), 0)
}

func PrimitveObjInt64ToTIme(value interface{}) time.Time {
	return time.Unix(int64(value.(primitive.DateTime)), 0)
}
