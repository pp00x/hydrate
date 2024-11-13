package middleware

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/pp00x/hydrate/config"
)

func Authenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Validate the signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(config.AppConfig.JWT.SecretKey), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        userID := uint(claims["user_id"].(float64))
        c.Set("userID", userID)
        c.Next()
    }
}