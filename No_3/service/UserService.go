package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"golang-crud-clean-architecture/model"
	"golang-crud-clean-architecture/model/entity"
	"golang-crud-clean-architecture/model/request"
	"golang-crud-clean-architecture/model/response"
	"golang-crud-clean-architecture/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type IUserService interface {
	CreateUser(userRequest *request.UserRequest) (*response.UserResponse, error)
	GetUsers() []response.UserResponse
	GetUserById(id int64) (*response.UserResponse, error)
	UpdateUser(userRequest *request.UserRequest, id int64) *response.UserResponse
	DeleteUser(id int64)
	Login(loginRequest *request.LoginRequest) (*response.LoginResponse, error)
	TokenIsValid(tokeStr string) bool
	GetRoleName(tokenStr string) string
}

type userService struct {
	UserRepo repository.IUserRepository
	RoleRepo repository.IRoleRepository
}

func NewUserService(userRepo repository.IUserRepository, roleRepo repository.IRoleRepository) IUserService {
	return &userService{
		UserRepo: userRepo,
		RoleRepo: roleRepo,
	}
}

func (u *userService) CreateUser(userRequest *request.UserRequest) (*response.UserResponse, error) {
	log.Info("ActionLog.CreateUser.Start")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	userRequest.Password = string(hashedPassword)

	user := &entity.User{}
	userResponse := &response.UserResponse{}

	user.Name = userRequest.Name
	user.LastName = userRequest.LastName
	user.UserName = userRequest.UserName
	user.Password = userRequest.Password
	user.Age = userRequest.Age
	user.RoleID = userRequest.RoleID
	user.Email = userRequest.Email
	user, err := u.UserRepo.Save(user)
	if err != nil {
		return nil, err
	}
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.LastName = user.LastName
	userResponse.UserName = user.UserName
	userResponse.Age = user.Age
	userResponse.RoleID = user.RoleID
	log.Info("ActionLog.CreateUser.End")
	return userResponse, err
}

func (u *userService) GetUsers() []response.UserResponse {
	log.Info("ActionLog.GetUsers.Start")

	users := make([]*entity.User, 0)

	userResponses := make([]response.UserResponse, 0)

	users, _ = u.UserRepo.FindAll(users)

	for i := range users {
		userResponse := response.UserResponse{}
		role, _ := u.RoleRepo.FindByRoleId(users[i].RoleID)

		userResponse.Id = users[i].Id
		userResponse.Name = users[i].Name
		userResponse.LastName = users[i].LastName
		userResponse.UserName = users[i].UserName
		userResponse.Email = users[i].Email
		userResponse.Age = users[i].Age
		userResponse.RoleID = role.Id
		userResponse.RoleName = role.Name
		userResponses = append(userResponses, userResponse)
	}

	log.Info("ActionLog.GetUsers.End")
	return userResponses
}

func (u *userService) GetUserById(id int64) (*response.UserResponse, error) {
	log.Info("ActionLog.GetUserById.Start")

	userResponse := &response.UserResponse{}
	user, err := u.UserRepo.FindByUserId(id)
	if err != nil {
		return nil, err
	}
	userResponse.RoleID = user.RoleID
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.LastName = user.LastName
	userResponse.UserName = user.UserName
	userResponse.Age = user.Age

	log.Info("ActionLog.GetUserById.End")
	return userResponse, err
}

func (u *userService) UpdateUser(userRequest *request.UserRequest, id int64) *response.UserResponse {
	log.Info("ActionLog.UpdateUser.Start")
	user := &entity.User{}
	userResponse := &response.UserResponse{}

	user.Name = userRequest.Name
	user.LastName = userRequest.LastName
	user.UserName = userRequest.UserName
	user.Password = userRequest.Password
	user.Age = userRequest.Age

	user, _ = u.UserRepo.UpdateUser(user, id)
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.LastName = user.LastName
	userResponse.UserName = user.UserName
	userResponse.Age = user.Age
	log.Info("ActionLog.UpdateUser.End")
	return userResponse
}

func (u *userService) DeleteUser(id int64) {
	log.Info("ActionLog.DeleteUser.Start")
	_ = u.UserRepo.DeleteUser(id)
	log.Info("ActionLog.DeleteUser.End")
}

func (u *userService) Login(loginRequest *request.LoginRequest) (*response.LoginResponse, error) {
	user := &entity.User{}

	user.Email = loginRequest.Email
	user.Password = loginRequest.Password

	user, _ = u.UserRepo.Login(user)

	role, _ := u.RoleRepo.FindByRoleId(user.RoleID)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return nil, err
	}
	expirationTime := time.Now().Add(time.Minute * 5).Unix()
	expirationTimeRefresh := time.Now().Add(time.Hour).Unix()

	claims := sendClaims(user, role, expirationTime)
	refreshClaims := sendClaims(user, role, expirationTimeRefresh)

	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtKey)

	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tr, err := tokenRefresh.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	loginResponse := &response.LoginResponse{AccessToken: t, RefreshToken: tr}

	return loginResponse, err
}

func (u *userService) TokenIsValid(tokeStr string) bool {
	t, err := jwt.Parse(tokeStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false
	}
	return t.Valid
}

func (u *userService) GetRoleName(tokenStr string) string {
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	return role
}

func sendClaims(user *entity.User, role *entity.Role, expirationTime int64) *model.JwtCustomClaims {

	refreshClaims := &model.JwtCustomClaims{
		FullName: user.Name + " " + user.LastName,
		Username: user.UserName,
		Email:    user.Email,
		Role:     role.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	return refreshClaims
}
