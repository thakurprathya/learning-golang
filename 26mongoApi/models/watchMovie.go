package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchMovie struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
