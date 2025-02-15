package handler

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`     //name of the meal
	MealType  string             `bson:"mealType"` //type of meal (lunch, dinner, breakfast)
	Date      time.Time          `bson:"date"`     //time eaten
	Image     string             `bson:"image"`    //image of the meal (URL or path to image)
	Macros    Macros             `bson:"macros"`   //macros for the meal
	CreatedAt time.Time          `bson:"createdAt"`
}

type Macros struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"` //name of the meal
	Calories  float64            `bson:"calories"`
	Protein   float64            `bson:"protein"`
	Carbs     float64            `bson:"carbs"`
	Fat       float64            `bson:"fat"`
	CreatedAt time.Time          `bson:"createdAt"`
}

//Used in bulk upload - not using id
type Macro struct {
	Name     string    `bson:"name"`
	Calories float64   `bson:"calories"`
	Protein  float64   `bson:"protein"`
	Carbs    float64   `bson:"carbs"`
	Fat      float64   `bson:"fat"`
	CreatedAt  time.Time `bson:"createdAt"`
}