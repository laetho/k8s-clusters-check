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
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s-clusters-check/pkg/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"reflect"
	"strings"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare deployments on clusters",
	Long:  `Compare deployments on clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		Conf = initConfig()
		cmdCompare()
	},
}

func init() {
	RootCmd.AddCommand(compareCmd)
}


func cmdCompare() {
	glog.V(2).Infof("Number of clusters in config: %v", len(Conf.Clusters))
	glog.V(2).Infof("Number of namespaces in config: %v", len(Conf.NameSpaces))

	maps := PodMaps{}

	// Get maps from all clusters
	for _,c := range Conf.Clusters {
		config := rest.Config{
			Host : c.Url,
			BearerToken: c.Token,
		}
		client := k8sclient.GetCoreRESTConfigClient(&config)
		glog.V(2).Info(config.Host)
		maps = append(maps, podMapFromCluster(client))
	}
	glog.V(3).Infof("PodMaps: %v",maps)
	if len(maps) == 2 {
		if reflect.DeepEqual(maps[0], maps[1]) {
			glog.Info("OK, The clusters are equal")
			fmt.Println("OK, The clusters are equal")
		} else {
			glog.Error("ERROR, The clusters differ")
			fmt.Println("ERROR, The clusters differ")
		}
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func podMapFromCluster(c *corev1client.CoreV1Client) PodMap {

	podmap := PodMap{}

	for _, ns := range Conf.NameSpaces {
		lst, err := c.Pods(ns.Namespace).List(metav1.ListOptions{})
		if err != nil {
			glog.Warningf("Unable to fetch pods for cluster: %v", err)
			return podmap
		}

		for _,pod := range lst.Items {
			if !StringInSlice(pod.Labels["deploymentconfig"],ns.Deployments) {
				continue
			}
			// build podmap, must handle multiple containers in pod
			glog.V(2).Infof("Containers in pod, %v : %v", pod.Name ,len(pod.Spec.Containers))
			for _, cnt := range pod.Spec.Containers {
				glog.V(3).Infof("Name: %v",cnt.Name)
				glog.V(3).Infof("Image: %v",cnt.Image)
				imageref := strings.SplitAfterN(cnt.Image,"/", 2)
				podmap[ns.Namespace+"."+pod.Labels["deploymentconfig"]+"."+cnt.Name] = imageref[1]
			}
		}
	}

	return podmap
}