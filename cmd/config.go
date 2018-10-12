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
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var InitConfigCmd = &cobra.Command{
	Use: "init-config",
	Short: "Initialize config",
	Long: "Initialize an empty config structure",
	Run: func (cmd *cobra.Command, args []string) {
		if ba, err := json.Marshal(Config{}); err != nil {
			glog.Errorf("%v", err)
		}
	},
}

var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "Config manipulates the current configuration",
	Long: "Manipulates the ~/.config/k8scc.json configuration file.",
	Run: func (cmd *cobra.Command, args []string) {
		fmt.Println("config")
	},
}



func init() {
	RootCmd.AddCommand(InitConfigCmd)
	RootCmd.AddCommand(ConfigCmd)
}
