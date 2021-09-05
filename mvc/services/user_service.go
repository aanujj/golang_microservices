package services

import (
	"github.com/golang_microservices/mvc/domain"
	"github.com/golang_microservices/mvc/utils"
)

func GetUser(userID int) (*domain.User,*utils.MicroserviceError){
	return domain.GetEndUser(userID)
}

