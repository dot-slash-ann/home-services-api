package UsersService

import (
	"os"
	"time"

	"github.com/dot-slash-ann/home-services-api/database"
	UsersDto "github.com/dot-slash-ann/home-services-api/dtos/users"
	UsersEntity "github.com/dot-slash-ann/home-services-api/entities/users"
	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(createUserDto UsersDto.CreateUserDto) (UsersEntity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 10)

	if err != nil {
		return UsersEntity.User{}, err
	}

	user := UsersEntity.User{
		Email:    createUserDto.Email,
		Password: string(hash),
	}

	if err := lib.HandleDatabaseError(database.Connection.Create(&user)); err != nil {
		return UsersEntity.User{}, err
	}

	return user, nil
}

func Login(loginUserDto UsersDto.LoginUserDto) (UsersEntity.User, string, error) {
	var user UsersEntity.User

	if err := lib.HandleDatabaseError(database.Connection.First(&user, "email = ?", loginUserDto.Email)); err != nil {
		return UsersEntity.User{}, "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserDto.Password))

	if err != nil {
		return UsersEntity.User{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// TODO: make this an env var secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return UsersEntity.User{}, "", err
	}

	return user, tokenString, nil
}

func FindAll() ([]UsersEntity.User, error) {
	var users []UsersEntity.User

	if err := lib.HandleDatabaseError(database.Connection.Find(&users)); err != nil {
		return []UsersEntity.User{}, err
	}

	return users, nil
}

func FindOne(id string) (UsersEntity.User, error) {
	var user UsersEntity.User

	if err := lib.HandleDatabaseError(database.Connection.First(&user, id)); err != nil {
		return UsersEntity.User{}, err
	}

	return user, nil
}
