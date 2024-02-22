package user

import "time"

type UserFormatter struct {
	ID        int       `json:"ID"`
	Username  string    `json:"Name"`
	Email     string    `json:"Email"`
	Token     string    `json:"Token"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
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
