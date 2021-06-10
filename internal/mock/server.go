package mock

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const SettingsFile = ".env"

func StartServer() {
	LoadEnv()

	r := NewRouter()
	r.Run()
}

func LoadEnv() {
	// Note, any defined environment variable is used over the ones defined in .env
	if _, err := os.Stat(SettingsFile); err == nil {
		fmt.Printf("Testing setup: Loading (%s)\n", SettingsFile)
		err := godotenv.Load(SettingsFile)
		if err != nil {
			fmt.Printf("Error loading file (%s), error: %v\n", SettingsFile, err)
			return
		}
	}
}
