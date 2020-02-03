// Copyright 2019 TriggerMesh, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clusterbuildtemplate

import (
	buildv1alpha1 "github.com/knative/build/pkg/apis/build/v1alpha1"
	"github.com/triggermesh/tm/pkg/client"
	"github.com/triggermesh/tm/pkg/printer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (cbt *ClusterBuildtemplate) GetTable(list *buildv1alpha1.ClusterBuildTemplateList) printer.Table {
	table := printer.Table{
		Headers: []string{
			"Name",
		},
		Rows: make([][]string, 0, len(list.Items)),
	}

	for _, item := range list.Items {
		table.Rows = append(table.Rows, cbt.Row(&item))
	}
	return table
}

func (cbt *ClusterBuildtemplate) Row(item *buildv1alpha1.ClusterBuildTemplate) []string {
	name := item.Name

	row := []string{
		name,
	}

	return row
}

func (cbt *ClusterBuildtemplate) List(clientset *client.ConfigSet) (*buildv1alpha1.ClusterBuildTemplateList, error) {
	return clientset.Build.BuildV1alpha1().ClusterBuildTemplates().List(metav1.ListOptions{})
}
