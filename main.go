package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	a := App{}
	godotenv.Load()
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	fmt.Println("os.Getenv(\"APP_DB_NAME\")")
	fmt.Println(os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
