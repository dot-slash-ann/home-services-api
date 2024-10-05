package users

import (
	"os"
	"time"

	usersDto "github.com/dot-slash-ann/home-services-api/dtos/users"
	"github.com/dot-slash-ann/home-services-api/entities/users"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersService interface {
	SignUp(usersDto.CreateUserDto) (users.User, error)
	Login(usersDto.LoginUserDto) (users.User, string, error)
	FindAll() ([]users.User, error)
	FindOne(string) (users.User, error)
	Update(string, usersDto.UpdateUserDto) (users.User, error)
}

type UsersServiceImpl struct {
	database *gorm.DB
}

func NewUsersService(database *gorm.DB) *UsersServiceImpl {
	return &UsersServiceImpl{
		database: database,
	}
}

func (service UsersServiceImpl) SignUp(createUserDto usersDto.CreateUserDto) (users.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 10)

	if err != nil {
		return users.User{}, err
	}

	user := users.User{
		Email:    createUserDto.Email,
		Password: string(hash),
	}

	if result := service.database.Create(&user); result.Error != nil {
		return users.User{}, result.Error
	}

	return user, nil
}

func (service UsersServiceImpl) Login(loginUserDto usersDto.LoginUserDto) (users.User, string, error) {
	var user users.User

	if result := service.database.First(&user, "email = ?", loginUserDto.Email); result.Error != nil {
		return users.User{}, "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserDto.Password))

	if err != nil {
		return users.User{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * 5).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return users.User{}, "", err
	}

	if result := service.database.Model(&user).Updates(users.User{
		SessionToken: tokenString,
	}); result.Error != nil {
		return users.User{}, "", result.Error
	}

	return user, tokenString, nil
}

func (service UsersServiceImpl) FindAll() ([]users.User, error) {
	var usersList []users.User

	if result := service.database.Find(&usersList); result.Error != nil {
		return []users.User{}, result.Error
	}

	return usersList, nil
}

func (service UsersServiceImpl) FindOne(id string) (users.User, error) {
	var user users.User

	if result := service.database.First(&user, id); result.Error != nil {
		return users.User{}, result.Error
	}

	return user, nil
}

func (service UsersServiceImpl) Update(id string, updateUserDto usersDto.UpdateUserDto) (users.User, error) {
	var user users.User

	updatedUser := users.User{
		SessionToken: updateUserDto.SessionToken,
	}

	if result := service.database.First(&user, id); result.Error != nil {
		return users.User{}, result.Error
	}

	if result := service.database.Model(&user).Updates(updatedUser); result.Error != nil {
		return users.User{}, result.Error
	}

	return user, nil
}
