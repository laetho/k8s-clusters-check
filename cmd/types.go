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

type Config struct {
	Clusters []ConfigCluster			`json:"clusters"`
	NameSpaces []ConfigNamespace 	`json:"namespaces"`
}

type ConfigCluster struct {
	Url string 						`json:"url"`	// Master api url
	User string						`json:"user"`	// Typically a service account
	Token string					`json:"token"`	// Token for service account
}

type ConfigNamespace struct {
	Namespace string				`json:"namespace"`
	Deployments []ConfigDeployment	`json:"deployments"`
}

type ConfigDeployment struct {
	Deployment string
}

