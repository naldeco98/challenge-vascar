package repository

import (
	"context"
	"database/sql"

	"github.com/naldeco98/challenge-vascar/internal/domain"
)

type Reports interface {
	CreateCommentReport(ctx context.Context, report *domain.ReportComment) (int, error)
	CreatePostReport(ctx context.Context, report *domain.ReportPost) (int, error)
	GetCommentById(ctx context.Context, id int) (domain.Comment, error)
	GetPostById(ctx context.Context, id int) (domain.Post, error)
}

type repository struct {
	db *sql.DB
}

func NewReportRepository(database *sql.DB) Reports {
	return &repository{db: database}
}

func (r *repository) CreateCommentReport(ctx context.Context, report *domain.ReportComment) (int, error) {
	stmt, err := r.db.Prepare("INSERT INTO report_comments (reason, user_id, comment_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(&report.Reason, &report.UserId, &report.CommentId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	*&report.Id = int(id)
	return int(id), nil
}

func (r *repository) CreatePostReport(ctx context.Context, report *domain.ReportPost) (int, error) {
	stmt, err := r.db.Prepare("INSERT INTO report_posts (reason, user_id, post_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&report.Reason, &report.UserId, &report.PostId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	*&report.Id = int(id)
	return int(id), nil
}

func (r *repository) GetCommentById(ctx context.Context, id int) (domain.Comment, error) {
	stmt, err := r.db.Prepare("SELECT id, text, creation_date FROM comments WHERE id = ?")
	if err != nil {
		return domain.Comment{}, err
	}
	defer stmt.Close()
	var result domain.Comment
	if err := stmt.QueryRow(id).Scan(&result.Id, &result.Text, &result.Created); err != nil {
		return domain.Comment{}, err
	}
	return result, nil
}

func (r *repository) GetPostById(ctx context.Context, id int) (domain.Post, error) {
	stmt, err := r.db.Prepare("SELECT id, text, creation_date FROM posts WHERE id = ?")
	if err != nil {
		return domain.Post{}, err
	}
	defer stmt.Close()
	var result domain.Post
	if err := stmt.QueryRow(id).Scan(&result.Id, &result.Text, &result.Created); err != nil {
		return domain.Post{}, err
	}
	return result, nil
}
