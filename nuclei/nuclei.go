package nuclei

import "time"

// type NucleiResult struct {
// }

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
	GetSubdomains() error
	AddSubdomain(url string, list string) error
}

type Repository interface {
	GetSubdomains() error
	AddSubdomain() error
}
