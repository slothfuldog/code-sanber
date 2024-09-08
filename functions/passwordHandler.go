package functions

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/argon2"
)

func PasswordGenerator(password string) (encryptedPass string, e error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("(PASSGENERATOR: 1000) %s\n", err)
		return "", err
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		fmt.Printf("(PASSGENERATOR:1001) %s\n", er)
		return "", err
	}

	encrypted := argon2.IDKey([]byte(password), []byte(os.Getenv("salt")), 2, 64*1024, 8, 32)

	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)

	return string(encryptedBase64), nil
}

func VerifyPassword(encryptedPass, password string) (isMatch bool, e error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("(VERIFYPASS: 1000) ", err)
		return false, err
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		fmt.Println("(VERIFYPASS:1001) ", er)
		return false, er
	}

	decodedPass, err := base64.StdEncoding.DecodeString(encryptedPass)

	if err != nil {
		fmt.Println("(VERIFYPASS:1003) Error decoding base64: ", err)
		return false, err
	}

	encrypted := argon2.IDKey([]byte(decodedPass), []byte(os.Getenv("salt")), 2, 64*1024, 8, 32)

	if string(decodedPass) != string(encrypted) {
		fmt.Println("(VERIFYPASS:1002) PASSWORD IS NOT MATCH")
		return false, fmt.Errorf("PASSWORD IS NOT MATCH")
	}

	return true, nil
}
