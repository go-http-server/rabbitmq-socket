package main

import (
	"fmt"
	"log"

	"github.com/go-http-server/rabbitmq-socket/utils"
)

func main() {
	envVars, err := utils.LoadEnvironmentVariables("./")
	if err != nil {
		log.Fatal("cannot load environment variables: ", err)
	}
	fmt.Println(envVars.HTTP_SERVER_ADDRESS)
}
