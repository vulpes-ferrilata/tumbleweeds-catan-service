package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Harbor struct {
	ID   primitive.ObjectID
	Q    int
	R    int
	Type string
}
