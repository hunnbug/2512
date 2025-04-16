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

	//
	//объявляем структуру, в объект которой будут парситься данные о логине и пароле с фронта
	//
	type loggedUser struct {
		Username string
		Password string
	}

	var _loggedUser loggedUser

	err := ctx.BindJSON(&_loggedUser)

	if err != nil {

		logging.WriteLog("произошла ошибка при парсинге логина и пароля: ", err)

	}

	//
	//объект, в который будет пароситься найденный пользователь из БД
	//
	var user models.User

	tx := database.DB.Begin()

	//
	//запрос поиска записи в БД по логину, парсится в user, возвращается ошибка и проверяется
	//
	err = tx.First(&user, "username = ?", _loggedUser.Username).Error

	if err != nil {

		logging.WriteLog("неверный логин: ", err)

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

		logging.WriteLog("пароль не совпадает с хешем")

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неверный пароль!"})

		return
	}

	//
	//создаем пейлод для жвт токена
	//
	payload := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	}

	//
	//создаем токен
	//
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//
	//подписываем токен
	//
	token, err := t.SignedString(environment.Env.JwtToken)

	if err != nil {

		logging.WriteLog("ошибка во время создания токена: ", err)

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Err: err, Message: "невозможно создать токен"})

		return
	}

	//
	//объявляем структуру для того, чтоб отдавать токен на фронт
	//
	type responseToken struct {
		Token string
	}

	logging.WriteLog("Авторизация пользователя", token[0:38])

	ctx.JSON(200, responseToken{Token: token})
}
