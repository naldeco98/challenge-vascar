package reports

import (
	"context"
	"database/sql"

	"github.com/naldeco98/challenge-vascar/internal/domain"
)

type Repository interface {
	CreateMassage(ctx context.Context, report *domain.ReportMessage) error
	CreatePost(ctx context.Context, report *domain.ReportPost) error
}

type repository struct {
	db *sql.DB
}

func NewRepostiroy(database *sql.DB) Repository {
	return &repository{db: database}
}

func (r *repository) CreateMassage(ctx context.Context, report *domain.ReportMessage) error {
	return nil
}

func (r *repository) CreatePost(ctx context.Context, report *domain.ReportPost) error {
	return nil
}
