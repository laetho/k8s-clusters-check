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


var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "Config manipulates the current configuration",
	Long: "Manipulates the ~/.config/k8scc.json configuration file.",
}

var ConfigInitCmd = &cobra.Command{
	Use: "init-config",
	Short: "Initialize config",
	Long: "Initialize an empty config structure",
	Run: func (cmd *cobra.Command, args []string) {
		conf := Config{}
		ba, err := json.Marshal(conf)
		if err != nil {
			glog.Errorf("%v", err)
		}
		fmt.Println(string(ba))
	},
}

var ConfigListCmd = &cobra.Command{
	Use: "list",
	Short: "List current configuration",
	Long: "Manipulates the ~/.config/k8scc.json configuration file.",
	Run: func (cmd *cobra.Command, args []string) {
		Conf = initConfig()
		configList()
	},
}

func init() {
	RootCmd.AddCommand(ConfigCmd)
	ConfigCmd.AddCommand(ConfigInitCmd)
	ConfigCmd.AddCommand(ConfigListCmd)

}

func configList() {
	fmt.Printf("Tracking %v namespace(s):\n",len(Conf.NameSpaces))
	for _,v := range Conf.NameSpaces {
		fmt.Printf(" - %v:\n", v.Namespace)
	}
	for _,v := range Conf.NameSpaces {
		fmt.Printf("Deployments in namespace %v:\n", v.Namespace)
		for k,_ := range v.Deployments {
			fmt.Printf(" - %v\n", v.Deployments[k])
		}
	}
	fmt.Printf("Across %v kubernetes cluster API's:\n", len(Conf.Clusters))
	for _,c := range Conf.Clusters {
		fmt.Printf(" - %v\n", c.Url)
	}
}