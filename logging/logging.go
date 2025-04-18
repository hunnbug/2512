package logging

import (
	"log"
)

func WriteLog(logger *log.Logger, v ...any) {

	logger.Println(v...)

}

func CheckLogError(err error) {
	if err != nil {
		log.Println("ошибка во время открытия .log: ", err)
	}
}

func TxDenied(v ...any) {
	WriteLog(ERROR, "Транзакция отменена", v)
	WriteLog(ERROR, "----------------------------------------------")
}
