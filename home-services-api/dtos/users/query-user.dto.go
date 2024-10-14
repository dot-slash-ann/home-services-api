package users

import (
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/gin-gonic/gin"
)

func UserToJson(user entities.User) gin.H {
	return gin.H{
		"id":    user.ID,
		"email": user.Email,
	}
}

func ManyUsersToJson(users []entities.User) []gin.H {
	results := make([]gin.H, 0, len(users))

	for _, user := range users {
		results = append(results, UserToJson(user))
	}

	return results
}
