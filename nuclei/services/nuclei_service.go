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

func (n NucleiService) GetSubdomains() error {
	fmt.Println("Service: get subdomain")
	n.Repository.GetSubdomains()
	return nil
}

// http://testphp.vulnweb.com
func (n NucleiService) AddSubdomain(url string, list string) error {
	var args string
	fmt.Println("Service: Add subdomain")
	envProps := viper.GetString(`dbns.nuclei.args`)
	if url != "" {
		args = `-u ` + url + ` `
	}
	args += envProps
	cmd := exec.Command("nuclei", strings.Split(args, " ")...)

	stderr, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		m := scanner.Text()

		var reqBody nuclei.NucleiResult
		if err := json.Unmarshal([]byte(m), &reqBody); err != nil {
			log.Println("ERR", err)
		}
		fmt.Println("NUCLEI | ", reqBody.Info.Severity, " | ", reqBody.Host)
	}
	cmd.Wait()
	n.Repository.AddSubdomain()
	return nil
}
