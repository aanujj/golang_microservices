package domain

import (
	"fmt"
	"github.com/golang_microservices/mvc/utils"
	"net/http"
)

var (
	endUsers = map[int]*User{
		1999 : {Id: 1999,FirstName: "Italia",LastName: "SpaghetiMafia",Email: "anuj@Italia.com"},
		1998 : {Id: 1999,FirstName: "Portugal",LastName: "Euro2016",Email: "ronaldo@Portugal.com"},
	}
)
func GetEndUser(ID int) (*User,*utils.MicroserviceError){
	user := endUsers[ID]
	if user != nil{
		return user,nil
	}
	return nil, &utils.MicroserviceError{
		Message: fmt.Sprintf(" %v Id not Found",ID),
		Status: http.StatusNotFound,
		Code: "Not Found",
	}
}
