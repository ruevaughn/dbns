package repositories

import (
	"database/sql"
	"strings"

	"github.com/xm1k3/dbns/nuclei"
)

var (
	templateID string
	host       string
	name       string
	tags       string
	matchedAt  string
)

type PsqlNucleiRepository struct {
	DB    *sql.DB
	Table string
}

func (n PsqlNucleiRepository) GetSubdomains(severity string) ([]nuclei.NucleiDB, error) {
	var rows []nuclei.NucleiDB
	records, err := n.DB.Query(`select templateid,host,severity,name,tags,matched_at from "` + n.Table + `" where severity = '` + severity + `'`)
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

func (n PsqlNucleiRepository) AddSubdomain(res nuclei.NucleiResult) error {

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
	INSERT INTO ` + n.Table + ` (templateid, host, severity, name, tags, matcher_name, type, matched_at, reference, curl)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (templateid, host) DO UPDATE
	SET last_change = now();`
	_, err := n.DB.Exec(sqlStatement, res.TemplateID, res.Host, res.Info.Severity, res.Info.Name, tagsdb, res.MatcherName, res.Type, res.MatchedAt, refsdb, res.CurlCommand)
	if err != nil {
		return err
	}
	return nil
}
