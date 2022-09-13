package repositories

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/FleexSecurity/dbns/nuclei"
)

var (
	templateID string
	host       string
	severity   string
	name       string
	tags       string
	matchedAt  string
)

type PsqlNucleiRepository struct {
	DB    *sql.DB
	Table string
}

func (p PsqlNucleiRepository) GetResultsBySeverity(sev string, latest int) ([]nuclei.NucleiDB, error) {
	availableSeverity := []string{"info", "low", "medium", "high", "critical"}
	if !p.stringInSlice(sev, availableSeverity) {
		return []nuclei.NucleiDB{}, nuclei.ErrInvalidSeverity
	}
	var rows []nuclei.NucleiDB
	query := `select templateid,host,severity,name,tags,matched_at from "` + p.Table + `" where severity = '` + sev + `'`
	if latest != 0 {
		query += ` and last_change > now() - interval '` + strconv.Itoa(latest) + ` day';`
	}
	records, err := p.DB.Query(query)
	if err != nil {
		return []nuclei.NucleiDB{}, err
	}
	for records.Next() {
		err = records.Scan(&templateID, &host, &sev, &name, &tags, &matchedAt)
		if err != nil {
			return []nuclei.NucleiDB{}, err
		}
		rows = append(rows, nuclei.NucleiDB{
			TemplateID: templateID,
			Host:       host,
			Severity:   sev,
			Name:       name,
			Tags:       tags,
			MatchedAt:  matchedAt,
		})
	}
	return rows, nil
}

func (p PsqlNucleiRepository) GetAllResults(latest int) ([]nuclei.NucleiDB, error) {
	var rows []nuclei.NucleiDB
	query := `select templateid,host,severity,name,tags,matched_at from "` + p.Table + `"`
	if latest != 0 {
		query += ` where last_change > now() - interval '` + strconv.Itoa(latest) + ` day';`
	}
	records, err := p.DB.Query(query)
	if err != nil {
		return []nuclei.NucleiDB{}, err
	}
	for records.Next() {
		err = records.Scan(&templateID, &host, &severity, &name, &tags, &matchedAt)
		if err != nil {
			return []nuclei.NucleiDB{}, err
		}
		rows = append(rows, nuclei.NucleiDB{
			TemplateID: templateID,
			Host:       host,
			Severity:   severity,
			Name:       name,
			Tags:       tags,
			MatchedAt:  matchedAt,
		})
	}
	return rows, nil
}

func (p PsqlNucleiRepository) AddSubdomain(res nuclei.NucleiResult) error {

	tagsdb := ""
	for _, author := range res.Info.Tags {
		tagsdb += author + `,`
	}
	tagsdb = strings.TrimSuffix(tagsdb, `,`)

	refsdb := ""
	for _, refs := range res.Info.Reference {
		refsdb += refs + `,`
	}
	refsdb = strings.TrimSuffix(refsdb, `,`)

	sqlStatement := `
	INSERT INTO ` + p.Table + ` (templateid, host, severity, name, tags, matcher_name, type, matched_at, reference)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	ON CONFLICT (templateid, host) DO UPDATE
	SET last_change = now();`
	_, err := p.DB.Exec(sqlStatement, res.TemplateID, res.Host, res.Info.Severity, res.Info.Name, tagsdb, res.MatcherName, res.Type, res.MatchedAt, refsdb)
	if err != nil {
		return err
	}
	return nil
}

func (p PsqlNucleiRepository) stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
