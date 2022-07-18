package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/naldeco98/challenge-vascar/internal/service"
)

type reportComment struct {
	Reason    string `json:"reason" binding:"required"`
	UserId    int    `json:"user_id" binding:"required"`
	CommentId int    `json:"comment_id" binding:"required"`
}

type reportPost struct {
	Reason string `json:"reason" binding:"required"`
	UserId int    `json:"user_id" binding:"required"`
	PostId int    `json:"post_id" binding:"required"`
}

type Handler struct {
	service service.Reports
}

func NewHandler(service service.Reports) *Handler {
	return &Handler{service: service}
}

// ReportComment godoc
// @Summary      Create a report for a comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param data body reportComment true "report a comment"
// @Success      200  {object} string
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      422  {object} string
// @Failure      500  {object} string
// @Router       /reports/comments [post]
func (h *Handler) ReportComment() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req reportComment
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				for _, f := range verr {
					ctx.String(400, f.Field()+" is "+f.Tag())
					return
				}
			}
			switch err.(type) {
			case *json.UnmarshalTypeError:
				sliced := strings.Split(err.Error(), " ")
				ctx.String(422, sliced[8]+" is '"+sliced[11]+"' but got '"+sliced[3]+"'")
				return
			}
			ctx.String(422, "wrong request body")
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

// ReportPost godoc
// @Summary      Create a report for a post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param data body reportPost true "report a post"
// @Success      200  {object} string
// @Failure      400  {object} string
// @Failure      404  {object} string
// @Failure      422  {object} string
// @Failure      500  {object} string
// @Router       /reports/posts [post]
func (h *Handler) ReportPost() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req reportPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				for _, f := range verr {
					ctx.String(400, f.Field()+" is "+f.Tag())
					return
				}
			}
			switch err.(type) {
			case *json.UnmarshalTypeError:
				sliced := strings.Split(err.Error(), " ")
				ctx.String(422, sliced[8]+" is '"+sliced[11]+"' but got '"+sliced[3]+"'")
				return
			}
			ctx.String(422, "wrong request body")
			return
		}

		switch {
		case req.UserId < 0:
			ctx.String(400, "user_id cant be negative")
			return
		case req.PostId < 0:
			ctx.String(400, "post_id cant be negative")
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
