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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/inftyai/llmaz/api/inference/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// BackendRuntimeConfigApplyConfiguration represents an declarative configuration of the BackendRuntimeConfig type for use
// with apply.
type BackendRuntimeConfigApplyConfiguration struct {
	Name      *v1alpha1.BackendName                   `json:"name,omitempty"`
	Version   *string                                 `json:"version,omitempty"`
	Args      []string                                `json:"args,omitempty"`
	Envs      []v1.EnvVar                             `json:"envs,omitempty"`
	Resources *ResourceRequirementsApplyConfiguration `json:"resources,omitempty"`
}

// BackendRuntimeConfigApplyConfiguration constructs an declarative configuration of the BackendRuntimeConfig type for use with
// apply.
func BackendRuntimeConfig() *BackendRuntimeConfigApplyConfiguration {
	return &BackendRuntimeConfigApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *BackendRuntimeConfigApplyConfiguration) WithName(value v1alpha1.BackendName) *BackendRuntimeConfigApplyConfiguration {
	b.Name = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *BackendRuntimeConfigApplyConfiguration) WithVersion(value string) *BackendRuntimeConfigApplyConfiguration {
	b.Version = &value
	return b
}

// WithArgs adds the given value to the Args field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Args field.
func (b *BackendRuntimeConfigApplyConfiguration) WithArgs(values ...string) *BackendRuntimeConfigApplyConfiguration {
	for i := range values {
		b.Args = append(b.Args, values[i])
	}
	return b
}

// WithEnvs adds the given value to the Envs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Envs field.
func (b *BackendRuntimeConfigApplyConfiguration) WithEnvs(values ...v1.EnvVar) *BackendRuntimeConfigApplyConfiguration {
	for i := range values {
		b.Envs = append(b.Envs, values[i])
	}
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *BackendRuntimeConfigApplyConfiguration) WithResources(value *ResourceRequirementsApplyConfiguration) *BackendRuntimeConfigApplyConfiguration {
	b.Resources = value
	return b
}
