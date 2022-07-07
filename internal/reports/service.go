package reports

import (
	"context"

	"github.com/naldeco98/challenge-vascar/internal/domain"
)

type Service interface {
	CreateMessageReport(ctx context.Context, report *domain.ReportMessage) error
	CreatePostReport(ctx context.Context, report *domain.ReportPost) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) CreateMessageReport(ctx context.Context, report *domain.ReportMessage) error {
	return nil
}

func (s *service) CreatePostReport(ctx context.Context, report *domain.ReportPost) error {
	return nil
}
