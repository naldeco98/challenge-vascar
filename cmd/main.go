package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/challenge-vascar/cmd/handler"
	"github.com/naldeco98/challenge-vascar/internal/repository"
	"github.com/naldeco98/challenge-vascar/internal/service"
	"github.com/naldeco98/challenge-vascar/pkg/storage"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error ocurred: ", err)
		}
	}()

	db, err := storage.GetDatabaseConnection("/home/naldeco/github/challenge-vascar/my_database.db")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

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
