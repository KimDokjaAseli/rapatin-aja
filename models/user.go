package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email            string             `bson:"email" json:"email"`
	Password         string             `bson:"password" json:"-"`
	Name             string             `bson:"name" json:"name"`
	Position         string             `bson:"position" json:"position"`
	Bio              string             `bson:"bio" json:"bio"`
	ProfileImagePath string             `bson:"profileImagePath" json:"profileImagePath"`
	Role             string             `bson:"role" json:"role"`
}
