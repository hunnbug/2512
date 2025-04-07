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

	err := logging.WriteLog("получен запрос")

	logging.CheckLogError(err)

	var _loggedUser loggedUser

	err = ctx.BindJSON(&_loggedUser)

	if err != nil {

		e := logging.WriteLog("произошла ошибка при парсинге логина и пароля: ", err)

		logging.CheckLogError(e)
	}

	err = logging.WriteLog("получен логин: ", _loggedUser)

	logging.CheckLogError(err)

	//
	//объект, в который будет пароситься найденный пользователь из БД
	//
	var user models.User

	//
	//начинаем транзакцию к БД
	//
	tx := database.DB.Begin()

	//
	//запрос поиска записи в БД по логину, парсится в user, возвращается ошибка и проверяется
	//
	err = tx.First(&user, "username = ?", _loggedUser.Username).Error

	if err != nil {

		e := logging.WriteLog("неверный логин: ", err)

		logging.CheckLogError(e)

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

		e := logging.WriteLog("пароль не совпадает с хешем")

		logging.CheckLogError(e)

		//при несовпадении пароля и хеша отдаем ошибку
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неверный пароль!"})

		return
	}

	err = logging.WriteLog("taken user:", user)

	logging.CheckLogError(err)

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

		e := logging.WriteLog("ошибка во время создания токена: ", err)

		logging.CheckLogError(e)

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Err: err, Message: "невозможно создать токен"})

		return
	}

	err = logging.WriteLog("jwt токен:", token)

	logging.CheckLogError(err)

	//
	//объявляем структуру для того, чтоб отдавать токен на фронт
	//
	type responseToken struct {
		Token string
	}

	ctx.JSON(200, responseToken{Token: token})
}
