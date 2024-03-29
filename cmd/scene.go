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

	"github.com/andreykaipov/goobs/api/requests/scenes"
        "github.com/b-turchyn/obs-control/client"
	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var sceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "Switch to scene",
	Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
          c := client.NewClient()

          _, err := c.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
            SceneName: args[0],
          })
          if err != nil {
            log.Fatal(err)
          }
        },
}

func init() {
	rootCmd.AddCommand(sceneCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sceneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sceneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
