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

// CreateCRDCertmanagersCertificatesPlatformTbdIo creates the CustomResourceDefinition resource with name certmanagers.certificates.platform.tbd.io.
func CreateCRDCertmanagersCertificatesPlatformTbdIo(
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
					"controller-gen.kubebuilder.io/version": "v0.15.0",
				},
				"labels": map[string]interface{}{
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "certmanagers.certificates.platform.tbd.io",
			},
			"spec": map[string]interface{}{
				"group": "certificates.platform.tbd.io",
				"names": map[string]interface{}{
					"kind":     "CertManager",
					"listKind": "CertManagerList",
					"plural":   "certmanagers",
					"singular": "certmanager",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "CertManager is the Schema for the certmanagers API.",
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
										"description": "CertManagerSpec defines the desired state of CertManager.",
										"properties": map[string]interface{}{
											"aws": map[string]interface{}{
												"properties": map[string]interface{}{
													"roleARN": map[string]interface{}{
														"default": "",
														"description": `(Default: "")


	The AWS IAM Role ARN to use for validating public DNS records for issuing public certificates.`,
														"type": "string",
													},
												},
												"type": "object",
											},
											"controller": map[string]interface{}{
												"properties": map[string]interface{}{
													"image": map[string]interface{}{
														"default": "quay.io/jetstack/cert-manager-controller:v1.14.4",
														"description": `(Default: "quay.io/jetstack/cert-manager-controller:v1.14.4")


	Image to use for cert-manager controller deployment.`,
														"type": "string",
													},
													"replicas": map[string]interface{}{
														"default": 2,
														"description": `(Default: 2)


	Number of replicas to use for the cert-manager controller deployment.`,
														"type": "integer",
													},
													"resources": map[string]interface{}{
														"properties": map[string]interface{}{
															"limits": map[string]interface{}{
																"properties": map[string]interface{}{
																	"memory": map[string]interface{}{
																		"default": "64Mi",
																		"description": `(Default: "64Mi")


	Memory limits to use for cert-manager controller deployment.`,
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


	CPU requests to use for cert-manager controller deployment.`,
																		"type": "string",
																	},
																	"memory": map[string]interface{}{
																		"default": "32Mi",
																		"description": `(Default: "32Mi")


	Memory requests to use for cert-manager controller deployment.`,
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
											"injector": map[string]interface{}{
												"properties": map[string]interface{}{
													"image": map[string]interface{}{
														"default": "quay.io/jetstack/cert-manager-cainjector:v1.14.4",
														"description": `(Default: "quay.io/jetstack/cert-manager-cainjector:v1.14.4")


	Image to use for cert-manager CA injector deployment.`,
														"type": "string",
													},
													"replicas": map[string]interface{}{
														"default": 2,
														"description": `(Default: 2)


	Number of replicas to use for the cert-manager cainjector deployment.`,
														"type": "integer",
													},
													"resources": map[string]interface{}{
														"properties": map[string]interface{}{
															"limits": map[string]interface{}{
																"properties": map[string]interface{}{
																	"memory": map[string]interface{}{
																		"default": "128Mi",
																		"description": `(Default: "128Mi")


	Memory limits to use for cert-manager CA injector deployment.`,
																		"type": "string",
																	},
																},
																"type": "object",
															},
															"requests": map[string]interface{}{
																"properties": map[string]interface{}{
																	"cpu": map[string]interface{}{
																		"default": "50m",
																		"description": `(Default: "50m")


	CPU requests to use for cert-manager CA injector deployment.`,
																		"type": "string",
																	},
																	"memory": map[string]interface{}{
																		"default": "64Mi",
																		"description": `(Default: "64Mi")


	Memory requests to use for cert-manager CA injector deployment.`,
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
											"namespace": map[string]interface{}{
												"default":     "tbd-certificates-system",
												"description": "(Default: \"tbd-certificates-system\")",
												"type":        "string",
											},
											"webhook": map[string]interface{}{
												"properties": map[string]interface{}{
													"image": map[string]interface{}{
														"default": "quay.io/jetstack/cert-manager-webhook:v1.14.4",
														"description": `(Default: "quay.io/jetstack/cert-manager-webhook:v1.14.4")


	Image to use for cert-manager webhook deployment.`,
														"type": "string",
													},
													"replicas": map[string]interface{}{
														"default": 2,
														"description": `(Default: 2)


	Number of replicas to use for the cert-manager webhook deployment.`,
														"type": "integer",
													},
													"resources": map[string]interface{}{
														"properties": map[string]interface{}{
															"limits": map[string]interface{}{
																"properties": map[string]interface{}{
																	"memory": map[string]interface{}{
																		"default": "64Mi",
																		"description": `(Default: "64Mi")


	Memory limits to use for cert-manager webhook deployment.`,
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


	CPU requests to use for cert-manager webhook deployment.`,
																		"type": "string",
																	},
																	"memory": map[string]interface{}{
																		"default": "32Mi",
																		"description": `(Default: "32Mi")


	Memory requests to use for cert-manager webhook deployment.`,
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
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "CertManagerStatus defines the observed state of CertManager.",
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

	return mutate.MutateCRDCertmanagersCertificatesPlatformTbdIo(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDTrustmanagersCertificatesPlatformTbdIo creates the CustomResourceDefinition resource with name trustmanagers.certificates.platform.tbd.io.
func CreateCRDTrustmanagersCertificatesPlatformTbdIo(
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
					"controller-gen.kubebuilder.io/version": "v0.15.0",
				},
				"labels": map[string]interface{}{
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "trustmanagers.certificates.platform.tbd.io",
			},
			"spec": map[string]interface{}{
				"group": "certificates.platform.tbd.io",
				"names": map[string]interface{}{
					"kind":     "TrustManager",
					"listKind": "TrustManagerList",
					"plural":   "trustmanagers",
					"singular": "trustmanager",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "TrustManager is the Schema for the trustmanagers API.",
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
										"description": "TrustManagerSpec defines the desired state of TrustManager.",
										"properties": map[string]interface{}{
											"controller": map[string]interface{}{
												"properties": map[string]interface{}{
													"image": map[string]interface{}{
														"default": "quay.io/jetstack/trust-manager:v0.9.2",
														"description": `(Default: "quay.io/jetstack/trust-manager:v0.9.2")


	Image to use for trust-manager controller deployment.`,
														"type": "string",
													},
													"replicas": map[string]interface{}{
														"default": 2,
														"description": `(Default: 2)


	Number of replicas to use for the trust-manager controller deployment.`,
														"type": "integer",
													},
													"resources": map[string]interface{}{
														"properties": map[string]interface{}{
															"limits": map[string]interface{}{
																"properties": map[string]interface{}{
																	"memory": map[string]interface{}{
																		"default": "64Mi",
																		"description": `(Default: "64Mi")


	Memory limits to use for trust-manager controller deployment.`,
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


	CPU requests to use for trust-manager controller deployment.`,
																		"type": "string",
																	},
																	"memory": map[string]interface{}{
																		"default": "32Mi",
																		"description": `(Default: "32Mi")


	Memory requests to use for trust-manager controller deployment.`,
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
											"namespace": map[string]interface{}{
												"default":     "tbd-certificates-system",
												"description": "(Default: \"tbd-certificates-system\")",
												"type":        "string",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "TrustManagerStatus defines the observed state of TrustManager.",
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

	return mutate.MutateCRDTrustmanagersCertificatesPlatformTbdIo(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceCertificatesOperatorControllerManager creates the ServiceAccount resource with name certificates-operator-controller-manager.
func CreateServiceAccountNamespaceCertificatesOperatorControllerManager(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name":      "certificates-operator-controller-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceCertificatesOperatorControllerManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateRoleNamespaceCertificatesOperatorLeaderElectionRole creates the Role resource with name certificates-operator-leader-election-role.
func CreateRoleNamespaceCertificatesOperatorLeaderElectionRole(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name":      "certificates-operator-leader-election-role",
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

	return mutate.MutateRoleNamespaceCertificatesOperatorLeaderElectionRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=certmanagers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=certmanagers/status,verbs=get

// CreateClusterRoleCertificatesOperatorCertificatesCertmanagerEditorRole creates the ClusterRole resource with name certificates-operator-certificates-certmanager-editor-role.
func CreateClusterRoleCertificatesOperatorCertificatesCertmanagerEditorRole(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "certificates-operator-certificates-certmanager-editor-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"certmanagers",
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
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"certmanagers/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertificatesOperatorCertificatesCertmanagerEditorRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=certmanagers,verbs=get;list;watch
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=certmanagers/status,verbs=get

// CreateClusterRoleCertificatesOperatorCertificatesCertmanagerViewerRole creates the ClusterRole resource with name certificates-operator-certificates-certmanager-viewer-role.
func CreateClusterRoleCertificatesOperatorCertificatesCertmanagerViewerRole(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "certificates-operator-certificates-certmanager-viewer-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"certmanagers",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"certmanagers/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertificatesOperatorCertificatesCertmanagerViewerRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=trustmanagers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=trustmanagers/status,verbs=get

// CreateClusterRoleCertificatesOperatorCertificatesTrustmanagerEditorRole creates the ClusterRole resource with name certificates-operator-certificates-trustmanager-editor-role.
func CreateClusterRoleCertificatesOperatorCertificatesTrustmanagerEditorRole(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "certificates-operator-certificates-trustmanager-editor-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"trustmanagers",
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
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"trustmanagers/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertificatesOperatorCertificatesTrustmanagerEditorRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=trustmanagers,verbs=get;list;watch
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=trustmanagers/status,verbs=get

// CreateClusterRoleCertificatesOperatorCertificatesTrustmanagerViewerRole creates the ClusterRole resource with name certificates-operator-certificates-trustmanager-viewer-role.
func CreateClusterRoleCertificatesOperatorCertificatesTrustmanagerViewerRole(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "certificates-operator-certificates-trustmanager-viewer-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"trustmanagers",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"trustmanagers/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertificatesOperatorCertificatesTrustmanagerViewerRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges/finalizers,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges/status,verbs=patch;update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders/finalizers,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders/status,verbs=patch;update
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests/finalizers,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests/status,verbs=patch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates/finalizers,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates/status,verbs=patch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers/status,verbs=patch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers/status,verbs=patch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=signers,verbs=approve
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=get;list;update;watch
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests/status,verbs=patch;update
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=sign
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=certmanagers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=certmanagers/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=trustmanagers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=certificates.platform.tbd.io,resources=trustmanagers/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;get;patch;update
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=create;delete;get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gateways,verbs=get;list;watch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gateways/finalizers,verbs=update
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=httproutes,verbs=create;delete;get;list;update;watch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=httproutes/finalizers,verbs=update
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=create;delete;get;list;update;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/finalizers,verbs=update
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=route.openshift.io,resources=routes/custom-host,verbs=create
// +kubebuilder:rbac:groups=trust.cert-manager.io,resources=bundles,verbs=get;list;update;watch
// +kubebuilder:rbac:groups=trust.cert-manager.io,resources=bundles/finalizers,verbs=update
// +kubebuilder:rbac:groups=trust.cert-manager.io,resources=bundles/status,verbs=patch

// CreateClusterRoleCertificatesOperatorManagerRole creates the ClusterRole resource with name certificates-operator-manager-role.
func CreateClusterRoleCertificatesOperatorManagerRole(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "certificates-operator-manager-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"orders",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"orders/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"orders/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
					},
				},
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
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"validatingwebhookconfigurations",
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
						"apiextensions.k8s.io",
					},
					"resources": []interface{}{
						"customresourcedefinitions",
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
						"apiregistration.k8s.io",
					},
					"resources": []interface{}{
						"apiservices",
					},
					"verbs": []interface{}{
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
						"authorization.k8s.io",
					},
					"resources": []interface{}{
						"subjectaccessreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificaterequests",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
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
						"certificaterequests/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificaterequests/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
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
						"deletecollection",
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
						"certificates/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"clusterissuers",
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
						"clusterissuers/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"issuers",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
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
						"issuers/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"signers",
					},
					"verbs": []interface{}{
						"approve",
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
						"get",
						"list",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.k8s.io",
					},
					"resources": []interface{}{
						"certificatesigningrequests/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.k8s.io",
					},
					"resources": []interface{}{
						"signers",
					},
					"verbs": []interface{}{
						"sign",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"certmanagers",
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
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"certmanagers/status",
					},
					"verbs": []interface{}{
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"trustmanagers",
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
						"certificates.platform.tbd.io",
					},
					"resources": []interface{}{
						"trustmanagers/status",
					},
					"verbs": []interface{}{
						"get",
						"patch",
						"update",
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
						"create",
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
						"configmaps",
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
						"endpoints",
					},
					"verbs": []interface{}{
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
						"pods",
					},
					"verbs": []interface{}{
						"create",
						"delete",
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
						"secrets",
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
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"gateways",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"gateways/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"httproutes",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"httproutes/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses/finalizers",
					},
					"verbs": []interface{}{
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
				map[string]interface{}{
					"apiGroups": []interface{}{
						"route.openshift.io",
					},
					"resources": []interface{}{
						"routes/custom-host",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"trust.cert-manager.io",
					},
					"resources": []interface{}{
						"bundles",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"trust.cert-manager.io",
					},
					"resources": []interface{}{
						"bundles/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"trust.cert-manager.io",
					},
					"resources": []interface{}{
						"bundles/status",
					},
					"verbs": []interface{}{
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertificatesOperatorManagerRole(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceCertificatesOperatorLeaderElectionRolebinding creates the RoleBinding resource with name certificates-operator-leader-election-rolebinding.
func CreateRoleBindingNamespaceCertificatesOperatorLeaderElectionRolebinding(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name":      "certificates-operator-leader-election-rolebinding",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "certificates-operator-leader-election-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "certificates-operator-controller-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceCertificatesOperatorLeaderElectionRolebinding(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertificatesOperatorManagerRolebinding creates the ClusterRoleBinding resource with name certificates-operator-manager-rolebinding.
func CreateClusterRoleBindingCertificatesOperatorManagerRolebinding(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
				},
				"name": "certificates-operator-manager-rolebinding",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "certificates-operator-manager-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "certificates-operator-controller-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertificatesOperatorManagerRolebinding(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCertificatesOperatorControllerManager creates the Deployment resource with name certificates-operator-controller-manager.
func CreateDeploymentNamespaceCertificatesOperatorControllerManager(
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
					"app":                                  "certificates-operator",
					"app.kubernetes.io/component":          "certificates-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-config-operator",
					"app.kubernetes.io/name":               "certificates-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "unstable",
					"capabilities.tbd.io/capability":       "platform-config",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
					"control-plane":                        "controller-manager",
				},
				"name":      "certificates-operator-controller-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app":                                  "certificates-operator",
						"app.kubernetes.io/component":          "certificates-operator",
						"app.kubernetes.io/instance":           "manager",
						"app.kubernetes.io/managed-by":         "platform-config-operator",
						"app.kubernetes.io/name":               "certificates-operator",
						"app.kubernetes.io/part-of":            "platform",
						"app.kubernetes.io/version":            "unstable",
						"capabilities.tbd.io/capability":       "certificates",
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
							"app":                                  "certificates-operator",
							"app.kubernetes.io/component":          "certificates-operator",
							"app.kubernetes.io/instance":           "manager",
							"app.kubernetes.io/managed-by":         "platform-config-operator",
							"app.kubernetes.io/name":               "certificates-operator",
							"app.kubernetes.io/part-of":            "platform",
							"app.kubernetes.io/version":            "unstable",
							"capabilities.tbd.io/capability":       "certificates",
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
															"certificates-operator",
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
								"image": "quay.io/tbd-paas/certificates-operator:v0.0.0-alpha.2",
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
							"kubernetes.io/os": "linux",
							"tbd.io/node-type": "platform",
						},
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
							"runAsUser":    1001,
						},
						"serviceAccountName":            "certificates-operator-controller-manager",
						"terminationGracePeriodSeconds": 10,
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceCertificatesOperatorControllerManager(resourceObj, parent, reconciler, req)
}
