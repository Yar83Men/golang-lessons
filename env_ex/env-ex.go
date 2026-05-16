package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//envVar := os.Getenv("ENV_VAR")
	//fmt.Println(envVar)

	_, err := os.Stat(".env")
	log.Printf("Stat env: %v", err)

	wd, err := os.Getwd()
	fmt.Println("WD:", wd, "err:", err)

	if err := godotenv.Load("D:\\go_lessons\\.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	if _, err := godotenv.Read("VAR_ENV"); err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}
	envVar, _ := godotenv.Read("VAR_ENV")
	fmt.Println(envVar)

	//listEnv := os.Environ()
	//fmt.Println(listEnv)
	//
	//for _, v := range listEnv {
	//	fmt.Println(v)
	//}
}
