package auth

import (
	"errors"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// SignedDetails
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateToken generates a jwt token
func GenerateToken(email string, firstName string, lastName string, uid string) (signedToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return signedToken, err
}

//ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err = errors.New("Couldn't get signed details")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}

	return claims, err
}
