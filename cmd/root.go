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
	"github.com/golang/glog"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var (
	CfgFile string = ""
	Conf *Config
)

var header = &doc.GenManHeader{
	Title: "k8s-clusters-check",
	Section: "3",
}

var RootCmd = &cobra.Command{
	Use:   "k8s-clusters-check",
	Short: "k8s-clusters-check",
	Long: `Runs against a list of k8s clusters and monitors namespaces and deployment resources differences.`,
	//Run: func(cmd *cobra.Command, args []string) {
	// Do Stuff Here
	//},
}



func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

func initConfig() *Config {
	viper.SetConfigType("json")

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home + "/.config")
	viper.SetConfigName("k8scc")

	if len(CfgFile) > 0 {
		viper.SetConfigFile(CfgFile)
	}

	viper.AutomaticEnv()
	conf := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		glog.Errorf("%v", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(conf); err != nil {
		glog.Errorf("Failed to unmarshal %v", conf)
		glog.Errorf("%v", err)
		os.Exit(1)
	}

	return conf
}

func Execute() error {
	if err := RootCmd.Execute(); err != nil {
		return err
	}
	return nil
}