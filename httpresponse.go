package gopeach

import (
	"strconv"
	"time"
)

// Error represents an http error.
type HTTPResponse struct {
	Timestamp int64
	Status    int
	Error     string
	Message   string
	Path      string
}

// NewError returns a new Error.
func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{}
}

// Created returns a new HTTPResponse with status 201.
func (e *HTTPResponse) Created(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 201
	e.Error = ""
	e.Message = message
	e.Path = path
	return e
}

// Accepted returns a new HTTPResponse with status 202.
func (e *HTTPResponse) Accepted(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 202
	e.Error = ""
	e.Message = message
	e.Path = path
	return e
}

// BadRequest returns a new HTTPResponse with status 400.
func (e *HTTPResponse) BadRequest(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 400
	e.Error = "bad request"
	e.Message = message
	e.Path = path
	return e
}

// Unauthorized returns a new HTTPResponse with status 401.
func (e *HTTPResponse) Unauthorized(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 401
	e.Error = "unauthorized"
	e.Message = message
	e.Path = path
	return e
}

// Forbidden returns a new HTTPResponse with status 403.
func (e *HTTPResponse) Forbidden(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 403
	e.Error = "forbidden"
	e.Message = message
	e.Path = path
	return e
}

// NotFound returns a new HTTPResponse with status 404.
func (e *HTTPResponse) NotFound(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 404
	e.Error = "not found"
	e.Message = message
	e.Path = path
	return e
}

// Conflict returns a new HTTPResponse with status 409.
func (e *HTTPResponse) Conflict(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 409
	e.Error = "conflict"
	e.Message = message
	e.Path = path
	return e
}

// InternalServerError returns a new HTTPResponse with status 500.
func (e *HTTPResponse) InternalServerError(message, path string) *HTTPResponse {
	e.Timestamp = time.Now().Unix()
	e.Status = 500
	e.Error = "internal server error"
	e.Message = message
	e.Path = path
	return e
}

// ToJson returns the HTTPResponse as a JSON string.
func (e *HTTPResponse) ToJson() string {
	// FormatInt is 1.3 times faster than Itoa
	// FormatInt + int64() is 1.1 times faster than Itoa
	return `{
		"timestamp": ` + strconv.FormatInt(e.Timestamp, 10) + `,
		"status": ` + strconv.FormatInt(int64(e.Status), 10) + `,
		"error": "` + e.Error + `",
		"message": "` + e.Message + `",
		"path": "` + e.Path + `"
	}`
}
