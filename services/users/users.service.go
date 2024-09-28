package UsersService

import (
	"github.com/dot-slash-ann/home-services-api/database"
	UsersDto "github.com/dot-slash-ann/home-services-api/dtos/users"
	UsersEntity "github.com/dot-slash-ann/home-services-api/entities/users"
	"github.com/dot-slash-ann/home-services-api/lib"
)

func SignUp(createUserDto UsersDto.CreateUserDto) (UsersEntity.User, error) {
	user := UsersEntity.User{
		Email:    createUserDto.Email,
		Password: createUserDto.Password,
	}

	if err := lib.HandleDatabaseError(database.Connection.Create(&user)); err != nil {
		return UsersEntity.User{}, err
	}

	return user, nil
}
