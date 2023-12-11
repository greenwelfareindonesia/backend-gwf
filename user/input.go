package user

type RegisterUserInput struct {
	Username string `json:"Username" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type UpdateUserInput struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
