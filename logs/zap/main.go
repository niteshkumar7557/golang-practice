package main

import (
	"log"
	"go.uber.org/zap"
)

func main(){

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in init zap logger")
	}

	defer logger.Sync()

	logger.Info("This is an info message.")

	logger.Info("User logged in", zap.String("user", "Nitesh Kumar"), zap.String("method", "GET"))

}
