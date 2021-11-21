package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
	"github.com/xm1k3/dbns/nuclei"
)

type NucleiService struct {
	Repository nuclei.Repository
}

func (n NucleiService) GetSubdomains(severity string, printFlags string) error {
	rows, err := n.Repository.GetSubdomains(severity)
	if err != nil {
		return err
	}
	for _, row := range rows {
		fmt.Println(row.Severity, row.Host, row.IP)
	}
	return nil
}

func (n NucleiService) AddSubdomain(url string, list string) error {
	var args string
	envProps := viper.GetString(`dbns.nuclei.args`)
	if url != "" {
		args = `-u ` + url + ` `
	}
	if list != "" {
		args = `-l ` + list + ` `
	}
	args += envProps
	cmd := exec.Command("nuclei", strings.Split(args, " ")...)

	stderr, _ := cmd.StdoutPipe()
	cmd.Start()

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
	cmd.Wait()
	return nil
}
