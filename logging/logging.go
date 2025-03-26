package logging

import (
	"io"
	"log"
	"os"
)

func WriteLog(v ...any) error {

	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	log.SetOutput(file)

	log.Println(v...)

	log.SetOutput(io.Writer(os.Stdout))

	log.Println(v...)

	return nil

}

func CheckLogError(err error) {
	if err != nil {
		log.Println("ошибка во время открытия .log: ", err)
	}
}
