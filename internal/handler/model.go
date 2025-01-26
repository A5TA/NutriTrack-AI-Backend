package handler

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	ID        primitive.ObjectID `bson:"_id"`
	MealType  string             `bson:"mealType"` //type of meal (lunch, dinner, breakfast)
	Date      time.Time          `bson:"date"`     //time eaten
	Image     string             `bson:"image"`    //image of the meal (URL or path to image)
	CreatedAt time.Time          `bson:"createdAt"`
}
