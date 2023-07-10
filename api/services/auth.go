package services

import (
	"fmt"
	"urlshorter/models"
	"urlshorter/utils"
)

type AuthService struct {
}

func (service *AuthService) Auth(user RegisterRequest) string {
	db := utils.GetDBConnection()
	defer db.Close()

	userModel := &models.UserModel{DB: db}

	tokenModel := &models.TokenModel{DB: db}

	userData, _ := userModel.Get(user.Name, user.Password)

	id := userData.ID

	if id == 0 {
		id = userModel.Create(user.Name, user.Password)
	}

	token, _ := tokenModel.Create(id)

	fmt.Println(tokenModel.GetUser("x2XjIOxYWa0xLD3IGfRItarS0qNsooqgv7l3u0NUre20KpHWHlv0OyCuWLSYy2GJ"))

	return token
}
