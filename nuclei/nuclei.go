package nuclei

import (
	"errors"
	"time"
)

var ErrInvalidJsonBody = errors.New("invalid json body")
var ErrInvalidUrlOrList = errors.New("you need to insert an url or list")

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
	GetSubdomains(severity string, printFlags string, delimiter string) error
	AddSubdomain(url string, list string) error
}

type Repository interface {
	GetSubdomains(severity string) ([]NucleiDB, error)
	AddSubdomain(res NucleiResult) error
}
