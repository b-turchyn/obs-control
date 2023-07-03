/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/andreykaipov/goobs/api/requests/filters"
	"github.com/b-turchyn/obs-control/client"
	"github.com/spf13/cobra"
)

// transitionCmd represents the transition command
var filterCmd = &cobra.Command{
	Use:   "filter <scene> <filter> <true|false>",
	Short: "Change source filters",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
          c := client.NewClient()

          filterEnabled := args[2] == "true"

          _, err := c.Filters.SetSourceFilterEnabled(&filters.SetSourceFilterEnabledParams{
            SourceName: args[0],
            FilterName: args[1],
            FilterEnabled: &filterEnabled,
          })

          if err != nil {
            log.Fatal(err)
          }
        },
}

func init() {
	rootCmd.AddCommand(filterCmd)
}
