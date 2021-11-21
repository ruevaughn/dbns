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
	"github.com/spf13/cobra"
	"github.com/xm1k3/dbns/config"
	"github.com/xm1k3/dbns/nuclei/repositories"
	"github.com/xm1k3/dbns/nuclei/services"
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
		psqlDB = config.Connect()
		repository := repositories.PsqlNucleiRepository{
			DB:    psqlDB,
			Table: "nuclei",
		}
		service := services.NucleiService{
			Repository: repository,
		}
		service.GetSubdomains(severityFlag, printFlag, delimiterFlag)
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.Flags().StringP("severity", "s", "high", "Severity flag")
	dbCmd.Flags().StringP("print", "p", "sh", "Print flags (t,h,s,n,g,i)")
	dbCmd.Flags().StringP("delimiter", "d", " - ", "Delimiter")
}
