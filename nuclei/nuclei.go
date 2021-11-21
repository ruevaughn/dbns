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
	IP         string
}

type NucleiResult struct {
	TemplateID  string    `json:"templateID"`
	Info        Info      `json:"info"`
	MatcherName string    `json:"matcher_name"`
	Type        string    `json:"type"`
	Host        string    `json:"host"`
	Matched     string    `json:"matched"`
	IP          string    `json:"ip"`
	Timestamp   time.Time `json:"timestamp"`
}

type Info struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

type Service interface {
	GetSubdomains(severity string, printFlags string, delimiter string) error
	AddSubdomain(url string, list string) error
}

type Repository interface {
	GetSubdomains(severity string) ([]NucleiDB, error)
	AddSubdomain(res NucleiResult) error
}
