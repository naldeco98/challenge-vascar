package hanlder

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
		Reason    string `json:"reason" binding:"required"`
		UserId    int    `json:"user_id" binding:"required"`
		CommentId int    `json:"comment_id" binding:"required"`
	}

	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.String(400, err.Error())
			return
		}

		switch {
		case req.UserId < 0:
			ctx.String(400, "user_id cant be negative")
			return
		case req.CommentId < 0:
			ctx.String(400, "comment_id cant be negative")
			return
		case req.Reason == "":
			ctx.String(400, "the reason cant be empty")
			return
		}

		err := h.service.CommentReport(ctx, req.CommentId, req.UserId, req.Reason)
		if err != nil {
			err = fmt.Errorf("error creating report: %w", err)
			ctx.String(500, err.Error())
			return
		}
		ctx.String(201, "report_created")
	}
}

func (h *Hanlder) ReportPost() gin.HandlerFunc {

	type request struct {
		Reason string `json:"reason" binding:"required"`
		UserId int    `json:"user_id" binding:"required"`
		PostId int    `json:"post_id" binding:"required"`
	}

	return func(ctx *gin.Context) {
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.String(400, err.Error())
			return
		}

		switch {
		case req.UserId < 0:
			ctx.String(400, "user_id cant be negative")
			return
		case req.PostId < 0:
			ctx.String(400, "comment_id cant be negative")
			return
		case req.Reason == "":
			ctx.String(400, "the reason cant be empty")
			return
		}

		err := h.service.PostReport(ctx, req.PostId, req.UserId, req.Reason)
		if err != nil {
			err = fmt.Errorf("error creating report: %w", err)
			ctx.String(500, err.Error())
			return
		}
		ctx.String(201, "report_created")
	}
}
