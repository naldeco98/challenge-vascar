package service

import (
	"context"
	"fmt"

	"github.com/naldeco98/challenge-vascar/internal/domain"
	"github.com/naldeco98/challenge-vascar/internal/repository"
)

type Reports interface {
	CommentReport(ctx context.Context, commentId, userId int, reason string) error
	PostReport(ctx context.Context, postId, userId int, reason string) error
}

type service struct {
	r repository.Reports
}

func NewReportService(r repository.Reports) Reports {
	return &service{r: r}
}

func (s *service) CommentReport(ctx context.Context, commentId, userId int, reason string) error {
	_, err := s.r.GetCommentById(ctx, commentId)
	if err != nil {
		return fmt.Errorf("comment not found")
	}
	report := &domain.ReportComment{
		Reason:    reason,
		UserId:    userId,
		CommentId: commentId,
	}
	_, err = s.r.CreateCommentReport(ctx, report)
	if err != nil {
		return fmt.Errorf("error creating comment report: %v", err)
	}
	return nil
}

func (s *service) PostReport(ctx context.Context, postId, userId int, reason string) error {
	_, err := s.r.GetPostById(ctx, postId)
	if err != nil {
		return fmt.Errorf("post not found")
	}
	report := &domain.ReportPost{
		Reason: reason,
		UserId: userId,
		PostId: postId,
	}
	_, err = s.r.CreatePostReport(ctx, report)
	if err != nil {
		return fmt.Errorf("error creating comment report: %v", err)
	}
	return nil
}
