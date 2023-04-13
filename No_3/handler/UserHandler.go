package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang-crud-clean-architecture/exception"
	"golang-crud-clean-architecture/model/request"
	"golang-crud-clean-architecture/service"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{service: userService}
}

var bearerPrefix = "Bearer "

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), bearerPrefix)[1]

	if h.service.GetRoleName(authToken) != "ADMIN" {
		exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("user role must be ADMIN"))
		return
	}

	if h.service.TokenIsValid(authToken) {
		var user request.UserRequest
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid request body"))
			return
		}
		resp, err := h.service.CreateUser(&user)
		if err != nil {
			exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("user can not be created"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	} else {
		exception.BearerUnauthorized(w)
		return
	}

}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), bearerPrefix)[1]
	if h.service.TokenIsValid(authToken) {
		resp := h.service.GetUsers()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	} else {
		exception.BearerUnauthorized(w)
		return
	}

}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	if h.service.TokenIsValid(authToken) {
		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)
		resp, err := h.service.GetUserById(id)
		if err != nil {
			exception.SendErrorResponse(w, http.StatusNotFound, fmt.Sprintf("user not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	} else {
		exception.BearerUnauthorized(w)
		return
	}

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	if h.service.GetRoleName(authToken) != "ADMIN" {
		exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("user role must be ADMIN"))
		return
	}

	if h.service.TokenIsValid(authToken) {
		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)
		userRequest := &request.UserRequest{}
		json.NewDecoder(r.Body).Decode(&userRequest)
		h.service.UpdateUser(userRequest, id)
		json.NewEncoder(w).Encode(userRequest)
	} else {
		exception.BearerUnauthorized(w)
		return
	}

}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	if h.service.GetRoleName(authToken) != "ADMIN" {
		exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("user role must be ADMIN"))
		return
	}

	if h.service.TokenIsValid(authToken) {
		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)
		h.service.DeleteUser(id)
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("The User is Deleted Successfully!")
	} else {
		exception.BearerUnauthorized(w)
		return
	}

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&login)
	resp, err := h.service.Login(&login)
	if err != nil {
		exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("user can not be created"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
