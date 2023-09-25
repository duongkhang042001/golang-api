package helpers

import (
	logger "core-api/pkg/logging"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var log *zap.SugaredLogger = logger.InitLogger()

func ComparePassword(hashedPwd string, plainPassword []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Errorln(err)
		return false
	}
	return true
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Errorln(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
