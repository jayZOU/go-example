package main

import (
	"log"

	"github.com/jayZOU/go-example/example/blogspider/email"
	"github.com/joho/godotenv"
)

func main() {
	//加载本地配置
	loadEnv()

	err := email.Send()
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
