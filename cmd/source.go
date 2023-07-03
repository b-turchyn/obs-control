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

	"github.com/andreykaipov/goobs/api/requests/sceneitems"
	"github.com/b-turchyn/obs-control/client"
	"github.com/spf13/cobra"
)

// transitionCmd represents the transition command
var sourceCmd = &cobra.Command{
	Use:   "source <scene> <source> <true|false>",
	Short: "Enable or disable a source",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
          c := client.NewClient()

          enabled := args[2] == "true"

          itemId, err := c.SceneItems.GetSceneItemId(&sceneitems.GetSceneItemIdParams{
            SceneName: args[0],
            SourceName: args[1],
          })
          if err != nil {
            log.Fatal(err)
          }

          _, err = c.SceneItems.SetSceneItemEnabled(&sceneitems.SetSceneItemEnabledParams{
            SceneName: args[0],
            SceneItemId: itemId.SceneItemId,
            SceneItemEnabled: &enabled,
          })

          if err != nil {
            log.Fatal(err)
          }
        },
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}
