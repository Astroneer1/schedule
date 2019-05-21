package common

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Erro　loading .env file")
	}
}
