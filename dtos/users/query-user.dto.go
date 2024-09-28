package UsersDto

import (
	UsersEntity "github.com/dot-slash-ann/home-services-api/entities/users"
	"github.com/gin-gonic/gin"
)

func UserToJson(user UsersEntity.User) gin.H {
	return gin.H{
		"id":    user.ID,
		"email": user.Email,
	}
}
