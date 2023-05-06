package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FormatterUser(user User, Token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    Token,
	}
	return formatter
}
