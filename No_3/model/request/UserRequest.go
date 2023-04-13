package request

type UserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	RoleID   int64  `json:"roleID"`
}
