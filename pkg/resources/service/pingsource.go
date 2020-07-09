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

package service

import (
	"github.com/triggermesh/tm/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	eventingv1alpha2 "knative.dev/eventing/pkg/apis/sources/v1alpha2"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

func (s *Service) pingSource(cron string, data string) *eventingv1alpha2.PingSource {
	return &eventingv1alpha2.PingSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      s.Name,
			Namespace: s.Namespace,
			Labels: map[string]string{
				"cli.triggermesh.io/schedule": s.Name,
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "PingSource",
			APIVersion: "sources.knative.dev/v1alpha2",
		},
		Spec: eventingv1alpha2.PingSourceSpec{
			Schedule: cron,
			JsonData: data,
			SourceSpec: duckv1.SourceSpec{
				Sink: duckv1.Destination{
					Ref: &duckv1.KReference{
						APIVersion: "serving.knative.dev/v1alpha1",
						Kind:       "Service",
						Name:       s.Name,
						Namespace:  s.Namespace,
					},
				},
			},
		},
	}
}

func (s *Service) createPingSource(ps *eventingv1alpha2.PingSource, clientset *client.ConfigSet) error {
	return nil
}
