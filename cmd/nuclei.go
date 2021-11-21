/*
Copyright © 2021 xm1k3

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"

	"github.com/spf13/cobra"
	"github.com/xm1k3/dbns/config"
	"github.com/xm1k3/dbns/nuclei/repositories"
	"github.com/xm1k3/dbns/nuclei/services"
)

var (
	psqlDB *sql.DB
)

// nucleiCmd represents the nuclei command
var nucleiCmd = &cobra.Command{
	Use:   "nuclei",
	Short: "Nuclei Scanner command",
	Long:  "Nuclei Scanner command",
	Run: func(cmd *cobra.Command, args []string) {
		listPath, _ := cmd.Flags().GetString("list")
		url, _ := cmd.Flags().GetString("url")
		psqlDB = config.Connect()
		repository := repositories.PsqlNucleiRepository{
			DB:    psqlDB,
			Table: "nucleires",
		}
		service := services.NucleiService{
			Repository: repository,
		}
		service.AddSubdomain(url, listPath)
	},
}

func init() {
	rootCmd.AddCommand(nucleiCmd)

	nucleiCmd.Flags().StringP("list", "l", "", "subdomains list path")
	nucleiCmd.Flags().StringP("url", "u", "", "url to scan")
}
