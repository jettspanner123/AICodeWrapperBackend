package lib

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	jwtModels "github.com/jettspanner123/AICodeWrapperBackend/models/jwt"
)

var ACCESS_JWT_SECRET = []byte(os.Getenv("ACCESS_JWT_SECRET"))
var REFRESH_JWT_SECRET = []byte(os.Getenv("REFRESH_JWT_SECRET"))

func generateAccessToken(userClaims jwtModels.TokenClaims) (string, error) {
	claims := jwt.MapClaims{
		"username": userClaims.Username,
		"user_id":  userClaims.UserId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(ACCESS_JWT_SECRET)
}

func generateRefreshToken(userClaims jwtModels.TokenClaims) (string, error) {
	claims := jwt.MapClaims{
		"username": userClaims.Username,
		"userId":   userClaims.UserId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(REFRESH_JWT_SECRET)
}

func GenerateTokenPair(userClaims jwtModels.TokenClaims) (*jwtModels.TokenPair, error) {
	accessToken, err := generateAccessToken(userClaims)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateRefreshToken(userClaims)
	if err != nil {
		return nil, err
	}
	return &jwtModels.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func ValidateToken(options jwtModels.TokenValidationOptions) (*jwt.Token, error) {
	secret := ACCESS_JWT_SECRET
	if options.IsRefresh {
		secret = REFRESH_JWT_SECRET
	}

	return jwt.Parse(options.Token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
