package helper

import (
	"log"
	"os"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllToken(email, firstName, lastName, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims :=&SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour *time.Duration(48)).Unix(),
		},
	}
	token,err :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString([]byte(SECRET_KEY))
	if err!=nil{
		log.Panic(err)
		return
	}
	refreshToken,err :=jwt.NewWithClaims(jwt.SigningMethodES256,refreshClaims).SignedString([]byte(SECRET_KEY))
	if err!=nil{
		log.Panic(err)
		return
	}
	return token,refreshToken,err
}

func UpdateAllTokens() {

}

func ValidateToken() {

}
