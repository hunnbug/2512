package environment

import (
	"main/logging"
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	DbConnectionString string
	JwtToken           []byte
}

type s3Environment struct {
	PublicKey  string
	PrivateKey string
	Bucket     string
}

var Env environment

func InitEnv() {

	err := godotenv.Load()

	if err != nil {
		logging.WriteLog(logging.DEBUG, err)
	}

	Env = environment{
		DbConnectionString: os.Getenv("DB_STRING"),
		JwtToken:           []byte(os.Getenv("JWT_KEY")),
	}

}

var S3 s3Environment

func InitS3Enviroment() {
	err := godotenv.Load()

	if err != nil {
		logging.WriteLog(logging.DEBUG, err)
	}

	S3 = s3Environment{
		PublicKey:  os.Getenv("S3_PUBLIC"),
		PrivateKey: os.Getenv("S3_PRIVATE"),
		Bucket:     os.Getenv("BUCKET"),
	}
}
