/*
Copyright 2024.

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

package platformconfig

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	deployv1alpha1 "github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1"
	"github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1/platformconfig/mutate"
)

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// CreateNamespacePlatformCertificatesNamespace creates the Namespace resource with name parent.Spec.Platform.Certificates.Namespace.
func CreateNamespacePlatformCertificatesNamespace(
	parent *deployv1alpha1.PlatformConfig,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				// controlled by field: platform.certificates.namespace
				// Namespace where
				//  the capability components will be deployed.
				"name": parent.Spec.Platform.Certificates.Namespace,
				"labels": map[string]interface{}{
					"capabilities.tbd.io/capability":                "platform-config",
					"capabilities.tbd.io/version":                   "v0.0.1",
					"capabilities.tbd.io/platform-version":          "unstable",
					"app.kubernetes.io/version":                     "unstable",
					"app.kubernetes.io/part-of":                     "platform",
					"app.kubernetes.io/managed-by":                  "platform-config-operator",
					"certificates.platform.tbd.io/inject-ca-bundle": "true",
				},
				"annotations": map[string]interface{}{
					"operator-builder.nukleros.io/ready-path":  ".status.created",
					"operator-builder.nukleros.io/ready-value": "true",
					// controlled by field: platform.certificates.deploymentSize
					// Size of the
					//  deployment for the underlying capability.  Must be one of small, medium, or large.
					//  +kubebuilder:validation:Enum:small;medium;large
					"certificates.platform.tbd.io/deployment-size": parent.Spec.Platform.Certificates.DeploymentSize,
				},
			},
		},
	}

	return mutate.MutateNamespacePlatformCertificatesNamespace(resourceObj, parent, reconciler, req)
}
