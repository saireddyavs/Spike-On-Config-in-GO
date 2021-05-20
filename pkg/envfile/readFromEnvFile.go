package envfile

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	err:=godotenv.Load("../.env");

	if err != nil {
		fmt.Println("Error loading environment")
	}

	return os.Getenv(key)
}