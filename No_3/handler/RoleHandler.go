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
)

type RoleHandler struct {
	service service.IRoleService
}

func NewRoleHandler(roleService service.IRoleService) *RoleHandler {
	return &RoleHandler{service: roleService}
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role request.RoleRequest
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid request body"))
		return
	}
	resp, err := h.service.CreateRole(&role)
	if err != nil {
		exception.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("user can not be created"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *RoleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := h.service.GetRoles()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *RoleHandler) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	resp := h.service.GetRoleById(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	roleRequest := &request.RoleRequest{}
	json.NewDecoder(r.Body).Decode(&roleRequest)
	h.service.UpdateRole(roleRequest, id)
	json.NewEncoder(w).Encode(roleRequest)
}

func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	h.service.DeleteRole(id)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("The User is Deleted Successfully!")
}
