// +build !ignore_autogenerated

/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.Condition":                         schema_pkg_apis_svcneg_v1beta1_Condition(ref),
		"k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.NegObjectReference":                schema_pkg_apis_svcneg_v1beta1_NegObjectReference(ref),
		"k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.ServiceNetworkEndpointGroup":       schema_pkg_apis_svcneg_v1beta1_ServiceNetworkEndpointGroup(ref),
		"k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.ServiceNetworkEndpointGroupStatus": schema_pkg_apis_svcneg_v1beta1_ServiceNetworkEndpointGroupStatus(ref),
	}
}

func schema_pkg_apis_svcneg_v1beta1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NegCondition contains details for the current condition of this NEG.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of the condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "ObservedGeneration will not be set for ServiceNetworkEndpointGroup as the spec is empty.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "The reason for the condition's last transition",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition. This field may be empty.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status", "lastTransitionTime", "reason", "message"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_svcneg_v1beta1_NegObjectReference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NegObjectReference is the object reference to the NEG resource in GCE",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "The unique identifier for the NEG resource in GCE API.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"selfLink": {
						SchemaProps: spec.SchemaProps{
							Description: "SelfLink is the GCE Server-defined fully-qualified URL for the GCE NEG resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"networkEndpointType": {
						SchemaProps: spec.SchemaProps{
							Description: "NetworkEndpointType: Type of network endpoints in this network endpoint group.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"id"},
			},
		},
	}
}

func schema_pkg_apis_svcneg_v1beta1_ServiceNetworkEndpointGroup(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.ServiceNetworkEndpointGroupSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.ServiceNetworkEndpointGroupStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.ServiceNetworkEndpointGroupSpec", "k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.ServiceNetworkEndpointGroupStatus"},
	}
}

func schema_pkg_apis_svcneg_v1beta1_ServiceNetworkEndpointGroupStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceNetworkEndpointGroupStatus is the status for a ServiceNetworkEndpointGroup resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"networkEndpointGroups": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"id",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.NegObjectReference"),
									},
								},
							},
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Last time the NEG syncer syncs associated NEGs.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.Condition"),
									},
								},
							},
						},
					},
					"lastSyncTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the NEG syncer syncs associated NEGs.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time", "k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.Condition", "k8s.io/ingress-gce/pkg/apis/svcneg/v1beta1.NegObjectReference"},
	}
}
