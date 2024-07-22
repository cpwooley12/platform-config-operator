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

// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=certmanagers,verbs=get;list;watch;create;update;patch;delete

// CreateCertManagerConfig creates the CertManager resource with name config.
func CreateCertManagerConfig(
	parent *deployv1alpha1.PlatformConfig,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "certificates.platform.tbd.io/v1alpha1",
			"kind":       "CertManager",
			"metadata": map[string]interface{}{
				"name": "config",
				"labels": map[string]interface{}{
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/version":            "unstable",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
				},
				"annotations": map[string]interface{}{
					"operator-builder.nukleros.io/ready-path":  ".status.created",
					"operator-builder.nukleros.io/ready-value": "true",
				},
			},
			"spec": map[string]interface{}{
				"namespace": parent.Spec.Platform.Certificates.Namespace, //  controlled by field: platform.certificates.namespace
				"aws": map[string]interface{}{
					"roleARN": "",
				},
				"injector": map[string]interface{}{
					"replicas": 2,
					"image":    "quay.io/jetstack/cert-manager-cainjector:v1.14.4",
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"cpu":    "50m",
							"memory": "64Mi",
						},
						"limits": map[string]interface{}{
							"memory": "128Mi",
						},
					},
				},
				"controller": map[string]interface{}{
					"replicas": 2,
					"image":    "quay.io/jetstack/cert-manager-controller:v1.14.4",
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"cpu":    "25m",
							"memory": "32Mi",
						},
						"limits": map[string]interface{}{
							"memory": "64Mi",
						},
					},
				},
				"webhook": map[string]interface{}{
					"replicas": 2,
					"image":    "quay.io/jetstack/cert-manager-webhook:v1.14.4",
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"cpu":    "25m",
							"memory": "32Mi",
						},
						"limits": map[string]interface{}{
							"memory": "64Mi",
						},
					},
				},
			},
		},
	}

	return mutate.MutateCertManagerConfig(resourceObj, parent, reconciler, req)
}
