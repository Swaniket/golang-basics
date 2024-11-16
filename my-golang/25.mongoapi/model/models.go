package model

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson/primitive"
)

type TNetflix struct {
	ID      primitive.ObjectID `json: "_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json: "movie,omitempty"`
	Watched bool               `json: "watched,omitempty"`
}
