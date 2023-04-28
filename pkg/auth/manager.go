package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

const timeAlive = 30

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) *Manager {
	if signingKey == "" {
		return nil
	}

	return &Manager{signingKey: signingKey}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId     int64 `json:"user_id"`
	UserRoleId int64 `json:"user_role_id"`
}

func (m *Manager) GenerateJWT(id int64, role int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeAlive * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix()}, id, role})
	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) RefreshJWT() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (m *Manager) Parse(accessToken string) (interface{}, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i any, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["user_id"], nil
}
