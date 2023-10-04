package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID primitive.ObjectID `bson:"_id"`
	First_name *string `json:"first_name" validate:"required,min=2,max=20"`
	Last_name *string `json:"last_name" validate:"required,min=2,max=20"`
	Phone *string `json:"phone" validate:"required"`
	Email *string `json:"email" validate:"email,required"`
	Password *string `json:"password" validate:"required,min=6"`
	Avatar *string `json:"avatar"`
	Token *string `json:"token"`
	Refresh_token *string `json:"refresh_token"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User_id string `json:"user_id"`
}