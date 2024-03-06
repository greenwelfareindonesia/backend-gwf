package user

import "time"

type UserFormatter struct {
	ID        int       `json:"ID"`
	Username  string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func PostFormatterUser(user User, Token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     Token,
		CreatedAt: user.CreatedAt,
	}
	return formatter
}

func UpdatedFormatterUser(user User, Token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     Token,
		UpdatedAt: user.UpdatedAt,
	}
	return formatter
}
