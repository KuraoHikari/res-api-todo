package controller

import (
	"encoding/json"
	"golang-res-api-coba/helper"
	"golang-res-api-coba/model/web"
	"golang-res-api-coba/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl {
		UserService: userService,
	}
}
type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (controller *UserControllerImpl)Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	userCreateRequest := web.UserCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(request)
	helper.PanicError(err)

	userResponse := controller.UserService.Create(request.Context(),userCreateRequest)
	webRes := WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webRes)
	helper.PanicError(errEncode)
}
func (controller *UserControllerImpl)Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	userUpdateRequest := web.UserUpdateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(request)
	helper.PanicError(err)

	userId := params.ByName("userId")
	id,errStr := strconv.Atoi(userId)
	helper.PanicError(errStr)
	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest) 
	webRes := WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webRes)
	helper.PanicError(errEncode)
}
func (controller *UserControllerImpl)Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	userId := params.ByName("userId")
	id,errStr := strconv.Atoi(userId)
	helper.PanicError(errStr)
	controller.UserService.Delete(request.Context(), id)
	webRes := WebResponse{
		Code: 200,
		Status: "OK",
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webRes)
	helper.PanicError(errEncode)
}
func (controller *UserControllerImpl)FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
func (controller *UserControllerImpl)FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)