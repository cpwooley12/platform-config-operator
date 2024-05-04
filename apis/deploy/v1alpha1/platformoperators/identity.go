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

package platformoperators

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	deployv1alpha1 "github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1"
	"github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1/platformoperators/mutate"
)

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDAwspodidentitywebhooksIdentityPlatformTbdIo creates the CustomResourceDefinition resource with name awspodidentitywebhooks.identity.platform.tbd.io.
func CreateCRDAwspodidentitywebhooksIdentityPlatformTbdIo(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.14.0",
				},
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "awspodidentitywebhooks.identity.platform.tbd.io",
			},
			"spec": map[string]interface{}{
				"group": "identity.platform.tbd.io",
				"names": map[string]interface{}{
					"kind":     "AWSPodIdentityWebhook",
					"listKind": "AWSPodIdentityWebhookList",
					"plural":   "awspodidentitywebhooks",
					"singular": "awspodidentitywebhook",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "AWSPodIdentityWebhook is the Schema for the awspodidentitywebhooks API.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "AWSPodIdentityWebhookSpec defines the desired state of AWSPodIdentityWebhook.",
										"properties": map[string]interface{}{
											"image": map[string]interface{}{
												"default": "amazon/amazon-eks-pod-identity-webhook:v0.5.3",
												"description": `(Default: "amazon/amazon-eks-pod-identity-webhook:v0.5.3")


	Image to use for AWS pod identity webhook deployment.`,
												"type": "string",
											},
											"namespace": map[string]interface{}{
												"default":     "tbd-identity-system",
												"description": "(Default: \"tbd-identity-system\")",
												"type":        "string",
											},
											"replicas": map[string]interface{}{
												"default": 2,
												"description": `(Default: 2)


	Number of replicas to use for the AWS pod identity webhook deployment.`,
												"type": "integer",
											},
											"resources": map[string]interface{}{
												"properties": map[string]interface{}{
													"limits": map[string]interface{}{
														"properties": map[string]interface{}{
															"memory": map[string]interface{}{
																"default": "64Mi",
																"description": `(Default: "64Mi")


	Memory limits to use for AWS pod identity webhook deployment.`,
																"type": "string",
															},
														},
														"type": "object",
													},
													"requests": map[string]interface{}{
														"properties": map[string]interface{}{
															"cpu": map[string]interface{}{
																"default": "25m",
																"description": `(Default: "25m")


	CPU requests to use for AWS pod identity webhook deployment.`,
																"type": "string",
															},
															"memory": map[string]interface{}{
																"default": "32Mi",
																"description": `(Default: "32Mi")


	Memory requests to use for AWS pod identity webhook deployment.`,
																"type": "string",
															},
														},
														"type": "object",
													},
												},
												"type": "object",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "AWSPodIdentityWebhookStatus defines the observed state of AWSPodIdentityWebhook.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"description": `PhaseCondition describes an event that has occurred during a phase
of the controller reconciliation loop.`,
													"properties": map[string]interface{}{
														"lastModified": map[string]interface{}{
															"description": "LastModified defines the time in which this component was updated.",
															"type":        "string",
														},
														"message": map[string]interface{}{
															"description": "Message defines a helpful message from the phase.",
															"type":        "string",
														},
														"phase": map[string]interface{}{
															"description": "Phase defines the phase in which the condition was set.",
															"type":        "string",
														},
														"state": map[string]interface{}{
															"description": "PhaseState defines the current state of the phase.",
															"enum": []interface{}{
																"Complete",
																"Reconciling",
																"Failed",
																"Pending",
															},
															"type": "string",
														},
													},
													"required": []interface{}{
														"lastModified",
														"message",
														"phase",
														"state",
													},
													"type": "object",
												},
												"type": "array",
											},
											"created": map[string]interface{}{
												"type": "boolean",
											},
											"dependenciesSatisfied": map[string]interface{}{
												"type": "boolean",
											},
											"resources": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "ChildResource is the resource and its condition as stored on the workload custom resource's status field.",
													"properties": map[string]interface{}{
														"condition": map[string]interface{}{
															"description": "ResourceCondition defines the current condition of this resource.",
															"properties": map[string]interface{}{
																"created": map[string]interface{}{
																	"description": "Created defines whether this object has been successfully created or not.",
																	"type":        "boolean",
																},
																"lastModified": map[string]interface{}{
																	"description": "LastModified defines the time in which this resource was updated.",
																	"type":        "string",
																},
																"message": map[string]interface{}{
																	"description": "Message defines a helpful message from the resource phase.",
																	"type":        "string",
																},
															},
															"required": []interface{}{
																"created",
															},
															"type": "object",
														},
														"group": map[string]interface{}{
															"description": "Group defines the API Group of the resource.",
															"type":        "string",
														},
														"kind": map[string]interface{}{
															"description": "Kind defines the kind of the resource.",
															"type":        "string",
														},
														"name": map[string]interface{}{
															"description": "Name defines the name of the resource from the metadata.name field.",
															"type":        "string",
														},
														"namespace": map[string]interface{}{
															"description": "Namespace defines the namespace in which this resource exists in.",
															"type":        "string",
														},
														"version": map[string]interface{}{
															"description": "Version defines the API Version of the resource.",
															"type":        "string",
														},
													},
													"required": []interface{}{
														"group",
														"kind",
														"name",
														"namespace",
														"version",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDAwspodidentitywebhooksIdentityPlatformTbdIo(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceIdentityOperatorControllerManager creates the ServiceAccount resource with name identity-operator-controller-manager.
func CreateServiceAccountNamespaceIdentityOperatorControllerManager(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name":      "identity-operator-controller-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceIdentityOperatorControllerManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateRoleNamespaceIdentityOperatorLeaderElectionRole creates the Role resource with name identity-operator-leader-election-role.
func CreateRoleNamespaceIdentityOperatorLeaderElectionRole(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name":      "identity-operator-leader-election-role",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"update",
						"patch",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"update",
						"patch",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"events",
					},
					"verbs": []interface{}{
						"create",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceIdentityOperatorLeaderElectionRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=create;get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;get;patch;update
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;get;patch;update
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=identity.platform.tbd.io,resources=awspodidentitywebhooks,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=identity.platform.tbd.io,resources=awspodidentitywebhooks/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=create;delete;get;list;patch;update;watch

// CreateClusterRoleIdentityOperatorManagerRole creates the ClusterRole resource with name identity-operator-manager-role.
func CreateClusterRoleIdentityOperatorManagerRole(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "identity-operator-manager-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"mutatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"deployments",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.k8s.io",
					},
					"resources": []interface{}{
						"certificatesigningrequests",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"events",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"namespaces",
					},
					"verbs": []interface{}{
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"serviceaccounts",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"services",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"identity.platform.tbd.io",
					},
					"resources": []interface{}{
						"awspodidentitywebhooks",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"identity.platform.tbd.io",
					},
					"resources": []interface{}{
						"awspodidentitywebhooks/status",
					},
					"verbs": []interface{}{
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"rbac.authorization.k8s.io",
					},
					"resources": []interface{}{
						"clusterrolebindings",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"rbac.authorization.k8s.io",
					},
					"resources": []interface{}{
						"clusterroles",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"rbac.authorization.k8s.io",
					},
					"resources": []interface{}{
						"rolebindings",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"rbac.authorization.k8s.io",
					},
					"resources": []interface{}{
						"roles",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIdentityOperatorManagerRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:verbs=get,urls=/metrics

// CreateClusterRoleIdentityOperatorMetricsReader creates the ClusterRole resource with name identity-operator-metrics-reader.
func CreateClusterRoleIdentityOperatorMetricsReader(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "identity-operator-metrics-reader",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"nonResourceURLs": []interface{}{
						"/metrics",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIdentityOperatorMetricsReader(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=authentication.k8s.io,resources=tokenreviews,verbs=create
// +kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create

// CreateClusterRoleIdentityOperatorProxyRole creates the ClusterRole resource with name identity-operator-proxy-role.
func CreateClusterRoleIdentityOperatorProxyRole(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "identity-operator-proxy-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authentication.k8s.io",
					},
					"resources": []interface{}{
						"tokenreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authorization.k8s.io",
					},
					"resources": []interface{}{
						"subjectaccessreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIdentityOperatorProxyRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceIdentityOperatorLeaderElectionRolebinding creates the RoleBinding resource with name identity-operator-leader-election-rolebinding.
func CreateRoleBindingNamespaceIdentityOperatorLeaderElectionRolebinding(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name":      "identity-operator-leader-election-rolebinding",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "identity-operator-leader-election-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "identity-operator-controller-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceIdentityOperatorLeaderElectionRolebinding(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIdentityOperatorManagerRolebinding creates the ClusterRoleBinding resource with name identity-operator-manager-rolebinding.
func CreateClusterRoleBindingIdentityOperatorManagerRolebinding(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "identity-operator-manager-rolebinding",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "identity-operator-manager-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "identity-operator-controller-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIdentityOperatorManagerRolebinding(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIdentityOperatorProxyRolebinding creates the ClusterRoleBinding resource with name identity-operator-proxy-rolebinding.
func CreateClusterRoleBindingIdentityOperatorProxyRolebinding(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "identity-operator-proxy-rolebinding",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "identity-operator-proxy-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "identity-operator-controller-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIdentityOperatorProxyRolebinding(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceIdentityOperatorControllerManagerMetricsService creates the Service resource with name identity-operator-controller-manager-metrics-service.
func CreateServiceNamespaceIdentityOperatorControllerManagerMetricsService(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
					"control-plane":                        "controller-manager",
				},
				"name":      "identity-operator-controller-manager-metrics-service",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "https",
						"port":       8443,
						"protocol":   "TCP",
						"targetPort": "https",
					},
				},
				"selector": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "v0.0.1",
					"capabilities.tbd.io/capability":       "identity",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
					"control-plane":                        "controller-manager",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceIdentityOperatorControllerManagerMetricsService(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceIdentityOperatorControllerManager creates the Deployment resource with name identity-operator-controller-manager.
func CreateDeploymentNamespaceIdentityOperatorControllerManager(
	parent *deployv1alpha1.PlatformOperators,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                                  "identity-operator",
					"app.kubernetes.io/component":          "identity-operator",
					"app.kubernetes.io/created-by":         "capabilities-identity-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "identity-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
					"control-plane":                        "controller-manager",
				},
				"name":      "identity-operator-controller-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app":                                  "identity-operator",
						"app.kubernetes.io/component":          "identity-operator",
						"app.kubernetes.io/instance":           "manager",
						"app.kubernetes.io/managed-by":         "platform-config-operator",
						"app.kubernetes.io/name":               "identity-operator",
						"app.kubernetes.io/part-of":            "platform",
						"app.kubernetes.io/version":            "unstable",
						"capabilities.tbd.io/capability":       "identity",
						"capabilities.tbd.io/platform-version": "unstable",
						"capabilities.tbd.io/version":          "v0.0.1",
						"control-plane":                        "controller-manager",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"kubectl.kubernetes.io/default-container": "manager",
						},
						"labels": map[string]interface{}{
							"app":                                  "identity-operator",
							"app.kubernetes.io/component":          "identity-operator",
							"app.kubernetes.io/instance":           "manager",
							"app.kubernetes.io/managed-by":         "platform-config-operator",
							"app.kubernetes.io/name":               "identity-operator",
							"app.kubernetes.io/part-of":            "platform",
							"app.kubernetes.io/version":            "unstable",
							"capabilities.tbd.io/capability":       "identity",
							"capabilities.tbd.io/platform-version": "unstable",
							"capabilities.tbd.io/version":          "v0.0.1",
							"control-plane":                        "controller-manager",
						},
					},
					"spec": map[string]interface{}{
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"podAffinityTerm": map[string]interface{}{
											"labelSelector": map[string]interface{}{
												"matchExpressions": []interface{}{
													map[string]interface{}{
														"key":      "app.kubernetes.io/name",
														"operator": "In",
														"values": []interface{}{
															"identity-operator",
														},
													},
												},
											},
											"topologyKey": "kubernetes.io/hostname",
										},
										"weight": 100,
									},
								},
							},
						},
						"containers": []interface{}{
							map[string]interface{}{
								"args": []interface{}{
									"--secure-listen-address=0.0.0.0:8443",
									"--upstream=http://127.0.0.1:8080/",
									"--logtostderr=true",
									"--v=0",
								},
								"image": "gcr.io/kubebuilder/kube-rbac-proxy:v0.13.1",
								"name":  "kube-rbac-proxy",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8443,
										"name":          "https",
										"protocol":      "TCP",
									},
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "64Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "5m",
										"memory": "16Mi",
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
								},
							},
							map[string]interface{}{
								"args": []interface{}{
									"--health-probe-bind-address=:8081",
									"--metrics-bind-address=127.0.0.1:8080",
									"--leader-elect",
								},
								"command": []interface{}{
									"/manager",
								},
								"image": "quay.io/tbd-paas/identity-operator:latest",
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/healthz",
										"port": 8081,
									},
									"initialDelaySeconds": 15,
									"periodSeconds":       20,
								},
								"name": "manager",
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/readyz",
										"port": 8081,
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       10,
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"cpu":    "125m",
										"memory": "64Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "10m",
										"memory": "16Mi",
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/arch": "arm64",
							"kubernetes.io/os":   "linux",
							"tbd.io/node-type":   "platform",
						},
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
							"runAsUser":    1001,
						},
						"serviceAccountName":            "identity-operator-controller-manager",
						"terminationGracePeriodSeconds": 10,
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceIdentityOperatorControllerManager(resourceObj, parent, reconciler, req)
}
