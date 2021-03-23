/*
Copyright 2021 Gravitational, Inc.

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

package types

import (
	"time"

	"github.com/gravitational/teleport/api/defaults"
	"github.com/gravitational/trace"
)

// PAMConfig represents a PAM configuration resource.
type PAMConfig interface {
	// Inherit generic resource requiements.
	Resource

	// GetEnabled checks if PAM integration should be enabled.
	GetEnabled() bool

	// SetEnabled configures whether PAM integration should be enabled or not.
	SetEnabled(bool)

	// GetServiceName returns the PAM service name to use.
	GetServiceName() string

	// SetServiceName sets the PAM service name to use.
	SetServiceName(string)

	// GetUsePAMAuth checks if PAM authentication should be used.
	GetUsePAMAuth() bool

	// SetUsePAMAuth configures whether PAM authentatication should be used.
	SetUsePAMAuth(bool)

	// GetEnvironmont fetches enviroment mappings to set for PAM modules.
	GetEnvironment() map[string]string

	// SetEnvironment sets environment mappings to set for PAM modules.
	SetEnvironment(map[string]string)

	// CheckAndSetDefaults configures the resource with default values if empty and validates invariants.
	CheckAndSetDefaults() error
}

// NewPAMConfig is a convenience function for creating a PAM configuration from a specification.
func NewPAMConfig(spec PAMConfigSpecV1) (PAMConfig, error) {
	conf := PAMConfigV1{
		Kind:    KindPAMConfig,
		Version: V1,
		Metadata: Metadata{
			Name:      MetaNamePAMConfig,
			Namespace: defaults.Namespace,
		},
		Spec: spec,
	}

	if err := conf.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	return &conf, nil
}

// DefaultPAMConfig creates a default PAM config.
func DefaultPAMConfig() (PAMConfig, error) {
	return NewPAMConfig(PAMConfigSpecV1{})
}

// PAMConfigV1 implements PAMConfig.
type PAMConfigV1 struct {
	// Kind is a resource kind - always resource.
	Kind string `json:"kind"`

	// SubKind is a resource sub kind.
	SubKind string `json:"sub_kind,omitempty"`

	// Version is a resource version.
	Version string `json:"version"`

	// Metadata is metadata about the resource.
	Metadata Metadata `json:"metadata"`

	// Spec is the specification of the resource.
	Spec PAMConfigSpecV1 `json:"spec"`
}

// GetEnabled checks if PAM integration should be enabled.
func (c *PAMConfigV1) GetEnabled() bool {
	return c.Spec.Enabled
}

// SetEnabled configures whether PAM integration should be enabled or not.
func (c *PAMConfigV1) SetEnabled(enabled bool) {
	c.Spec.Enabled = enabled
}

// GetServiceName returns the PAM service name to use.
func (c *PAMConfigV1) GetServiceName() string {
	return c.Spec.ServiceName
}

// SetServiceName sets the PAM service name to use.
func (c *PAMConfigV1) SetServiceName(serviceName string) {
	c.Spec.ServiceName = serviceName
}

// GetUsePAMAuth checks if PAM authentication should be used.
func (c *PAMConfigV1) GetUsePAMAuth() bool {
	return c.Spec.UsePAMAuth
}

// SetUsePAMAuth configures whether PAM authentatication should be used.
func (c *PAMConfigV1) SetUsePAMAuth(enabled bool) {
	c.Spec.UsePAMAuth = enabled
}

// GetEnvironmont fetches enviroment mappings to set for PAM modules.
func (c *PAMConfigV1) GetEnvironment() map[string]string {
	return c.Spec.Environment
}

// SetEnvironment sets environment mappings to set for PAM modules.
func (c *PAMConfigV1) SetEnvironment(environment map[string]string) {
	c.Spec.Environment = environment
}

// CheckAndSetDefaults configures the resource with default values if empty and validates invariants.
func (c *PAMConfigV1) CheckAndSetDefaults() error {
	// make sure we have defaults for all metadata fields
	err := c.Metadata.CheckAndSetDefaults()
	if err != nil {
		return trace.Wrap(err)
	}

	if c.Spec.ServiceName == "" {
		return trace.BadParameter("service name cannot be empty")
	}

	return nil
}

// GetVersion returns resource version.
func (c *PAMConfigV1) GetVersion() string {
	return c.Version
}

// GetName returns the name of the resource.
func (c *PAMConfigV1) GetName() string {
	return c.Metadata.Name
}

// SetName sets the name of the resource.
func (c *PAMConfigV1) SetName(e string) {
	c.Metadata.Name = e
}

// SetExpiry sets expiry time for the object.
func (c *PAMConfigV1) SetExpiry(expires time.Time) {
	c.Metadata.SetExpiry(expires)
}

// Expiry returns object expiry setting.
func (c *PAMConfigV1) Expiry() time.Time {
	return c.Metadata.Expiry()
}

// SetTTL sets Expires header using the provided clock.
// Use SetExpiry instead.
// DELETE IN 7.0.0
func (c *PAMConfigV1) SetTTL(clock Clock, ttl time.Duration) {
	c.Metadata.SetTTL(clock, ttl)
}

// GetMetadata returns object metadata.
func (c *PAMConfigV1) GetMetadata() Metadata {
	return c.Metadata
}

// GetResourceID returns resource ID.
func (c *PAMConfigV1) GetResourceID() int64 {
	return c.Metadata.ID
}

// SetResourceID sets resource ID.
func (c *PAMConfigV1) SetResourceID(id int64) {
	c.Metadata.ID = id
}

// GetKind returns resource kind.
func (c *PAMConfigV1) GetKind() string {
	return c.Kind
}

// GetSubKind returns resource subkind.
func (c *PAMConfigV1) GetSubKind() string {
	return c.SubKind
}

// SetSubKind sets resource subkind.
func (c *PAMConfigV1) SetSubKind(sk string) {
	c.SubKind = sk
}

// PAMConfigSpecV1 specifies the PAM configuration.
type PAMConfigSpecV1 struct {
	// Whether PAM integration is enabled or not. Defaults to false.
	Enabled bool `json:"enabled"`

	// The PAM service name. Defaults to `sshd`.
	ServiceName string `json:"service_name"`

	// Whether PAM authentication is enabled or not. Defaults to false.
	UsePAMAuth bool `json:"use_pam_auth"`

	// Environment mappings set for PAM modules.
	Environment map[string]string `json:"environment"`
}
