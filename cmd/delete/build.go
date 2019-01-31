/*
Copyright (c) 2018 TriggerMesh, Inc

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

package delete

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/triggermesh/tm/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Build struct {
	Name      string
	Namespace string
}

var b Build

func cmdDeleteBuild(clientset *client.ConfigSet) *cobra.Command {
	return &cobra.Command{
		Use:     "build",
		Aliases: []string{"builds"},
		Short:   "Delete knative build resource",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			b.Name = args[0]
			b.Namespace = client.Namespace
			if err := b.DeleteBuild(clientset); err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Build is being deleted")
		},
	}
}

// DeleteBuild removes knative build object
func (b *Build) DeleteBuild(clientset *client.ConfigSet) error {
	return clientset.Build.BuildV1alpha1().Builds(b.Namespace).Delete(b.Name, &metav1.DeleteOptions{})
}
