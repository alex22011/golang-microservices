package controllers

import (
	"encoding/json"
	"github.com/alex22011/golang-microservices/mvc/services"
	"github.com/alex22011/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)
func GetUser(resp http.ResponseWriter,req *http.Request) {
//userIdParam :=req.URL.Query().Get("user_id")
//fmt.Printf( "about to process user_id %v",userId)
userId,err :=strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
if  err!=nil {
	apiErr :=&utils.ApplicationError{
		Message : "userid must be a number",
		StatusCode : http.StatusNotFound,
		Code : "bad user",
	}
	jsonvalue,_ :=json.Marshal(apiErr)
	resp.WriteHeader(apiErr.StatusCode)
	resp.Write(jsonvalue)
	return
}
user ,apiErr :=services.GetUser(userId)
if apiErr!=nil {
	//handle the error and return to the clients
	jsonvalue,_ :=json.Marshal(apiErr)
	resp.WriteHeader(apiErr.StatusCode)
	resp.Write([]byte(jsonvalue))
	return
}
// return user to the client
jsonValue,_:=json.Marshal(user)
resp.Write(jsonValue)
}




