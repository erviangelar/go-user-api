package jwt

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/erviangelar/go-user-api/common/config"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	KeyType  string
	jwt.StandardClaims
}

type JWTRefreshClaim struct {
	ID      string `json:"id"`
	KeyType string
	jwt.StandardClaims
}

// generate access token
func GenerateAccessToken(user *models.User) (string, error) {

	configs := config.LoadConfig()
	userID := user.UID.String()
	tokenType := "access"

	claims := JWTClaim{
		userID,
		user.Username,
		user.Name,
		strings.Join(user.Role, ","),
		tokenType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(configs.JwtExpiration)).Unix(),
			Issuer:    "auth.service",
		},
	}

	signBytes, err := os.ReadFile(configs.AccessTokenPrivateKeyPath)
	if err != nil {
		return "", errors.New("could not generate access token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", errors.New("could not generate access token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// generate refresh token
func GenerateRefreshToken(user *models.User) (string, error) {
	configs := config.LoadConfig()
	userID := user.UID.String()
	tokenType := "refresh"

	claims := JWTRefreshClaim{
		userID,
		// cusKey,
		tokenType,
		jwt.StandardClaims{
			Issuer: "auth.service",
		},
	}

	signBytes, err := os.ReadFile(configs.RefreshTokenPrivateKeyPath)
	if err != nil {
		return "", errors.New("could not generate refresh token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", errors.New("could not generate refresh token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(signKey)
}

// validate access token
func ValidateAccessToken(tokenString string) (*models.User, error) {
	configs := config.LoadConfig()
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method in auth token")
		}
		verifyBytes, err := os.ReadFile(configs.AccessTokenPublicKeyPath)
		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid || claims.ID == "" || claims.KeyType != "access" {
		return nil, errors.New("invalid token: authentication failed")
	}
	user := models.User{}
	user.UID = uuid.FromStringOrNil(claims.ID)
	user.Role = strings.Split(claims.Role, ";")
	user.Name = claims.Name
	return &user, nil
}

// validate refresh token
func ValidateRefreshToken(tokenString string) (string, error) {

	configs := config.LoadConfig()
	token, err := jwt.ParseWithClaims(tokenString, &JWTRefreshClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method in auth token")
		}
		verifyBytes, err := os.ReadFile(configs.RefreshTokenPublicKeyPath)
		if err != nil {
			return nil, errors.New("invalid token: authentication failed")
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			return nil, errors.New("invalid token: authentication failed")
		}

		return verifyKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*JWTRefreshClaim)
	if !ok || !token.Valid || claims.ID == "" || claims.KeyType != "refresh" {
		return "", errors.New("invalid token: authentication failed")
	}
	return claims.ID, nil
}
