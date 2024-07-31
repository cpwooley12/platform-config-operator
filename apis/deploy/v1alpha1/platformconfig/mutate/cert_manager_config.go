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
	certManagerSmall  = "small"
	certManagerMedium = "medium"
	certManagerLarge  = "large"

	// injector resource requests, limits and replicas
	certManagersmallInjectorReplicas  = 1
	certManagermediumInjectorReplicas = 1
	certManagerlargeInjectorReplicas  = 2

	certManagersmallInjectorCPURequests    = "50m"
	certManagersmallInjectorMemoryRequests = "64Mi"
	certManagersmallInjectorMemoryLimits   = "128Mi"

	certManagermediumInjectorCPURequests    = "100m"
	certManagermediumInjectorMemoryRequests = "128Mi"
	certManagermediumInjectorMemoryLimits   = "256Mi"

	certManagerlargeInjectorCPURequests    = "150m"
	certManagerlargeInjectorMemoryRequests = "192Mi"
	certManagerlargeInjectorMemoryLimits   = "384Mi"

	// controller resource requests, limits and replicas
	certManagersmallControllerReplicas  = 1
	certManagermediumControllerReplicas = 1
	certManagerlargeControllerReplicas  = 2

	certManagersmallControllerCPURequests    = "25m"
	certManagersmallControllerMemoryRequests = "32Mi"
	certManagersmallControllerMemoryLimits   = "64Mi"

	certManagermediumControllerCPURequests    = "50m"
	certManagermediumControllerMemoryRequests = "64Mi"
	certManagermediumControllerMemoryLimits   = "96Mi"

	certManagerlargeControllerCPURequests    = "50m"
	certManagerlargeControllerMemoryRequests = "64Mi"
	certManagerlargeControllerMemoryLimits   = "96Mi"

	// webhook resource requests, limits and replicas
	certManagersmallWebhookReplicas  = 1
	certManagermediumWebhookReplicas = 1
	certManagerlargeWebhookReplicas  = 2

	certManagersmallWebhookCPURequests    = "25m"
	certManagersmallWebhookMemoryRequests = "32Mi"
	certManagersmallWebhookMemoryLimits   = "64Mi"

	certManagermediumWebhookCPURequests    = "50m"
	certManagermediumWebhookMemoryRequests = "64Mi"
	certManagermediumWebhookMemoryLimits   = "96Mi"

	certManagerlargeWebhookCPURequests    = "50m"
	certManagerlargeWebhookMemoryRequests = "64Mi"
	certManagerlargeWebhookMemoryLimits   = "96Mi"
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

	certManager := &certificatesv1alpha1.CertManager{}
	if certManager.Kind != "CertManager" {
		certManager.Kind = "CertManager"
	}

	if certManager.APIVersion != "certificates.platform.tbd.io/v1alpha1" {
		certManager.APIVersion = "certificates.platform.tbd.io/v1alpha1"
	}

	if certManager.Namespace != parent.Namespace {
		certManager.Namespace = parent.Namespace
	}
	// Ensure the Name is set
	if certManager.Name == "" {
		certManager.Name = "certificaterequests.cert-manager.io" // Set a default name or use a more appropriate naming strategy
	}
	err := resources.ToTyped(original, certManager)
	if err != nil {
		return nil, fmt.Errorf("failed to convert object to CertManager type: %w", err)
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
	case certManagerSmall:
		// modify injector for small deployment
		certManager.Spec.Injector.Replicas = certManagersmallInjectorReplicas
		certManager.Spec.Injector.Resources.Requests.Cpu = certManagersmallInjectorCPURequests
		certManager.Spec.Injector.Resources.Requests.Memory = certManagersmallInjectorMemoryRequests
		certManager.Spec.Injector.Resources.Limits.Memory = certManagersmallInjectorMemoryLimits

		// modify controller for small deployment
		certManager.Spec.Controller.Replicas = certManagersmallControllerReplicas
		certManager.Spec.Controller.Resources.Requests.Cpu = certManagersmallControllerCPURequests
		certManager.Spec.Controller.Resources.Requests.Memory = certManagersmallControllerMemoryRequests
		certManager.Spec.Controller.Resources.Limits.Memory = certManagersmallControllerMemoryLimits

		// modify webhook for small deployment
		certManager.Spec.Webhook.Replicas = certManagersmallWebhookReplicas
		certManager.Spec.Webhook.Resources.Requests.Cpu = certManagersmallWebhookCPURequests
		certManager.Spec.Webhook.Resources.Requests.Memory = certManagersmallWebhookMemoryRequests
		certManager.Spec.Webhook.Resources.Limits.Memory = certManagersmallWebhookMemoryLimits
	case certManagerMedium:
		// modify injector for medium deployment
		certManager.Spec.Injector.Replicas = certManagermediumInjectorReplicas
		certManager.Spec.Injector.Resources.Requests.Cpu = certManagermediumInjectorCPURequests
		certManager.Spec.Injector.Resources.Requests.Memory = certManagermediumInjectorMemoryRequests
		certManager.Spec.Injector.Resources.Limits.Memory = certManagermediumInjectorMemoryLimits

		// modify controller for medium deployment
		certManager.Spec.Controller.Replicas = certManagermediumControllerReplicas
		certManager.Spec.Controller.Resources.Requests.Cpu = certManagermediumControllerCPURequests
		certManager.Spec.Controller.Resources.Requests.Memory = certManagermediumControllerMemoryRequests
		certManager.Spec.Controller.Resources.Limits.Memory = certManagermediumControllerMemoryLimits

		// modify webhook for medium deployment
		certManager.Spec.Webhook.Replicas = certManagermediumWebhookReplicas
		certManager.Spec.Webhook.Resources.Requests.Cpu = certManagermediumWebhookCPURequests
		certManager.Spec.Webhook.Resources.Requests.Memory = certManagermediumWebhookMemoryRequests
		certManager.Spec.Webhook.Resources.Limits.Memory = certManagermediumWebhookMemoryLimits
	case certManagerLarge:
		// modify injector for large deployment
		certManager.Spec.Injector.Replicas = certManagerlargeInjectorReplicas
		certManager.Spec.Injector.Resources.Requests.Cpu = certManagerlargeInjectorCPURequests
		certManager.Spec.Injector.Resources.Requests.Memory = certManagerlargeInjectorMemoryRequests
		certManager.Spec.Injector.Resources.Limits.Memory = certManagerlargeInjectorMemoryLimits

		// modify controller for large deployment
		certManager.Spec.Controller.Replicas = certManagerlargeControllerReplicas
		certManager.Spec.Controller.Resources.Requests.Cpu = certManagerlargeControllerCPURequests
		certManager.Spec.Controller.Resources.Requests.Memory = certManagerlargeControllerMemoryRequests
		certManager.Spec.Controller.Resources.Limits.Memory = certManagerlargeControllerMemoryLimits

		// modify webhook for large deployment
		certManager.Spec.Webhook.Replicas = certManagerlargeWebhookReplicas
		certManager.Spec.Webhook.Resources.Requests.Cpu = certManagerlargeWebhookCPURequests
		certManager.Spec.Webhook.Resources.Requests.Memory = certManagerlargeWebhookMemoryRequests
		certManager.Spec.Webhook.Resources.Limits.Memory = certManagerlargeWebhookMemoryLimits
	default:
		return fmt.Errorf("invalid deployment size %s", deploymentSize)
	}

	return nil
}
