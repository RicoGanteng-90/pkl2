package service

import (
	"github.com/labstack/gommon/log"
	"golang-crud-clean-architecture/model/entity"
	"golang-crud-clean-architecture/model/request"
	"golang-crud-clean-architecture/model/response"
	"golang-crud-clean-architecture/repository"
)

type IRoleService interface {
	CreateRole(RoleRequest *request.RoleRequest) (*response.RoleResponse, error)
	GetRoles() []response.RoleResponse
	GetRoleById(id int64) *response.RoleResponse
	UpdateRole(roleRequest *request.RoleRequest, id int64) *response.RoleResponse
	DeleteRole(id int64)
}

type roleService struct {
	roleRepository repository.IRoleRepository
}

func NewRoleService(roleRepository repository.IRoleRepository) IRoleService {
	return &roleService{
		roleRepository: roleRepository,
	}
}

func (r roleService) CreateRole(roleRequest *request.RoleRequest) (*response.RoleResponse, error) {
	log.Info("ActionLog.CreateRole.Start")
	role := &entity.Role{}
	roleResponse := &response.RoleResponse{}
	role.Name = roleRequest.Name
	role, err := r.roleRepository.Save(role)
	roleResponse.Id = role.Id
	roleResponse.Name = role.Name
	log.Info("ActionLog.CreateRole.End")
	return roleResponse, err
}

func (r roleService) GetRoles() []response.RoleResponse {
	log.Info("ActionLog.GetRoles.Start")

	roles := make([]*entity.Role, 0)
	roleResponses := make([]response.RoleResponse, 0)
	roles, _ = r.roleRepository.FindAll(roles)

	for i := range roles {
		roleResponse := response.RoleResponse{}
		roleResponse.Id = roles[i].Id
		roleResponse.Name = roles[i].Name
		roleResponses = append(roleResponses, roleResponse)
	}
	log.Info("ActionLog.GetRoles.End")
	return roleResponses
}

func (r roleService) GetRoleById(id int64) *response.RoleResponse {
	log.Info("ActionLog.GetRoleById.Start")

	roleResponse := &response.RoleResponse{}
	role, _ := r.roleRepository.FindByRoleId(id)
	roleResponse.Id = role.Id
	roleResponse.Name = role.Name

	log.Info("ActionLog.GetRoleById.End")

	return roleResponse
}

func (r roleService) UpdateRole(roleRequest *request.RoleRequest, id int64) *response.RoleResponse {
	log.Info("ActionLog.UpdateRole.Start")

	role := &entity.Role{}
	roleResponse := &response.RoleResponse{}
	role.Name = roleRequest.Name
	role, _ = r.roleRepository.UpdateRole(role, id)
	roleResponse.Id = role.Id
	roleResponse.Name = role.Name
	log.Info("ActionLog.UpdateRole.End")
	return roleResponse
}

func (r roleService) DeleteRole(id int64) {
	log.Info("ActionLog.DeleteRole.Start")
	r.roleRepository.DeleteRole(id)
	log.Info("ActionLog.DeleteRole.End")
}
