package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func HandlePanic() int {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic :: ", r)
	}
	return 0
}

func SetGlobals() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
