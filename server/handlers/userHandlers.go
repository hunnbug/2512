package handlers

import (
	"log"
	"main/database"
	"main/environment"
	"main/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PostUserHandler(ctx *gin.Context) {

	type loggedUser struct {
		Username string
		Password string
	}

	log.Println("получен запрос ")

	var _loggedUser loggedUser

	ctx.BindJSON(&_loggedUser)

	log.Println("логин: ", _loggedUser)

	var user models.User

	tx := database.DB.Begin()

	err := tx.First(&user, "username = ?", _loggedUser.Username).Error

	if err != nil {

		log.Println("пользователь не найден")

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неправильный логин!"})

		tx.Rollback()

		return

	}

	tx.Commit()

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(_loggedUser.Password)) // линьганьгулигулигуливочалиньганьгулиньганьгу

	if err != nil {
		log.Println("пароль не совпадает с хешем")

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неправильный пароль!"})

		return
	}

	log.Println("найдено совпадение по пользователю:", user)

	payload := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString(environment.Env.JwtToken)

	if err != nil {
		log.Println("ошибка во время подписи токена: ", err)

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Err: err, Message: "ошибка на стороне сервера, попробуйте снова позже"})

		return
	}

	log.Println("jwt токен:", token)

	type responseToken struct {
		Token string
	}

	ctx.JSON(200, responseToken{Token: token})
}
