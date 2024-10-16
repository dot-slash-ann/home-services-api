package users

import (
	"errors"
	"net/http"

	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userService UsersService
}

func NewUsersController(service UsersService) *UsersController {
	return &UsersController{
		userService: service,
	}
}

func (controller *UsersController) SignUp(c *gin.Context) {
	var createUserDto CreateUserDto

	if err := c.ShouldBind(&createUserDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	user, err := controller.userService.SignUp(createUserDto)

	if err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": UserToJson(user),
	})
}

func (controller *UsersController) Login(c *gin.Context) {
	var loginUserDto LoginUserDto

	if err := c.ShouldBind(&loginUserDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	user, token, err := controller.userService.Login(loginUserDto)

	if err != nil {
		httpErr := httpErrors.UnauthorizedError(err, nil)

		c.Error(httpErr)

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", string(token), 60*60*24*7, "", "", true, true)

	//

	c.JSON(http.StatusOK, gin.H{
		"data": UserToJson(user),
	})
}

func (controller *UsersController) FindAll(c *gin.Context) {
	users, err := controller.userService.FindAll()

	if err != nil {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ManyUsersToJson(users),
	})
}

func (controller *UsersController) FindOne(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

		return
	}

	user, err := controller.userService.FindOne(id)

	if err != nil {
		httpErr := httpErrors.NotFoundError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": UserToJson(user),
	})
}
