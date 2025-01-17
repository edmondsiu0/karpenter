/*
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

package utils

import (
	"fmt"
	"regexp"

	v1 "k8s.io/api/core/v1"

	corev1alpha1 "github.com/aws/karpenter-core/pkg/apis/v1alpha1"
)

// ParseInstanceID parses the provider ID stored on the node to get the instance ID
// associated with a node
func ParseInstanceID(node *v1.Node) (string, error) {
	r := regexp.MustCompile(`aws:///(?P<AZ>.*)/(?P<InstanceID>.*)`)
	matches := r.FindStringSubmatch(node.Spec.ProviderID)
	if matches == nil {
		return "", fmt.Errorf("parsing instance id %s", node.Spec.ProviderID)
	}
	for i, name := range r.SubexpNames() {
		if name == "InstanceID" {
			return matches[i], nil
		}
	}
	return "", fmt.Errorf("parsing instance id %s", node.Spec.ProviderID)
}

// ParseInstanceID parses the provider ID stored on the node to get the instance ID
// associated with a node
func ParseMachineInstanceID(machine *corev1alpha1.Machine) (string, error) {
	r := regexp.MustCompile(`aws:///(?P<AZ>.*)/(?P<InstanceID>.*)`)
	matches := r.FindStringSubmatch(machine.Status.ProviderID)
	if matches == nil {
		return "", fmt.Errorf("parsing instance id %s", machine.Status.ProviderID)
	}
	for i, name := range r.SubexpNames() {
		if name == "InstanceID" {
			return matches[i], nil
		}
	}
	return "", fmt.Errorf("parsing instance id %s", machine.Status.ProviderID)
}
