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

package mutate

import (
	"fmt"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/resources"
	certificatesv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/certificates/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	deployv1alpha1 "github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1"
)

const (
	// deployment sizes
	trustManagerSmall  = "small"
	trustManagerMedium = "medium"
	trustManagerLarge  = "large"

	// controller resource requests, limits and replicas
	trustManagersmallControllerReplicas  = 1
	trustManagermediumControllerReplicas = 1
	trustManagerlargeControllerReplicas  = 2

	trustManagersmallControllerCPURequests    = "25m"
	trustManagersmallControllerMemoryRequests = "32Mi"
	trustManagersmallControllerMemoryLimits   = "64Mi"

	trustManagermediumControllerCPURequests    = "50m"
	trustManagermediumControllerMemoryRequests = "64Mi"
	trustManagermediumControllerMemoryLimits   = "96Mi"

	trustManagerlargeControllerCPURequests    = "50m"
	trustManagerlargeControllerMemoryRequests = "64Mi"
	trustManagerlargeControllerMemoryLimits   = "96Mi"
)

// MutateTrustManagerConfig mutates the TrustManager resource with name config.
func MutateTrustManagerConfig(
	original client.Object,
	parent *deployv1alpha1.PlatformConfig,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	trustManager := &certificatesv1alpha1.TrustManager{}
	err := resources.ToTyped(trustManager, original)
	if err != nil {
		return nil, fmt.Errorf("failed to convert object to TrustManager type: %w", err)
	}

	// trustManager, ok := original.(*certificatesv1alpha1.TrustManager)
	// if !ok {
	// 	return nil, fmt.Errorf("original object is not a TrustManager - found  %T", original)
	// }

	//apply the trustmanager configuration
	if err := applyTrustManagerConfig(trustManager, parent.Spec.Platform.Identity.DeploymentSize); err != nil {
		return nil, err
	}

	return []client.Object{trustManager}, nil
}

func applyTrustManagerConfig(trustManager *certificatesv1alpha1.TrustManager, deploymentSize string) error {
	switch deploymentSize {
	case trustManagerSmall:
		trustManager.Spec.Controller.Replicas = trustManagersmallControllerReplicas
		trustManager.Spec.Controller.Resources.Requests.Cpu = trustManagersmallControllerCPURequests
		trustManager.Spec.Controller.Resources.Requests.Memory = trustManagersmallControllerMemoryRequests
		trustManager.Spec.Controller.Resources.Limits.Memory = trustManagersmallControllerMemoryLimits
	case trustManagerMedium:
		trustManager.Spec.Controller.Replicas = trustManagermediumControllerReplicas
		trustManager.Spec.Controller.Resources.Requests.Cpu = trustManagermediumControllerCPURequests
		trustManager.Spec.Controller.Resources.Requests.Memory = trustManagermediumControllerMemoryRequests
		trustManager.Spec.Controller.Resources.Limits.Memory = trustManagermediumControllerMemoryLimits
	case trustManagerLarge:
		trustManager.Spec.Controller.Replicas = trustManagerlargeControllerReplicas
		trustManager.Spec.Controller.Resources.Requests.Cpu = trustManagerlargeControllerCPURequests
		trustManager.Spec.Controller.Resources.Requests.Memory = trustManagerlargeControllerMemoryRequests
		trustManager.Spec.Controller.Resources.Limits.Memory = trustManagerlargeControllerMemoryLimits
	default:
		return fmt.Errorf("invalid deployment size %s", deploymentSize)
	}

	return nil
}
