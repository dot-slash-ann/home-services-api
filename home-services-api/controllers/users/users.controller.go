package UsersController

import (
	"net/http"

	UsersDto "github.com/dot-slash-ann/home-services-api/dtos/users"
	"github.com/dot-slash-ann/home-services-api/lib"
	UsersService "github.com/dot-slash-ann/home-services-api/services/users"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var createUserDto UsersDto.CreateUserDto

	if !lib.HandleShouldBind(c, &createUserDto) {
		return
	}

	user, err := UsersService.SignUp(createUserDto)

	if err != nil {
		lib.HandleError(c, http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": UsersDto.UserToJson(user),
	})
}

func Login(c *gin.Context) {
	var loginUserDto UsersDto.LoginUserDto

	if !lib.HandleShouldBind(c, &loginUserDto) {
		return
	}

	user, token, err := UsersService.Login(loginUserDto)

	if err != nil {
		lib.HandleError(c, http.StatusUnauthorized, err)

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", string(token), 60*60*24*7, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"data": UsersDto.UserToJson(user),
	})
}

func FindAll(c *gin.Context) {
	users, err := UsersService.FindAll()

	// TODO: this is a bad response
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": UsersDto.ManyUsersToJson(users),
	})
}

func FindOne(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	user, err := UsersService.FindOne(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": UsersDto.UserToJson(user),
	})
}
