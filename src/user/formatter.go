package user

import "go-technical-test-synapsis/src/entity"

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatUser := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	}

	return formatUser
}
