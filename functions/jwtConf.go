package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	secretKey []byte
)

type Claims struct {
	UserData string `json:"user_data"` // Store JSON string
	jwt.StandardClaims
}

// Init function to load the secret key from .env
func init() {
	wd, _ := os.Getwd()

	curDir := fmt.Sprint(wd, "/.env")

	err := godotenv.Load(curDir)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the secret key from environment variables
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		log.Fatal("JWT_SECRET_KEY is not set in the environment variables")
	}
}

// EncodeJWT creates a new JWT token
func EncodeJWT(userData map[string]interface{}) (string, error) {
	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		return "", err
	}
	claims := Claims{
		UserData: string(userDataJSON),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
			Issuer:    "your_app",                            // Issuer of the token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// DecodeJWT parses and validates the JWT token
func DecodeJWT(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Convert JSON string back to map
		var userData map[string]interface{}
		err := json.Unmarshal([]byte(claims.UserData), &userData)
		if err != nil {
			return nil, err
		}
		return userData, nil
	}

	return nil, fmt.Errorf("invalid token")
}
