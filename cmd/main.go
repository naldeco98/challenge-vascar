package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	hanlder "github.com/naldeco98/challenge-vascar/cmd/handler"
	"github.com/naldeco98/challenge-vascar/internal/repository"
	"github.com/naldeco98/challenge-vascar/internal/service"
	"github.com/naldeco98/challenge-vascar/pkg/storage"
)

const FILE = "testing.db"

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error ocurred: ", err)
		}
	}()

	db, err := storage.GetDatabaseConnection("testing.db")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	repo := repository.NewRepostiroy(db)
	service := service.NewService(repo)
	handle := hanlder.NewHanlder(service)

	reports := r.Group("/reports")
	{
		reports.POST("comments", handle.ReportComment())
		reports.POST("posts", handle.ReportPost())
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
