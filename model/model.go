package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Todo      string              `json:"text" bson:"text,omitempty"`
	CreatedAt primitive.Timestamp `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt primitive.Timestamp `json:"updatedAt" bson:"updatedAt,omitempty"`
	IsDone    bool                `json:"isDone" bson:"isDone,omitempty"`
}
