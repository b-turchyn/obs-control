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
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

type step struct {
	Name    string            `yaml:"name"`
	Options map[string]string `yaml:"options"`
}

// playbookCmd represents the playbook command
var playbookCmd = &cobra.Command{
	Use:   "playbook",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		steps := getSteps(args[0])
		runSteps(steps)
	},
}

func init() {
	rootCmd.AddCommand(playbookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playbookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playbookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getSteps(filename string) []step {
	var result []step

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}
	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		log.Fatal("Unable to parse file: ", err)
	}

	return result
}

func runSteps(steps []step) {
	for i := 0; i < len(steps); i++ {
		steps[i].runStep()
	}
}

func (s *step) runStep() {
	cmd, _, err := rootCmd.Find([]string{s.Name})

	if err != nil {
		log.Fatal("Unknown step: ", s.Name)
	}

	if s.Name == "scene" {
		cmd.Run(cmd, []string{s.Options["scene"]})
	} else if s.Name == "sleep" {
		cmd.Run(cmd, []string{s.Options["duration"]})
	} else if s.Name == "transition" {
		cmd.Run(cmd, []string{s.Options["name"]})
	} else if s.Name == "filter" {
		cmd.Run(cmd, []string{s.Options["scene"], s.Options["filter"], s.Options["enabled"]})
	} else if s.Name == "source" {
		cmd.Run(cmd, []string{s.Options["scene"], s.Options["source"], s.Options["enabled"]})
	}

}
