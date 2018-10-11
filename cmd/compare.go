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
	"fmt"
	"os"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s-clusters-check/pkg/client"
	"k8s.io/client-go/rest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare deployments on clusters",
	Long:  `Compare deployments on clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdCompare()
	},
}

func init() {
	RootCmd.AddCommand(compareCmd)
}


func cmdCompare() {
	glog.Infof("Number of clusters in config: %v", len(Conf.Clusters))
	glog.Infof("Number of namespaces in config: %v", len(Conf.NameSpaces))

	// @todo replace this with actual
	config := rest.Config{
		Host: "https://master.ocp.norsk-tipping.no:8443",
		UserAgent: "k8s-clusters-check",
		BearerToken: "ZAGqtBd0LIzgpBvqWwvUkrjxhJ9oEsp-OUYg-7g68hk",
	}

	for k,v := range Conf.Clusters {
		fmt.Println(k,v)
	}



}

func podMapFromCluster(c *corev1client.CoreV1Client, config *rest.Config) PodMap {
	client := client.GetCoreRESTConfigClient(config)
	lst, err := client.Pods("vapidev").List(metav1.ListOptions{})
	if err != nil {
		glog.Errorf("%v",err)
		os.Exit(1)
	}
	fmt.Println(lst)

	return PodMap{}
}