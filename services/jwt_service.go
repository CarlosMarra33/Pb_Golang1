package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "segredo",
		issure:    "pb-api",
	}
}

type Claim struct {
	Sum string `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(email string) (string, error) {
	claim := &Claim{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

func (s *jwtService) ExtractEmailFromToken(tokenString string) (string, error) {
	// Define a chave secreta usada para assinar o token JWT
	// Você deve usar a mesma chave secreta que foi usada para assinar o token
	secret := []byte(s.secretKey)

	// Analisa o token JWT e extrai as informações do payload
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}

	// Extrai o email do payload do token JWT
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, found := claims["email"].(string)
		if !found {
			return "", errors.New("o token JWT não contém um campo 'email")
		}
		return email, nil
	} else {
		return "", errors.New("o token JWT é inválido")
	}
}
