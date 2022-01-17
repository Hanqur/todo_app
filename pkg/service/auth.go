package service

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"time"

	"github.com/Hanqur/todo_app"
	"github.com/Hanqur/todo_app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	
)

const (
	salt = "dsklfj9345hfwksd"
	signingKey = "fjwe9395htew993nmfw356"
	tokenTTL = 12 * time.Hour
) 

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}


type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// хешируем пароль и вызываем метод CreateUser у репозитория
func (s *AuthService) CreateUser(user todo_app.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// функция хеширования пароля
func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// получение пользователя из базы
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		fmt.Println("tut")
		return "", err
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{ 
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		}, 
				user.Id,
	})

	return token.SignedString([]byte(signingKey))


}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}