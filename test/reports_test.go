package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/challenge-vascar/cmd/handler"
	"github.com/naldeco98/challenge-vascar/internal/repository"
	"github.com/naldeco98/challenge-vascar/internal/service"
	"github.com/naldeco98/challenge-vascar/pkg/storage"
)

type mockDB struct {
	pathFile     string
	tempPathFile string
	db           *sql.DB
}

func NewMockDB(pathFile string) *mockDB {
	return &mockDB{
		pathFile:     pathFile,
		tempPathFile: "temp.db",
	}
}

func (m *mockDB) Begin() error {
	data, err := os.ReadFile(m.pathFile)
	if err != nil {
		return err
	}
	if err := os.WriteFile(m.tempPathFile, data, 0644); err != nil {
		return err
	}
	m.db, err = storage.GetDatabaseConnection(m.tempPathFile)
	if err != nil {
		return err
	}
	return nil
}

func (m mockDB) RollBack() error {
	if err := os.Remove(m.tempPathFile); err != nil {
		return err
	}
	return nil
}

func createServer(db *sql.DB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	repo := repository.NewReportRepository(db)
	service := service.NewReportService(repo)
	handle := handler.NewHandler(service)

	r := gin.New()

	reports := r.Group("/reports")
	{
		reports.POST("comments", handle.ReportComment())
		reports.POST("posts", handle.ReportPost())
	}
	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestReportCommentOk(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	mockRequest := `{"reason": "bad attitude","user_id": 1,"comment_id": 1}`
	// Act
	req, rr := createRequestTest(http.MethodPost, "/reports/comments", mockRequest)
	r.ServeHTTP(rr, req)
	// Assert
	if http.StatusCreated != rr.Code {
		t.Errorf("status code must be %d, but got %d", http.StatusCreated, rr.Code)
	}
	if string(rr.Body.Bytes()) != "report_created" {
		t.Errorf("response body must be %s, but got %s", "report_created", string(rr.Body.Bytes()))
	}
}

func TestReportPostOk(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	mockRequest := `{"reason": "bad attitude","user_id": 1,"post_id": 1}`
	// Act
	req, rr := createRequestTest(http.MethodPost, "/reports/posts", mockRequest)
	r.ServeHTTP(rr, req)
	// Assert
	if http.StatusCreated != rr.Code {
		t.Errorf("status code must be %d, but got %d", http.StatusCreated, rr.Code)
	}
	if string(rr.Body.Bytes()) != "report_created" {
		t.Errorf("response body must be %s, but got %s", "report_created", string(rr.Body.Bytes()))
	}
}

func TestReportCommentErrorBadRequest(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	cases := map[string]string{
		`{"reason": "bad attitude","user_id": 1}`:                   `CommentId is required`,       // Missing "comment_id"
		`{"reason": "bad attitude","comment_id": 1}`:                `UserId is required`,          // Missing "user_id"
		`{"user_id": 1,"comment_id": 1}`:                            `Reason is required`,          // Missing "reason"
		`{"reason": "bad attitude","user_id": 0,"comment_id": 1}`:   `UserId is required`,          // "user_id" zero
		`{"reason": "bad attitude","user_id": -10,"comment_id": 1}`: `user_id cant be negative`,    // "user_id" negative
		`{"reason": "bad attitude","user_id": 1,"comment_id": 0}`:   `CommentId is required`,       // "comment_id" zero
		`{"reason": "bad attitude","user_id": 1,"comment_id": -10}`: `comment_id cant be negative`, // "comment_id" negative
	}
	// Act
	for send, responseExpected := range cases {
		req, rr := createRequestTest(http.MethodPost, "/reports/comments", send)
		r.ServeHTTP(rr, req)
		// Assert
		if http.StatusBadRequest != rr.Code {
			t.Errorf("status code must be '%d', but got '%d'", http.StatusBadRequest, rr.Code)
		}
		if string(rr.Body.Bytes()) != responseExpected {
			t.Errorf("response body must be '%s', but got '%s'", responseExpected, string(rr.Body.Bytes()))
		}
	}
}

func TestReportPostErrorBadRequest(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	cases := map[string]string{
		`{"reason": "bad attitude","user_id": 1}`:                `PostId is required`,       // Missing "post_id"
		`{"reason": "bad attitude","post_id": 1}`:                `UserId is required`,       // Missing "user_id"
		`{"user_id": 1,"post_id": 1}`:                            `Reason is required`,       // Missing "reason"
		`{"reason": "bad attitude","user_id": 0,"post_id": 1}`:   `UserId is required`,       // "user_id" zero
		`{"reason": "bad attitude","user_id": -10,"post_id": 1}`: `user_id cant be negative`, // "user_id" negative
		`{"reason": "bad attitude","user_id": 1,"post_id": 0}`:   `PostId is required`,       // "post_id" zero
		`{"reason": "bad attitude","user_id": 1,"post_id": -10}`: `post_id cant be negative`, // "post_id" negative
	}
	// Act
	for send, responseExpected := range cases {
		req, rr := createRequestTest(http.MethodPost, "/reports/posts", send)
		r.ServeHTTP(rr, req)
		// Assert
		if http.StatusBadRequest != rr.Code {
			t.Errorf("status code must be '%d', but got '%d'", http.StatusBadRequest, rr.Code)
		}
		if string(rr.Body.Bytes()) != responseExpected {
			t.Errorf("response body must be '%s', but got '%s'", responseExpected, string(rr.Body.Bytes()))
		}
	}
}

func TestReportCommentErrorUnprocessableEntity(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	cases := map[string]string{
		`{"reason": 5,"user_id": 1,"comment_id": 1}`:                `request.reason is 'string' but got 'number'`,  // reason is number
		`{"reason": {},"user_id": 1,"comment_id": 1}`:               `request.reason is 'string' but got 'object'`,  // reason is object
		`{"reason": "bad attitude","user_id": "1","comment_id": 1}`: `request.user_id is 'int' but got 'string'`,    // user_id is string
		`{"reason": "bad attitude","user_id": 1,"comment_id": "1"}`: `request.comment_id is 'int' but got 'string'`, // comment_id is string
		``: `wrong request body`, // Empty body
	}
	// Act
	for send, responseExpected := range cases {
		req, rr := createRequestTest(http.MethodPost, "/reports/comments", send)
		r.ServeHTTP(rr, req)
		// Assert
		if http.StatusUnprocessableEntity != rr.Code {
			t.Errorf("status code must be '%d', but got '%d'", http.StatusUnprocessableEntity, rr.Code)
		}
		if string(rr.Body.Bytes()) != responseExpected {
			fmt.Println(send, responseExpected)
			t.Errorf("response body must be '%s', but got '%s'", responseExpected, string(rr.Body.Bytes()))
		}
	}
}

func TestReportPostErrorUnprocessableEntity(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	cases := map[string]string{
		`{"reason": 5,"user_id": 1,"post_id": 1}`:                `request.reason is 'string' but got 'number'`, // reason is number
		`{"reason": {},"user_id": 1,"post_id": 1}`:               `request.reason is 'string' but got 'object'`, // reason is object
		`{"reason": "bad attitude","user_id": "1","post_id": 1}`: `request.user_id is 'int' but got 'string'`,   // user_id is string
		`{"reason": "bad attitude","user_id": 1,"post_id": "1"}`: `request.post_id is 'int' but got 'string'`,   // post_id is string
		``: `wrong request body`, // Empty body
	}
	// Act
	for send, responseExpected := range cases {
		req, rr := createRequestTest(http.MethodPost, "/reports/posts", send)
		r.ServeHTTP(rr, req)
		fmt.Println(send, responseExpected, rr.Code)
		// Assert
		if http.StatusUnprocessableEntity != rr.Code {
			t.Errorf("status code must be '%d', but got '%d'", http.StatusUnprocessableEntity, rr.Code)
		}
		if string(rr.Body.Bytes()) != responseExpected {
			t.Errorf("response body must be '%s', but got '%s'", responseExpected, string(rr.Body.Bytes()))
		}
	}
}

func TestReportCommentInternalServerError(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	cases := map[string]string{
		`{"reason": "bad attitude","user_id": 1,"comment_id": 1001}`: `error creating report: comment not found`, // comment not found
	}
	// Act
	for send, responseExpected := range cases {
		req, rr := createRequestTest(http.MethodPost, "/reports/comments", send)
		r.ServeHTTP(rr, req)
		// Assert
		if http.StatusInternalServerError != rr.Code {
			t.Errorf("status code must be '%d', but got '%d'", http.StatusInternalServerError, rr.Code)
		}
		if string(rr.Body.Bytes()) != responseExpected {
			t.Errorf("response body must be '%s', but got '%s'", responseExpected, string(rr.Body.Bytes()))
		}
	}
}

func TestReportPostInternalServerError(t *testing.T) {
	// Arange
	mockDatabase := NewMockDB("../my_database.db")
	if err := mockDatabase.Begin(); err != nil {
		t.Fatalf("fail connection to database: %v", err)
	}
	defer mockDatabase.RollBack()
	r := createServer(mockDatabase.db)
	cases := map[string]string{
		`{"reason": "bad attitude","user_id": 1,"post_id": 1001}`: `error creating report: post not found`, // post not found
	}
	// Act
	for send, responseExpected := range cases {
		req, rr := createRequestTest(http.MethodPost, "/reports/posts", send)
		r.ServeHTTP(rr, req)
		// Assert
		if http.StatusInternalServerError != rr.Code {
			t.Errorf("status code must be '%d', but got '%d'", http.StatusInternalServerError, rr.Code)
		}
		if string(rr.Body.Bytes()) != responseExpected {
			t.Errorf("response body must be '%s', but got '%s'", responseExpected, string(rr.Body.Bytes()))
		}
	}
}
