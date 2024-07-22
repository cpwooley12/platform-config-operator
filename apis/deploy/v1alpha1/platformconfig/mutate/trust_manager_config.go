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

	"sigs.k8s.io/controller-runtime/pkg/client"
	trustmanagerv1alpha1 "github.com/tbd-paas/capabilities-trust-manager-operator/apis/trustmanager/v1alpha1"
	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	deployv1alpha1 "github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1"
)

const (
	// deployment sizes
	Small  = "small"
	Medium = "medium"
	Large  = "large"

	// controller resource requests, limits and replicas
	smallControllerReplicas  = 1
	mediumControllerReplicas = 1
	largeControllerReplicas  = 2
	
	smallControllerCPURequests    = "25m"
	smallControllerMemoryRequests = "32Mi"
	smallControllerMemoryLimits   = "64Mi"

	mediumControllerCPURequests    = "50m"
	mediumControllerMemoryRequests = "64Mi"
	mediumControllerMemoryLimits   = "96Mi"

	largeControllerCPURequests    = "50m"
	largeControllerMemoryRequests = "64Mi"
	largeControllerMemoryLimits   = "96Mi"
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

	trustManager, ok := original.(*trustmanagerv1alpha1.TrustManager)
	if !ok {
		return nil, fmt.Errorf("original object is not a TrustManager - found  %T", original)
	}

	//apply the trustmanager configuration
	if err := applyTrustManagerConfig(trustmanager, parent.Spec.Platform.Identity.DeploymentSize); err != nil {
		return nil, err
	}

	return []client.Object{trustManager}, nil
}

func applyTrustManagerConfig(trustManager *trustmanagerv1alpha1.TrustManager, deploymentSize string) error {
	switch deploymentSize {
	case Small:
		trustManager.Spec.Controller.Replicas = smallControllerReplicas
		trustManager.Spec.Controller.Resources.Requests.CPU = smallControllerCPURequests
		trustManager.Spec.Controller.Resources.Requests.Memory = smallControllerMemoryRequests
		trustManager.Spec.Controller.Resources.Limits.Memory = smallControllerMemoryLimits
	case Medium:
		trustManager.Spec.Controller.Replicas = mediumControllerReplicas
		trustManager.Spec.Controller.Resources.Requests.CPU = mediumControllerCPURequests
		trustManager.Spec.Controller.Resources.Requests.Memory = mediumControllerMemoryRequests
		trustManager.Spec.Controller.Resources.Limits.Memory = mediumControllerMemoryLimits
	case Large:
		trustManager.Spec.Controller.Replicas = largeControllerReplicas
		trustManager.Spec.Controller.Resources.Requests.CPU = largeControllerCPURequests
		trustManager.Spec.Controller.Resources.Requests.Memory = largeControllerMemoryRequests
		trustManager.Spec.Controller.Resources.Limits.Memory = largeControllerMemoryLimits
	default:
		return fmt.Errorf("invalid deployment size %s", deploymentSize)
	}

	return nil
}
