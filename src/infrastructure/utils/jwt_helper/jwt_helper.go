package jwt_helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	configs "github.com/nhsh1997/go-boilerplate/config"
)


type IJwtHelper interface {
	SignIn(payload jwt.MapClaims) (string, error)
	Verify(tokenString string)
}

type JwtHelper struct {
	SecretKey string

}

func NewJwtHelper(config *configs.Configuration) IJwtHelper{
	return &JwtHelper{
		SecretKey: config.AuthSecret,
	}
}

func (j *JwtHelper) SignIn(payload jwt.MapClaims) (string, error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(j.SecretKey))

	return tokenString, err
}


func (j *JwtHelper) Verify(tokenString string){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}
