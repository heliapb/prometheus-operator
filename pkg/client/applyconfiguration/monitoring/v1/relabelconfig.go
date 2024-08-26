// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

// RelabelConfigApplyConfiguration represents a declarative configuration of the RelabelConfig type for use
// with apply.
type RelabelConfigApplyConfiguration struct {
	SourceLabels []v1.LabelName `json:"sourceLabels,omitempty"`
	Separator    *string        `json:"separator,omitempty"`
	TargetLabel  *string        `json:"targetLabel,omitempty"`
	Regex        *string        `json:"regex,omitempty"`
	Modulus      *uint64        `json:"modulus,omitempty"`
	Replacement  *string        `json:"replacement,omitempty"`
	Action       *string        `json:"action,omitempty"`
}

// RelabelConfigApplyConfiguration constructs a declarative configuration of the RelabelConfig type for use with
// apply.
func RelabelConfig() *RelabelConfigApplyConfiguration {
	return &RelabelConfigApplyConfiguration{}
}

// WithSourceLabels adds the given value to the SourceLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SourceLabels field.
func (b *RelabelConfigApplyConfiguration) WithSourceLabels(values ...v1.LabelName) *RelabelConfigApplyConfiguration {
	for i := range values {
		b.SourceLabels = append(b.SourceLabels, values[i])
	}
	return b
}

// WithSeparator sets the Separator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Separator field is set to the value of the last call.
func (b *RelabelConfigApplyConfiguration) WithSeparator(value string) *RelabelConfigApplyConfiguration {
	b.Separator = &value
	return b
}

// WithTargetLabel sets the TargetLabel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetLabel field is set to the value of the last call.
func (b *RelabelConfigApplyConfiguration) WithTargetLabel(value string) *RelabelConfigApplyConfiguration {
	b.TargetLabel = &value
	return b
}

// WithRegex sets the Regex field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Regex field is set to the value of the last call.
func (b *RelabelConfigApplyConfiguration) WithRegex(value string) *RelabelConfigApplyConfiguration {
	b.Regex = &value
	return b
}

// WithModulus sets the Modulus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Modulus field is set to the value of the last call.
func (b *RelabelConfigApplyConfiguration) WithModulus(value uint64) *RelabelConfigApplyConfiguration {
	b.Modulus = &value
	return b
}

// WithReplacement sets the Replacement field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replacement field is set to the value of the last call.
func (b *RelabelConfigApplyConfiguration) WithReplacement(value string) *RelabelConfigApplyConfiguration {
	b.Replacement = &value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *RelabelConfigApplyConfiguration) WithAction(value string) *RelabelConfigApplyConfiguration {
	b.Action = &value
	return b
}
