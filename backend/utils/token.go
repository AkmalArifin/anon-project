package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"os"
	"strconv"
	"strings"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateResetPasswordToken(userID int64) string {

	encodeUserID := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(userID, 10)))

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(encodeUserID))
	token := mac.Sum(nil)

	output := encodeUserID + "." + base64.StdEncoding.EncodeToString(token)

	return output
}

func VerifyResetPasswordToken(token string) (bool, error) {
	parsedToken := strings.Split(token, ".")

	messageToken, err := base64.StdEncoding.DecodeString(parsedToken[1])

	if err != nil {
		return false, errors.New("decoded falied")
	}

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(parsedToken[0]))
	expectedToken := mac.Sum(nil)

	isValid := hmac.Equal(expectedToken, messageToken)

	if !isValid {
		return false, errors.New("hmac not equal")
	}

	return true, nil
}
