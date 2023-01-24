package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	controllers "zun.com/demo/controllers"
	"zun.com/demo/middlewares"
)

var ginLambda *ginadapter.GinLambda

func main() {
	g := gin.Default()
	g.Use(middlewares.AuthenticationMiddleware)
	g.GET("/users", controllers.ListUserHandler)
	g.PUT("/users", controllers.UpdateUserHandler)
	g.POST("/users", controllers.AddUserHandler)

	env := os.Getenv("GIN_MODE")
	if env == "release" {
		ginLambda = ginadapter.New(g)
		lambda.Start(Handler)
	} else {
		g.Run(":8080")
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
