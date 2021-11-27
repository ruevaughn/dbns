package nuclei

import (
	"errors"
	"time"
)

var ErrInvalidJsonBody = errors.New("invalid json body")
var ErrInvalidUrlOrList = errors.New("you need to insert an url or list")
var ErrInvalidSeverity = errors.New("invalid severity")

type NucleiDB struct {
	TemplateID string
	Host       string
	Severity   string
	Name       string
	Tags       string
	MatchedAt  string
}

type NucleiResult struct {
	TemplateID  string    `json:"template-id"`
	Info        Info      `json:"info"`
	MatcherName string    `json:"matcher-name"`
	Type        string    `json:"type"`
	Host        string    `json:"host"`
	MatchedAt   string    `json:"matched-at"`
	Timestamp   time.Time `json:"timestamp"`
	CurlCommand string    `json:"curl-command"`
}

type Info struct {
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Reference   []string `json:"reference"`
	Severity    string   `json:"severity"`
}

type Service interface {
	GetResultsBySeverity(severity string, printFlags string, delimiter string, latest int) error
	GetAllResults(severity string, printFlags string, delimiter string) error
	Scan(url string, list string, info bool) error
}

type Repository interface {
	GetResultsBySeverity(severity string, latest int) ([]NucleiDB, error)
	GetAllResults(latest int) ([]NucleiDB, error)
	AddSubdomain(res NucleiResult) error
}
