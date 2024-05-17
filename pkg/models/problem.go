package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

// ProblemContentType is the default content type for a Problem response
const ProblemContentType = "application/problem+json"

// ProblemType contains a URI that identifies the problem type. This URI will,
// ideally, contain human-readable documentation for the problem when
// de-referenced.
type ProblemType string

const (
	// ProblemTypeDefault is the default problem type.
	ProblemTypeDefault = ProblemType("about:blank")
	// TODO: add more problem types that can be used to link to docs.
	// ...
)

// Problem is the RFC 7807 response body.
type Problem interface {
	Respond(w http.ResponseWriter)
	Error() string
	ProblemType() ProblemType
	ProblemTitle() string
	ProblemStatus() int
}

// StatusProblem is the RFC 7807 response body without additional fields.
type StatusProblem struct {
	Err error `json:"-"` // low-level runtime error

	// Type is a URI reference that identifies the problem type.
	Type ProblemType `json:"type"`
	// Title is a short, human-readable summary of the problem type.
	Title string `json:"title"`
	// Status is the HTTP status code generated by the origin server for this occurrence of the problem.
	Status int `json:"status"`
	// Detail is a human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`
	// Instance is a URI reference that identifies the specific occurrence of the problem.
	Instance string `json:"instance,omitempty"`
}

var _ Problem = (*StatusProblem)(nil)

func (p *StatusProblem) Error() string {
	if p.Err == nil {
		return fmt.Sprintf("[%s] %s", p.Title, p.Detail)
	}

	return fmt.Sprintf("[%s] %s - %s", p.Title, p.Err.Error(), p.Detail)
}

func (p *StatusProblem) RawError() error {
	return p.Err
}

func (p *StatusProblem) ProblemType() ProblemType {
	return p.Type
}

func (p *StatusProblem) ProblemStatus() int {
	return p.Status
}

func (p *StatusProblem) ProblemTitle() string {
	return p.Title
}

// Respond will render the problem as JSON to the provided ResponseWriter.
func (p *StatusProblem) Respond(w http.ResponseWriter) {
	RespondProblem(p, w)
}

// Respond will render the problem as JSON to the provided ResponseWriter.
func RespondProblem(problem Problem, w http.ResponseWriter) {
	// Respond
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	_ = enc.Encode(problem)

	w.Header().Set("Content-Type", ProblemContentType)
	w.WriteHeader(problem.ProblemStatus())
	_, _ = w.Write(buf.Bytes())
}

// NewStatusProblem will generate a problem for the provided HTTP status
// code. The Problem's Status field will be set to match the status argument,
// and the Title will be set to the default Go status text for that code.
func NewStatusProblem(ctx context.Context, err error, status int) *StatusProblem {
	var instance string
	reqID := middleware.GetReqID(ctx)
	if reqID != "" {
		instance = fmt.Sprintf("urn:request:%s", reqID)
	}

	// Set context canceled errors to 408.
	// Context canceled errors either happen when the client cancels the request or when a timeout happens in dependency.
	// If client cancels the request, the status code doesn't matter, because the client will not see the response.
	// If timeout happens in dependency, we want to return 408 to the client.
	if err != nil && strings.Contains(err.Error(), "context canceled") {
		status = http.StatusRequestTimeout
	}

	var detail string
	if err != nil {
		detail = err.Error()
	}

	if status == http.StatusInternalServerError {
		detail = ""
	}

	return &StatusProblem{
		Err:      err,
		Type:     ProblemTypeDefault,
		Title:    http.StatusText(status),
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}
}
