package hanlder

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/challenge-vascar/internal/domain"
	"github.com/naldeco98/challenge-vascar/internal/service"
)

type Hanlder struct {
	service service.Reports
}

func NewHanlder(service service.Reports) *Hanlder {
	return &Hanlder{service: service}
}

func (h *Hanlder) ReportComment() gin.HandlerFunc {

	type request struct {
		Reason    string `json:"reason" binding:"required" validate:""`
		UserId    int    `json:"user_id" binding:"required"`
		CommentId int    `json:"comment_id" binding:"required"`
	}

	return func(ctx *gin.Context) {
		var req request
		if err := validate(ctx, &req); err != nil {
			return
		}
		// if err := ctx.BindJSON(&req); err != nil {
		// 	ctx.JSON(400, gin.H{"error": err.Error()})
		// 	return
		// }
		_, err := h.service.CommentReport(ctx, req.CommentId, req.UserId, req.Reason)
		if err != nil {
			err = fmt.Errorf("error creating report: %w", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.String(201, "report_created")
	}
}

func (h *Hanlder) ReportPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.ReportPost
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		_, err := h.service.PostReport(ctx, &request)
		if err != nil {
			err = fmt.Errorf("error creating report: %w", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.String(201, "report_created")
	}
}

func validate(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	// TODO Validate fields
	return nil
}
