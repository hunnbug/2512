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

	log.Println("took POST req")

	var _loggedUser loggedUser

	ctx.BindJSON(&_loggedUser)

	log.Println("logged user: ", _loggedUser)

	var user models.User

	tx := database.DB.Begin()

	err := tx.First(&user, "username = ?", _loggedUser.Username).Error

	if err != nil {

		log.Println("user not found")

		ctx.JSON(http.StatusNotAcceptable, models.ErrorResponse{Err: err, Message: "cannot find user by login"})

		return

	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(_loggedUser.Password)) // линьганьгулигулигуливочалиньганьгулиньганьгу

	if err != nil {
		log.Println("given password does not match hash")

		ctx.JSON(http.StatusForbidden, models.ErrorResponse{Err: err, Message: "given password is not correct"})

		return
	}

	log.Println("taken user:", user)

	payload := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString(environment.Env.JwtToken)

	if err != nil {
		log.Println("an error occured while signing token: ", err)

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Err: err, Message: "cannot sign token"})

		return
	}

	log.Println("jwt token:", token)

	type responseToken struct {
		Token string
	}

	ctx.JSON(200, responseToken{Token: token})
}
