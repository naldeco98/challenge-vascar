package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/challenge-vascar/cmd/handler"
	"github.com/naldeco98/challenge-vascar/docs"
	"github.com/naldeco98/challenge-vascar/internal/repository"
	"github.com/naldeco98/challenge-vascar/internal/service"
	"github.com/naldeco98/challenge-vascar/pkg/storage"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Vascar Challenge
// @version 1.0

// @description This a documentation for vescar challenge the objective is to provide an API to report comments and posts.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error ocurred: ", err)
		}
	}()

	db, err := storage.GetDatabaseConnection("my_database.db")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	docs.SwaggerInfo.Host = "localhost:8080"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	repo := repository.NewReportRepository(db)
	service := service.NewReportService(repo)
	handle := handler.NewHandler(service)

	reports := r.Group("/reports")
	{
		reports.POST("comments", handle.ReportComment())
		reports.POST("posts", handle.ReportPost())
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
