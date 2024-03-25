package middlewares

import (
	"coffee-pos-backend/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": " Authorization header is required "})
			ctx.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authrization header must header must be in the format 'Bearer <token>'"})
			ctx.Abort()
			return
		}

		tokenStr := strings.TrimSpace(bearerToken[1])

		/// Verfify JWT
		token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			ctx.Set("username", claims.Username)
			ctx.Next()
		} else {
			log.Printf("Token claims: %#v, valid: %v\n", token.Claims, token.Valid)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token2"})
			ctx.Abort()
			return
		}

	}
}
