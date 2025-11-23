package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rest/internal/config"
	"rest/internal/dto"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"rest/internal/storage/postgres/user"
	httpUtils "rest/internal/utils/http"
)

type contextKey string

const CtxUserKey contextKey = "user"

func CreateJWT(secret []byte, userID int) (string, error) {
	expirationTime := config.AuthConfig.JWTExpire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expirationTime).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, userRepo user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token, %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("invalid token, %v", err)
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID, err := strconv.Atoi(claims["userId"].(string))
		if err != nil {
			log.Printf("failed to get user id from token, %v", err)
			permissionDenied(w)
			return
		}

		user, err := userRepo.GetById(userID)
		if err != nil {
			log.Printf("failed to get user by id, %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, CtxUserKey, user)

		handlerFunc(w, r.WithContext(ctx))
	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.AuthConfig.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	httpUtils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserFromCtx(r *http.Request) *dto.User {
	user, ok := r.Context().Value(CtxUserKey).(*dto.User)
	if !ok {
		return new(dto.User)
	}
	return user
}
