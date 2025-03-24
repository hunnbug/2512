package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	DbConnectionString string
	JwtToken           []byte
}

var Env environment

func InitEnv() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	Env = environment{
		DbConnectionString: os.Getenv("DB_STRING"),
		JwtToken:           []byte(os.Getenv("JWT_KEY")),
	}

}
