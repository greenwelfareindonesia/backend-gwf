package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int, role int) (string, error)
	ValidasiToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("greenWelfareKuhh")

func NewService() *jwtService {
	return &jwtService{}
}

type Claims struct {
	UserID int `json:"user_id"`
	Role   int `json:"role"`
	jwt.StandardClaims
}

// func (s *jwtService) GenerateToken(userID int, role int) (string, error) {
// 	claim := jwt.MapClaims{}
// 	claim["user_id"] = userID
// 	claim["role"] = role
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// 	signedToken, err := token.SignedString(SECRET_KEY)
// 	if err != nil {
// 		return signedToken, err
// 	}
// 	return signedToken, nil
// }

func (s *jwtService) GenerateToken(userID int, role int) (string, error) {
	expirationTime := time.Now().Add(10 * time.Hour)
	claim := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidasiToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
