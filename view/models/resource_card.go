package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResourceCard struct {
	ID       primitive.ObjectID
	Type     string
	Offering bool
}
