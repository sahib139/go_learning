package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	MovieId   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieName string             `json:"movieName,omitempty"`
	IsWatched bool               `json:"isWatched"`
}
