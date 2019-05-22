package common

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro　loading .env file")
	}
}
