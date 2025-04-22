package logging

import (
	"io"
	"log"
	"os"
)

var DEBUG *log.Logger
var ERROR *log.Logger

func Init() {

	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Println("ошибка во время открытия .log файла: ", err)
	}

	writer := io.MultiWriter(os.Stdout, file)

	DEBUG = log.New(writer, "DEBUG\t", log.LstdFlags)

	ERROR = log.New(writer, "ERROR\t", log.LstdFlags)

}
