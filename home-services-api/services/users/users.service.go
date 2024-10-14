package users

import (
	"os"
	"time"

	usersDto "github.com/dot-slash-ann/home-services-api/dtos/users"
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersService interface {
	SignUp(usersDto.CreateUserDto) (entities.User, error)
	Login(usersDto.LoginUserDto) (entities.User, string, error)
	FindAll() ([]entities.User, error)
	FindOne(string) (entities.User, error)
	Update(string, usersDto.UpdateUserDto) (entities.User, error)
}

type UsersServiceImpl struct {
	database *gorm.DB
}

func NewUsersService(database *gorm.DB) *UsersServiceImpl {
	return &UsersServiceImpl{
		database: database,
	}
}

func (service UsersServiceImpl) SignUp(createUserDto usersDto.CreateUserDto) (entities.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 10)

	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		Email:    createUserDto.Email,
		Password: string(hash),
	}

	if result := service.database.Create(&user); result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}

func (service UsersServiceImpl) Login(loginUserDto usersDto.LoginUserDto) (entities.User, string, error) {
	var user entities.User

	if result := service.database.First(&user, "email = ?", loginUserDto.Email); result.Error != nil {
		return entities.User{}, "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserDto.Password))

	if err != nil {
		return entities.User{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * 5).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return entities.User{}, "", err
	}

	if result := service.database.Model(&user).Updates(entities.User{
		SessionToken: tokenString,
	}); result.Error != nil {
		return entities.User{}, "", result.Error
	}

	return user, tokenString, nil
}

func (service UsersServiceImpl) FindAll() ([]entities.User, error) {
	var usersList []entities.User

	if result := service.database.Find(&usersList); result.Error != nil {
		return []entities.User{}, result.Error
	}

	return usersList, nil
}

func (service UsersServiceImpl) FindOne(id string) (entities.User, error) {
	var user entities.User

	if result := service.database.First(&user, id); result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}

func (service UsersServiceImpl) Update(id string, updateUserDto usersDto.UpdateUserDto) (entities.User, error) {
	var user entities.User

	updatedUser := entities.User{
		SessionToken: updateUserDto.SessionToken,
	}

	if result := service.database.First(&user, id); result.Error != nil {
		return entities.User{}, result.Error
	}

	if result := service.database.Model(&user).Updates(updatedUser); result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}
