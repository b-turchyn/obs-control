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

	"github.com/andreykaipov/goobs/api/requests/transitions"
	"github.com/b-turchyn/obs-control/client"
	"github.com/spf13/cobra"
)

// transitionCmd represents the transition command
var transitionCmd = &cobra.Command{
	Use:   "transition",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
          c := client.NewClient()
          defer c.Disconnect()

          _, err := c.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
            TransitionName: args[0],
          })

          if err != nil {
            log.Fatal(err)
          }
        },
}

func init() {
	rootCmd.AddCommand(transitionCmd)
}
