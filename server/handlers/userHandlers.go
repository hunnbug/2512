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
	"gorm.io/gorm"
)

//TODO сделать вывод ошибок на фронт

func PostUserHandler(ctx *gin.Context) {

	type loggedUser struct {
		Username string
		Password string
	}

	err := logging.WriteLog("получен запрос")

	logging.CheckLogError(err)

	var _loggedUser loggedUser

	ctx.BindJSON(&_loggedUser)

	err = logging.WriteLog("получен логин: ", _loggedUser)

	logging.CheckLogError(err)

	var user models.User

	database.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Find(&user, "username = ?", _loggedUser.Username).Error; err != nil {

			err := logging.WriteLog("неверный логин: ", err)

			logging.CheckLogError(err)

			ctx.JSON(http.StatusNotAcceptable, models.ErrorResponse{Err: err, Message: "неверный логин!"})

			return err

		}

		return nil
	})

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(_loggedUser.Password)) // линьганьгулигулигуливочалиньганьгулиньганьгу

	if err != nil {
		err := logging.WriteLog("пароль не совпадает с хешем")

		logging.CheckLogError(err)

		ctx.JSON(http.StatusForbidden, models.ErrorResponse{Err: err, Message: "неверный пароль!"})

		return
	}

	err = logging.WriteLog("taken user:", user)

	logging.CheckLogError(err)

	payload := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString(environment.Env.JwtToken)

	if err != nil {

		err := logging.WriteLog("ошибка во время создания токена: ", err)

		logging.CheckLogError(err)

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Err: err, Message: "невозможно создать токен"})

		return
	}

	err = logging.WriteLog("jwt токен:", token)

	logging.CheckLogError(err)

	type responseToken struct {
		Token string
	}

	ctx.JSON(200, responseToken{Token: token})
}
