// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package resources

import (
	"fmt"

	"github.com/elastic/elastic-package/internal/kibana"
	"github.com/elastic/go-resource"
)

type FleetAgentPolicy struct {
	// Provider is the name of the provider to use, defaults to "kibana".
	Provider string

	// Name of the policy.
	Name string

	// Description of the policy.
	Description string

	// Namespace to use for the policy.
	Namespace string

	// DataOutputID is the identifier of the output to use.
	DataOutputID string

	// Absent is set to true to indicate that the policy should not be present.
	Absent bool

	// PackagePolicies
	PackagePolicies []FleetPackagePolicy
}

type FleetPackagePolicy struct {
	// Provider is the name of the provider to use, defaults to "kibana".
	Provider string

	// Name of the policy.
	Name string

	// TemplateName is the policy template to use from the package manifest.
	TemplateName string

	// RootPath is the root of the source of the package to configure, from it we should
	// be able to read the manifest, the data stream manifests and the policy template to use.
	RootPath string

	// TODO: Config
	// Input
	// Vars

	// Absent is set to true to indicate that the policy should not be present.
	Absent bool
}

func (f *FleetPackagePolicy) String() string {
	return fmt.Sprintf("[FleetPackagePolicy:%s:%s]", f.Provider, f.Name)
}

func (f *FleetAgentPolicy) String() string {
	return fmt.Sprintf("[FleetAgentPolicy:%s:%s]", f.Provider, f.Name)
}

func (f *FleetAgentPolicy) provider(ctx resource.Context) (*KibanaProvider, error) {
	name := f.Provider
	if name == "" {
		name = DefaultKibanaProviderName
	}
	var provider *KibanaProvider
	ok := ctx.Provider(name, &provider)
	if !ok {
		return nil, fmt.Errorf("provider %q must be explicitly defined", name)
	}
	return provider, nil
}

func (f *FleetAgentPolicy) Get(ctx resource.Context) (current resource.ResourceState, err error) {
	return &FleetAgentPolicyState{expected: !f.Absent}, nil
}

func (f *FleetAgentPolicy) Create(ctx resource.Context) error {
	return nil
}

func (f *FleetAgentPolicy) Update(ctx resource.Context) error {
	return nil
}

type FleetAgentPolicyState struct {
	current  *kibana.Policy
	expected bool
}

func (s *FleetAgentPolicyState) Found() bool {
	return !s.expected || s.current != nil
}

func (s *FleetAgentPolicyState) NeedsUpdate(resource resource.Resource) (bool, error) {
	policy := resource.(*FleetAgentPolicy)
	return policy.Absent == (s.current == nil), nil
}
