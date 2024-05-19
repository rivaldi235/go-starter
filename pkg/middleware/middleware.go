package middleware

import (
	"os"
	"service-code/model/dto/json"
	jwtClaims "service-code/model/dto/jwt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	user, password, ok := c.Request.BasicAuth()

	if !ok {
		json.NewResponseUnauthorized(c, "Invalid Token", "01", "01")
		c.Abort()
		return
	}
	if user != os.Getenv("CLIENT_ID") || password != os.Getenv("CLIENT_SECRET") {
		json.NewResponseUnauthorized(c, "Unauthorized", "01", "01")
		c.Abort()
		return
	}
	c.Next()
}

var (
	applicationName  = "incubation-golang"
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte("incubation-golang")
)

func GenerateTokenJwt(username string, expiredAt int64) (string, error) {
	loginExpDuration := time.Duration(expiredAt) * time.Minute
	myExpiresAt := time.Now().Add(loginExpDuration).Unix()

	claims := jwtClaims.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationName,
			ExpiresAt: myExpiresAt,
		},
		Username: username,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)
	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			json.NewResponseUnauthorized(c, "Invalid Token", "01", "01")
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &jwtClaims.JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSignatureKey, nil
		})

		if err != nil {
			json.NewResponseUnauthorized(c, "Invalid Token", "01", "01")
			c.Abort()
			return
		}

		if !token.Valid {
			json.NewResponseForbidden(c, "Forbidden", "01", "01")
			c.Abort()
			return
		}

		c.Next()
	}
}
