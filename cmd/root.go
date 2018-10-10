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
	"flag"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

var (
	CfgFile string = ""
)

func init() {

}

func initConfig() {
	viper.SetConfigType("json")

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home + "/.config/")
	viper.SetConfigName("k8scc.json")

	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read:", viper.ConfigFileUsed())
	}
}

func Execute() error {
	initConfig()
	flag.Parse()
	if err := RootCmd.Execute(); err != nil {
		return err
	}
	return nil
}