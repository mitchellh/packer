// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package packer

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatMockProvisioner is an auto-generated flat version of MockProvisioner.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatMockProvisioner struct {
	PrepCalled       *bool         `cty:"prep_called" hcl:"prep_called"`
	PrepConfigs      []interface{} `cty:"prep_configs" hcl:"prep_configs"`
	ProvCalled       *bool         `cty:"prov_called" hcl:"prov_called"`
	ProvRetried      *bool         `cty:"prov_retried" hcl:"prov_retried"`
	ProvCommunicator Communicator  `cty:"prov_communicator" hcl:"prov_communicator"`
	ProvUi           Ui            `cty:"prov_ui" hcl:"prov_ui"`
}

// FlatMapstructure returns a new FlatMockProvisioner.
// FlatMockProvisioner is an auto-generated flat version of MockProvisioner.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*MockProvisioner) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatMockProvisioner)
}

// HCL2Spec returns the hcl spec of a MockProvisioner.
// This spec is used by HCL to read the fields of MockProvisioner.
// The decoded values from this spec will then be applied to a FlatMockProvisioner.
func (*FlatMockProvisioner) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"prep_called":       &hcldec.AttrSpec{Name: "prep_called", Type: cty.Bool, Required: false},
		"prep_configs":      &hcldec.AttrSpec{Name: "prep_configs", Type: cty.Bool, Required: false}, /* TODO(azr): could not find type */
		"prov_called":       &hcldec.AttrSpec{Name: "prov_called", Type: cty.Bool, Required: false},
		"prov_retried":      &hcldec.AttrSpec{Name: "prov_retried", Type: cty.Bool, Required: false},
		"prov_communicator": &hcldec.AttrSpec{Name: "prov_communicator", Type: cty.Bool, Required: false}, /* TODO(azr): could not find type */
		"prov_ui":           &hcldec.AttrSpec{Name: "prov_ui", Type: cty.Bool, Required: false},           /* TODO(azr): could not find type */
	}
	return s
}
