package myrouter

import (
	"github.com/gorilla/mux"
	"golang-crud-clean-architecture/config/db"
	"golang-crud-clean-architecture/handler"
	"golang-crud-clean-architecture/repository"
	"golang-crud-clean-architecture/service"
	"net/http"
)

func NewRooter(rooter *mux.Router) *mux.Router {

	NewUserHandler(rooter)
	NewRoleHandler(rooter)

	return rooter
}

func NewUserHandler(router *mux.Router) *mux.Router {
	userRepo := repository.NewUserRepository(db.DB)
	roleRepo := repository.NewRoleRepository(db.DB)
	var userService = service.NewUserService(userRepo, roleRepo)
	var h = handler.NewUserHandler(userService)

	router.HandleFunc("/v1/users", h.CreateUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/v1/users/auth", h.Login).Methods(http.MethodPost)

	router.HandleFunc("/v1/users", h.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/{id}", h.GetUserByID).Methods(http.MethodGet)

	router.HandleFunc("/v1/users/{id}", h.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/v1/users/{id}", h.DeleteUser).Methods(http.MethodDelete)

	return router
}

func NewRoleHandler(router *mux.Router) *mux.Router {

	roleRepo := repository.NewRoleRepository(db.DB)
	var roleService = service.NewRoleService(roleRepo)
	var h = handler.NewRoleHandler(roleService)

	router.HandleFunc("/v1/roles", h.CreateRole).Methods(http.MethodPost)
	router.HandleFunc("/v1/roles", h.GetRoles).Methods(http.MethodGet)
	router.HandleFunc("/v1/roles/{id}", h.GetRoleByID).Methods(http.MethodGet)
	router.HandleFunc("/v1/roles/{id}", h.UpdateRole).Methods(http.MethodPut)
	router.HandleFunc("/v1/roles/{id}", h.DeleteRole).Methods(http.MethodDelete)

	return router
}
