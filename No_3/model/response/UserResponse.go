package response

type UserResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	RoleID   int64  `json:"roleID"`
	RoleName string `json:"roleName"`
}
