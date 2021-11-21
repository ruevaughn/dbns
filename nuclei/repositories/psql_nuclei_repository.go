package repositories

import (
	"database/sql"

	"github.com/xm1k3/dbns/nuclei"
)

var (
	templateID string
	host       string
	severity   string
	name       string
	tags       string
	ip         string
)

type PsqlNucleiRepository struct {
	DB    *sql.DB
	Table string
}

func (n PsqlNucleiRepository) GetSubdomains(severity string) ([]nuclei.NucleiDB, error) {
	var rows []nuclei.NucleiDB
	records, err := n.DB.Query(`select templateid,host,severity,name,tags,ip from "` + n.Table + `" where severity = '` + severity + `'`)
	if err != nil {
		return []nuclei.NucleiDB{}, err
	}
	for records.Next() {
		err = records.Scan(&templateID, &host, &severity, &name, &tags, &ip)
		if err != nil {
			return []nuclei.NucleiDB{}, err
		}
		rows = append(rows, nuclei.NucleiDB{
			TemplateID: templateID,
			Host:       host,
			Severity:   severity,
			Name:       name,
			Tags:       tags,
			IP:         ip,
		})
	}
	return rows, nil
}

func (n PsqlNucleiRepository) AddSubdomain(res nuclei.NucleiResult) error {
	sqlStatement := `
	INSERT INTO ` + n.Table + ` (templateid, host, severity, name, tags, matcher_name, type, matched, ip)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	ON CONFLICT (templateid, host) DO UPDATE
	SET last_change = now();`
	_, err := n.DB.Exec(sqlStatement, res.TemplateID, res.Host, res.Info.Severity, res.Info.Name, res.Info.Tags, res.MatcherName, res.Type, res.Matched, res.IP)
	if err != nil {
		return err
	}
	return nil
}
