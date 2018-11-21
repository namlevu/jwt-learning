package auth

import (
  "log"
  "time"
  "fmt"
  
  jwt "github.com/dgrijalva/jwt-go"
)

const SecretKey = "anhmeodeptrainhatxom!#%&(0)"

func GenerateJWT(username string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
  claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) *jwt.Token {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      log.Println("Token parse error")
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
    log.Println("Token parse error")
    return nil
	}
  return token
}
