package initializer

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[ PROGRAM STATUS ] : Loaded Environmental Variables Successfully")
	}
}
