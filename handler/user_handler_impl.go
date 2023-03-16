package handler

import (
	"net/http"
	"r-G7D/go_restful/domain/web"
	"r-G7D/go_restful/helper"
	"r-G7D/go_restful/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

func (handler *UserHandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := handler.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "User Created",
		Data:   userResponse,
	}
	helper.WriteFromRequestBody(writer, webResponse)
}

func (handler *UserHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	handler.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "User Deleted",
	}
	helper.WriteFromRequestBody(writer, webResponse)
}

func (handler *UserHandlerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userUpdateRequest.ID = id

	userResponse := handler.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "User Updated",
		Data:   userResponse,
	}
	helper.WriteFromRequestBody(writer, webResponse)
}

func (handler *UserHandlerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		helper.PanicIfError(err)
	}
	userResponse := handler.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteFromRequestBody(writer, webResponse)
}

func (handler *UserHandlerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := handler.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}
	helper.WriteFromRequestBody(writer, webResponse)
}
