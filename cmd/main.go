package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	hanlder "github.com/naldeco98/challenge-vascar/cmd/handler"
	"github.com/naldeco98/challenge-vascar/internal/reports"
	"github.com/naldeco98/challenge-vascar/pkg/database"
)

const FILE = "testing.db"

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error ocurred: ", err)
		}
	}()

	db, err := database.GetDatabaseConnection("testing.db")
	
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	repo := reports.NewRepostiroy(db)
	service := reports.NewService(repo)
	handle := hanlder.NewHanlder(service)

	r.POST("comments/report", handle.ReportMessage())
	r.POST("posts/report", handle.ReportPost())

	if err := r.Run(); err != nil {
		panic(err)
	}
}
