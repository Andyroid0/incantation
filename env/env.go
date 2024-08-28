package env

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		str := fmt.Sprintf(
			"🚨 Error loading .env file.\n\t%v\n\tℹ️ Disregard if in production.",
			err)
		log.Print(str)
	}
}
