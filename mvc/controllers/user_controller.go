package controllers

import (
	"encoding/json"
	"github.com/golang_microservices/mvc/services"
	"github.com/golang_microservices/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	log.Println("user service started . . .")
	userId := request.URL.Query().Get("user_id")
	userInt, err := strconv.Atoi(userId)

	if err != nil {
		error := &utils.MicroserviceError{
			Message: "UserID must be integer",
			Status:  http.StatusBadRequest,
			Code:    "bad request ",
		}
		jsonValue, _ := json.Marshal(error)
		response.WriteHeader(error.Status)
		response.Write(jsonValue)
	}
	updateduser, error := services.GetUser(userInt)
	if error != nil {
		response.WriteHeader(error.Status)
		response.Write([]byte(error.Message))
	}
	jsonValue, err := json.Marshal(updateduser)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
	response.Write(jsonValue)

}
