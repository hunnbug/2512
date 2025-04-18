package handlers

import (
	"main/database"
	"main/environment"
	"main/logging"
	"main/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func LoginHandler(ctx *gin.Context) {

	type loggedUser struct {
		Username string
		Password string
	}

	var _loggedUser loggedUser

	err := ctx.BindJSON(&_loggedUser)

	if err != nil {

		logging.WriteLog(logging.ERROR, "произошла ошибка при парсинге логина и пароля: ", err)

	}

	var user models.User

	tx := database.DB.Begin()

	err = tx.First(&user, "username = ?", _loggedUser.Username).Error

	if err != nil {

		logging.WriteLog(logging.ERROR, "неверный логин: ", err)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неверный логин!"})

		tx.Rollback()

		return

	}

	tx.Commit()

	//
	//сравниваем хеш в БД с полученным паролем с фронта
	//
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(_loggedUser.Password))

	if err != nil {

		logging.WriteLog(logging.ERROR, "пароль не совпадает с хешем")

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неверный пароль!"})

		return
	}

	payload := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString(environment.Env.JwtToken)

	if err != nil {

		logging.WriteLog(logging.ERROR, "ошибка во время создания токена: ", err)

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Err: err, Message: "невозможно создать токен"})

		return
	}

	type responseToken struct {
		Token string
	}

	logging.WriteLog(logging.DEBUG, "Авторизация пользователя", token[0:38])

	ctx.JSON(200, responseToken{Token: token})
}
