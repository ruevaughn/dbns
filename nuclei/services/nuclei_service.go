package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/FleexSecurity/dbns/nuclei"
	output "github.com/projectdiscovery/nuclei/v2/pkg/output"
	"github.com/spf13/viper"
)

type NucleiService struct {
	Repository nuclei.Repository
}

func (n NucleiService) GetResultsBySeverity(severity string, printFlags string, delimiter string, latest int) error {
	rows, err := n.Repository.GetResultsBySeverity(severity, latest)
	if err != nil {
		return err
	}
	for _, row := range rows {
		out := n.FilterOutput(row, printFlags, delimiter)
		fmt.Println(strings.TrimSuffix(out, delimiter))
	}
	return nil
}

func (n NucleiService) GetAllResults(printFlags string, delimiter string, latest int) error {
	rows, err := n.Repository.GetAllResults(latest)
	if err != nil {
		return err
	}
	for _, row := range rows {
		out := n.FilterOutput(row, printFlags, delimiter)
		fmt.Println(strings.TrimSuffix(out, delimiter))
	}
	return nil
}

func (n NucleiService) Scan(url string, list string, info bool) error {
	var args string
	envProps := viper.GetString(`dbns.nuclei.args`)
	if url != "" {
		args = `-u ` + url + ` `
	}
	if list != "" {
		args = `-l ` + list + ` `
	}
	if !info {
		args += `-exclude-severity info `
	}
	args += envProps
	cmd := exec.Command("nuclei", strings.Split(args, " ")...)

	stderr, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		m := scanner.Text()
		var nujson nuclei.NucleiResult
		if err := json.Unmarshal([]byte(m), &nujson); err != nil {
			log.Println("ERR:", nuclei.ErrInvalidJsonBody)
		}
		err := n.Repository.AddSubdomain(nujson)
		if err != nil {
			return err
		}
	}
	err = cmd.Wait()
	if err != nil {
		return nuclei.ErrGenericError
	}
	return nil
}

func (n NucleiService) ScanTest(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)

	stderr, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		m := scanner.Text()
		var nujson output.ResultEvent
		if err := json.Unmarshal([]byte(m), &nujson); err != nil {
			// log.Println("ERR:", nuclei.ErrInvalidJsonBody)
			return nuclei.ErrInvalidJsonBody
		}

		fmt.Println("JSON:", nujson)
	}
	err = cmd.Wait()
	if err != nil {
		return nuclei.ErrGenericError
	}
	return nil
}

func (n NucleiService) FilterOutput(row nuclei.NucleiDB, printFlags string, delimiter string) string {
	outrow := ""
	for _, char := range printFlags {
		if char == 'h' {
			outrow += row.Host + delimiter
		} else if char == 't' {
			outrow += row.TemplateID + delimiter
		} else if char == 's' {
			outrow += row.Severity + delimiter
		} else if char == 'n' {
			outrow += row.Name + delimiter
		} else if char == 'm' {
			outrow += row.MatchedAt + delimiter
		} else if char == 'g' {
			outrow += row.Tags + delimiter
		}
	}
	return outrow
}
