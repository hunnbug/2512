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

func LoginHandler(ctx *gin.Context) {

	//
	//объявляем структуру, в объект которой будут парситься данные о логине и пароле с фронта
	//
	type loggedUser struct {
		Username string
		Password string
	}

	//
	//пишем лог при старте работы хендлера
	//
	err := logging.WriteLog("получен запрос")

	logging.CheckLogError(err)

	//
	//объявляем объект структуры для дальнейшего парсинга логина и пароля
	//
	var _loggedUser loggedUser

	//
	// получаем логин и пароль и парсим в loggedUser
	//
	err = ctx.BindJSON(&_loggedUser)

	if err != nil {

		e := logging.WriteLog("произошла ошибка при парсинге логина и пароля: ", err)

		logging.CheckLogError(e)
	}

	//
	//пишем лог после парсинга логина и пароля в обхект структуры
	//
	err = logging.WriteLog("получен логин: ", _loggedUser)

	logging.CheckLogError(err)

	//
	//объект, в который будет пароситься найденный пользователь из БД
	//
	var user models.User

	database.DB.Transaction(func(tx *gorm.DB) error {

		//
		//ищем пользователя по полученному с фронта юзернейму
		//
		if err := tx.Find(&user, "username = ?", _loggedUser.Username).Error; err != nil {

			e := logging.WriteLog("неверный логин: ", err)

			logging.CheckLogError(e)

			//
			//при ненаходе отдаем ошибку и 400 статус
			//
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неверный логин!"})

			return err

		}

		return nil
	})

	//
	//сравниваем хеш в БД с полученным паролем с фронта
	//
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(_loggedUser.Password))

	if err != nil {

		e := logging.WriteLog("пароль не совпадает с хешем")

		logging.CheckLogError(e)

		//
		//при несовпадении пароля и хеша отдаем ошибку
		//
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "неверный пароль!"})

		return
	}

	//
	//пишем в лог данные пользователя при совпадении пароля и логина
	//
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

		//
		//при ошибке подписи токена возвращаем 400 и отдаем ошибку
		//
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Err: err, Message: "невозможно создать токен"})

		return
	}

	//
	//логгируем жвт токен
	//
	err = logging.WriteLog("jwt токен:", token)

	logging.CheckLogError(err)

	//
	//объявляем структуру для того, чтоб отдавать токен на фронт
	//
	type responseToken struct {
		Token string
	}

	//
	//отдаем токен на фронт с кодом 200
	//
	ctx.JSON(200, responseToken{Token: token})
}
