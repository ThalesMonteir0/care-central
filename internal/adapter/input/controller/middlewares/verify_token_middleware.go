package middlewares

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

const (
	JWT_SECRET = "JWT_SECRET"
)

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError("invalid token")
		c.JSON(errRest.Code, errRest.Message)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := domain.UserDomain{
		ID:       int(claims["id"].(float64)),
		Name:     claims["name"].(string),
		Email:    claims["email"].(string),
		ClinicID: int(claims["clinic_id"].(float64)),
		Role:     claims["role"].(string),
	}

	c.Set("clinic_id", userDomain.ClinicID)
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix(token, "Bearer")
		token = strings.TrimSpace(token)
	}

	return token
}
