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
	certificatesv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/certificates/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	deployv1alpha1 "github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1"
)

const (
	// deployment sizes
	Small  = "small"
	Medium = "medium"
	Large  = "large"

	// injector resource requests, limits and replicas
	smallInjectorReplicas  = 1
	mediumInjectorReplicas = 1
	largeInjectorReplicas  = 2

	smallInjectorCPURequests    = "50m"
	smallInjectorMemoryRequests = "64Mi"
	smallInjectorMemoryLimits   = "128Mi"

	mediumInjectorCPURequests    = "100m"
	mediumInjectorMemoryRequests = "128Mi"
	mediumInjectorMemoryLimits   = "256Mi"

	largeInjectorCPURequests    = "150m"
	largeInjectorMemoryRequests = "192Mi"
	largeInjectorMemoryLimits   = "384Mi"

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

	// webhook resource requests, limits and replicas
	smallWebhookReplicas  = 1
	mediumWebhookReplicas = 1
	largeWebhookReplicas  = 2

	smallWebhookCPURequests    = "25m"
	smallWebhookMemoryRequests = "32Mi"
	smallWebhookMemoryLimits   = "64Mi"

	mediumWebhookCPURequests    = "50m"
	mediumWebhookMemoryRequests = "64Mi"
	mediumWebhookMemoryLimits   = "96Mi"

	largeWebhookCPURequests    = "50m"
	largeWebhookMemoryRequests = "64Mi"
	largeWebhookMemoryLimits   = "96Mi"
)

// MutateCertManagerConfig mutates the CertManager resource with name config.
func MutateCertManagerConfig(
	original client.Object,
	parent *deployv1alpha1.PlatformConfig,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	certManager, ok := original.(*certificatesv1alpha1.CertManager)
	if !ok {
		return nil, fmt.Errorf("original object is not a CertManager - found %T", original)
	}

	// apply the appropriate values to the CertManager resource
	if err := applyCertManagerConfig(certManager, parent.Spec.Platform.Certificates.DeploymentSize); err != nil {
		return nil, err
	}

	return []client.Object{certManager}, nil
}

// applyCertManagerConfig checks if s,m,l and pass in appropriate values
func applyCertManagerConfig(certManager *certificatesv1alpha1.CertManager, deploymentSize string) error {
	switch deploymentSize {
	case Small:
		// modify injector for small deployment
		certManager.Spec.Injector.Replicas = smallInjectorReplicas
		certManager.Spec.Injector.Resources.Requests.Cpu = smallInjectorCPURequests
		certManager.Spec.Injector.Resources.Requests.Memory = smallInjectorMemoryRequests
		certManager.Spec.Injector.Resources.Limits.Memory = smallInjectorMemoryLimits

		// modify controller for small deployment
		certManager.Spec.Controller.Replicas = smallControllerReplicas
		certManager.Spec.Controller.Resources.Requests.Cpu = smallControllerCPURequests
		certManager.Spec.Controller.Resources.Requests.Memory = smallControllerMemoryRequests
		certManager.Spec.Controller.Resources.Limits.Memory = smallControllerMemoryLimits

		// modify webhook for small deployment
		certManager.Spec.Webhook.Replicas = smallWebhookReplicas
		certManager.Spec.Webhook.Resources.Requests.Cpu = smallWebhookCPURequests
		certManager.Spec.Webhook.Resources.Requests.Memory = smallWebhookMemoryRequests
		certManager.Spec.Webhook.Resources.Limits.Memory = smallWebhookMemoryLimits
	case Medium:
		// modify injector for medium deployment
		certManager.Spec.Injector.Replicas = mediumInjectorReplicas
		certManager.Spec.Injector.Resources.Requests.Cpu = mediumInjectorCPURequests
		certManager.Spec.Injector.Resources.Requests.Memory = mediumInjectorMemoryRequests
		certManager.Spec.Injector.Resources.Limits.Memory = mediumInjectorMemoryLimits

		// modify controller for medium deployment
		certManager.Spec.Controller.Replicas = mediumControllerReplicas
		certManager.Spec.Controller.Resources.Requests.Cpu = mediumControllerCPURequests
		certManager.Spec.Controller.Resources.Requests.Memory = mediumControllerMemoryRequests
		certManager.Spec.Controller.Resources.Limits.Memory = mediumControllerMemoryLimits

		// modify webhook for medium deployment
		certManager.Spec.Webhook.Replicas = mediumWebhookReplicas
		certManager.Spec.Webhook.Resources.Requests.Cpu = mediumWebhookCPURequests
		certManager.Spec.Webhook.Resources.Requests.Memory = mediumWebhookMemoryRequests
		certManager.Spec.Webhook.Resources.Limits.Memory = mediumWebhookMemoryLimits
	case Large:
		// modify injector for large deployment
		certManager.Spec.Injector.Replicas = largeInjectorReplicas
		certManager.Spec.Injector.Resources.Requests.Cpu = largeInjectorCPURequests
		certManager.Spec.Injector.Resources.Requests.Memory = largeInjectorMemoryRequests
		certManager.Spec.Injector.Resources.Limits.Memory = largeInjectorMemoryLimits

		// modify controller for large deployment
		certManager.Spec.Controller.Replicas = largeControllerReplicas
		certManager.Spec.Controller.Resources.Requests.Cpu = largeControllerCPURequests
		certManager.Spec.Controller.Resources.Requests.Memory = largeControllerMemoryRequests
		certManager.Spec.Controller.Resources.Limits.Memory = largeControllerMemoryLimits

		// modify webhook for large deployment
		certManager.Spec.Webhook.Replicas = largeWebhookReplicas
		certManager.Spec.Webhook.Resources.Requests.Cpu = largeWebhookCPURequests
		certManager.Spec.Webhook.Resources.Requests.Memory = largeWebhookMemoryRequests
		certManager.Spec.Webhook.Resources.Limits.Memory = largeWebhookMemoryLimits
	default:
		return fmt.Errorf("invalid deployment size %s", deploymentSize)
	}

	return nil
}
