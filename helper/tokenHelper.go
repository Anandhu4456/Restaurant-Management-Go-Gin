package helper

import (
	"os"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct{
	Email string
	First_name string
	Last_name string
	Uid string
	jwt.StandardClaims
}
var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllToken(){
	
}

func UpdateAllTokens(){

}

func ValidateToken(){
	
}