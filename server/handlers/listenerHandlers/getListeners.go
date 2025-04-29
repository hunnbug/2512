package listenerHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetListeners(ctx *gin.Context) {

	const LIMIT_COUNT = 10

	logging.WriteLog(logging.DEBUG, "получен запрос на получение слушателей")

	type request struct {
		CurrentPage int
		FirstField  string
		SecondField string
		ThirdField  string
		FourthField string
		FifthField  string
		EmptyForm   bool
	}

	var _request request

	err := ctx.BindJSON(&_request)

	if err != nil {
		logging.WriteLog(logging.ERROR, "не удалось получить ответ от страницы: ", err)

	}

	var listeners []models.Listener

	//
	//Если фильтры не введены - работает обычная пагинация
	//
	if _request.EmptyForm {

		query := database.DB.Limit(LIMIT_COUNT).Offset((_request.CurrentPage - 1) * LIMIT_COUNT).Find(&listeners)

		if query.Error != nil {

			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Слушатели не найдены"})

			return

		}

	} else {

		type fieldsDTO struct {
			firstField  string
			secondField string
			thirdField  string
			fourthField string
			fifthField  string
		}

		fields := fieldsDTO{

			firstField:  _request.FirstField,
			secondField: _request.SecondField,
			thirdField:  _request.ThirdField,
			fourthField: _request.FourthField,
			fifthField:  _request.FifthField,
		}

		value := reflect.ValueOf(fields)

		var notNullFields []string

		//
		//перебор всех полей из фильтрации и отсечение пустых
		//
		for i := range value.NumField() {

			if value.Field(i).String() != "" {

				notNullFields = append(notNullFields, value.Field(i).String())

			}

		}

		var requestString string

		//
		//добавление всех строк для фильтрации в запрос
		//
		for i := 0; i < len(notNullFields)-1; i++ {

			requestString += "(firstname LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "secondname LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "middlename LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "contactphone LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "email LIKE '%" + notNullFields[i] + "%') AND "

		}

		requestString += "(firstname LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "secondname LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "middlename LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "contactphone LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "email LIKE '%" + notNullFields[len(notNullFields)-1] + "%')"

		database.DB.Find(&listeners, requestString)

	}

	ctx.JSON(http.StatusOK, listeners)
}
