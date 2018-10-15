/*
Copyright 2018 The k8s-clusters-check Authors

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
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

var GenerateCmd = &cobra.Command {
	Hidden: true,
	Use: "generate",
	Short: "Generate commands",
	Long: "Sub command for generating documentation, man pages and documentation",
}

var GenerateManPageCmd = &cobra.Command{
	Hidden: true,
	Use: "manpage",
	Short: "Generate manpage",
	Long: "Generates a man page",
	Run: func (cmd *cobra.Command, args []string) {
		err := doc.GenManTree(RootCmd, header, "./man")
		if err != nil {
			glog.Errorf("%v", err)
		}
	},
}

var GenerateCompletionsCmd = &cobra.Command{
	Hidden: true,
	Use: "completions",
	Short: "Generates bash completions",
	Long: "Generates bash completions for the k8s-clusters-check command.",
	Run: func (cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
	GenerateCmd.AddCommand(GenerateCompletionsCmd)
	GenerateCmd.AddCommand(GenerateManPageCmd)
}