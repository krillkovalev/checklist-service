package main

import (
	"api-service/application"
	"context"
	"fmt"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("error to start application: %v", err)
	}

}
