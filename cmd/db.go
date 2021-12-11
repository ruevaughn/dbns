/*
Copyright © 2021 FleexSecurity

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
	"log"

	"github.com/FleexSecurity/dbns/config"
	"github.com/FleexSecurity/dbns/nuclei/repositories"
	"github.com/FleexSecurity/dbns/nuclei/services"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Retrieve data from db",
	Long:  "Retrieve data from db",
	Run: func(cmd *cobra.Command, args []string) {
		severityFlag, _ := cmd.Flags().GetString("severity")
		printFlag, _ := cmd.Flags().GetString("print")
		delimiterFlag, _ := cmd.Flags().GetString("delimiter")
		latest, _ := cmd.Flags().GetInt("latest")
		psqlDB = config.Connect()
		repository := repositories.PsqlNucleiRepository{
			DB:    psqlDB,
			Table: "nuclei",
		}
		service := services.NucleiService{
			Repository: repository,
		}
		if severityFlag == "all" {
			err := service.GetAllResults(printFlag, delimiterFlag, latest)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			err := service.GetResultsBySeverity(severityFlag, printFlag, delimiterFlag, latest)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.Flags().StringP("severity", "s", "high", "Severity flag")
	dbCmd.Flags().StringP("print", "p", "sm", "Print flags (t,h,s,n,g,m)")
	dbCmd.Flags().StringP("delimiter", "d", " - ", "Delimiter")
	dbCmd.Flags().IntP("latest", "", 0, "see latest results")
}
