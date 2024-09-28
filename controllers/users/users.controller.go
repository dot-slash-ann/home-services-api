package UsersController

import (
	"net/http"

	UsersDto "github.com/dot-slash-ann/home-services-api/dtos/users"
	"github.com/dot-slash-ann/home-services-api/lib"
	UsersService "github.com/dot-slash-ann/home-services-api/services/users"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var createUserDto UsersDto.CreateUserDto

	if !lib.HandleShouldBind(c, &createUserDto) {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 10)

	// TODO: this is a bad response code
	if err != nil {
		lib.HandleError(c, http.StatusUnauthorized, err)

		return
	}

	createUserDto.Password = string(hash)

	user, err := UsersService.SignUp(createUserDto)

	if err != nil {
		lib.HandleError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": UsersDto.UserToJson(user),
	})
}
