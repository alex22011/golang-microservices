package services

import (
	"github.com/alex22011/golang-microservices/mvc/domain"
	"github.com/alex22011/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User,*utils.ApplicationError) {

	return domain.GetUser(userId)
}