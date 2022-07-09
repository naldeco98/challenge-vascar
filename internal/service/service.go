package service

import (
	"context"

	"github.com/naldeco98/challenge-vascar/internal/domain"
	"github.com/naldeco98/challenge-vascar/internal/repository"
)

type Reports interface {
	CommentReport(ctx context.Context, commentId, userId int, reason string) (int, error)
	PostReport(ctx context.Context, report *domain.ReportPost) (int, error)
}

type service struct {
	r repository.Reports
}

func NewService(r repository.Reports) Reports {
	return &service{r: r}
}

func (s *service) CommentReport(ctx context.Context, commentId, userId int, reason string) (int, error) {
	_, err := s.r.GetCommentById(ctx, commentId)
	if err != nil {
		return 0, err
	}
	report := &domain.ReportComment{
		Reason:    reason,
		UserId:    userId,
		CommentId: commentId,
	}
	id, err := s.r.CreateCommentReport(ctx, report)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) PostReport(ctx context.Context, report *domain.ReportPost) (int, error) {
	id, err := s.r.CreatePostReport(ctx, report)
	if err != nil {
		return 0, err
	}
	return id, nil
}
