package hanlder

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/challenge-vascar/internal/domain"
	"github.com/naldeco98/challenge-vascar/internal/reports"
)

type Hanlder struct {
	service reports.Service
}

func NewHanlder(service reports.Service) *Hanlder {
	return &Hanlder{service: service}
}

func (h *Hanlder) ReportMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.ReportMessage
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"error": "error binding JSON"})
			return
		}
		err := h.service.CreateMessageReport(ctx, &request)
		if err != nil {
			err = fmt.Errorf("error creating report: %w", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, fmt.Sprint("report created successfully"))
	}
}

func (h *Hanlder) ReportPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.ReportPost
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"error": "error binding JSON"})
			return
		}
		err := h.service.CreatePostReport(ctx, &request)
		if err != nil {
			err = fmt.Errorf("error creating report: %w", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, fmt.Sprint("report created successfully"))
	}
}
