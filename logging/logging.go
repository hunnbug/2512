package logging

import (
	"io"
	"log"
	"os"
)

func WriteLog(v ...any) {

	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		CheckLogError(err)
	}

	defer file.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	writer := io.MultiWriter(os.Stdout, file)

	log.SetOutput(writer)

	log.Println(v...)

	// log.SetOutput(file)

	// log.Println(v...)

	// log.SetOutput(io.Writer(os.Stdout))

	// log.Println(v...)

}

func CheckLogError(err error) {
	if err != nil {
		log.Println("ошибка во время открытия .log: ", err)
	}
}

func TxDenied(v ...any) {
	WriteLog("Транзакция отменена", v)
	WriteLog("----------------------------------------------")
}
