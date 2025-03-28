package token

import "math/rand"

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateAuthTokenAndRefreshToken(length int) (string, string) {
	authToken := randStringBytes(length)
	refreshToken := randStringBytes(length)

	return authToken, refreshToken

}

func GetAuthCode(length int) string {
	return randStringBytes(length)
}

func randStringBytes(length int) string {
	token := make([]byte, length)

	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}
	return string(token)
}
