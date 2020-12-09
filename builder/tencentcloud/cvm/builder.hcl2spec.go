// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package cvm

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string                    `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string                    `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string                    `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool                      `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool                      `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string                    `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string          `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string                   `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	SecretId                  *string                    `mapstructure:"secret_id" required:"true" cty:"secret_id" hcl:"secret_id"`
	SecretKey                 *string                    `mapstructure:"secret_key" required:"true" cty:"secret_key" hcl:"secret_key"`
	Region                    *string                    `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	Zone                      *string                    `mapstructure:"zone" required:"true" cty:"zone" hcl:"zone"`
	SkipValidation            *bool                      `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation" hcl:"skip_region_validation"`
	ImageName                 *string                    `mapstructure:"image_name" required:"true" cty:"image_name" hcl:"image_name"`
	ImageDescription          *string                    `mapstructure:"image_description" required:"false" cty:"image_description" hcl:"image_description"`
	Reboot                    *bool                      `mapstructure:"reboot" required:"false" cty:"reboot" hcl:"reboot"`
	ForcePoweroff             *bool                      `mapstructure:"force_poweroff" required:"false" cty:"force_poweroff" hcl:"force_poweroff"`
	Sysprep                   *bool                      `mapstructure:"sysprep" required:"false" cty:"sysprep" hcl:"sysprep"`
	ImageForceDelete          *bool                      `mapstructure:"image_force_delete" cty:"image_force_delete" hcl:"image_force_delete"`
	ImageCopyRegions          []string                   `mapstructure:"image_copy_regions" required:"false" cty:"image_copy_regions" hcl:"image_copy_regions"`
	ImageShareAccounts        []string                   `mapstructure:"image_share_accounts" required:"false" cty:"image_share_accounts" hcl:"image_share_accounts"`
	AssociatePublicIpAddress  *bool                      `mapstructure:"associate_public_ip_address" required:"false" cty:"associate_public_ip_address" hcl:"associate_public_ip_address"`
	SourceImageId             *string                    `mapstructure:"source_image_id" required:"false" cty:"source_image_id" hcl:"source_image_id"`
	SourceImageName           *string                    `mapstructure:"source_image_name" required:"false" cty:"source_image_name" hcl:"source_image_name"`
	InstanceType              *string                    `mapstructure:"instance_type" required:"true" cty:"instance_type" hcl:"instance_type"`
	InstanceName              *string                    `mapstructure:"instance_name" required:"false" cty:"instance_name" hcl:"instance_name"`
	DiskType                  *string                    `mapstructure:"disk_type" required:"false" cty:"disk_type" hcl:"disk_type"`
	DiskSize                  *int64                     `mapstructure:"disk_size" required:"false" cty:"disk_size" hcl:"disk_size"`
	DataDisks                 []FlattencentCloudDataDisk `mapstructure:"data_disks" cty:"data_disks" hcl:"data_disks"`
	VpcId                     *string                    `mapstructure:"vpc_id" required:"false" cty:"vpc_id" hcl:"vpc_id"`
	VpcName                   *string                    `mapstructure:"vpc_name" required:"false" cty:"vpc_name" hcl:"vpc_name"`
	VpcIp                     *string                    `mapstructure:"vpc_ip" cty:"vpc_ip" hcl:"vpc_ip"`
	SubnetId                  *string                    `mapstructure:"subnet_id" required:"false" cty:"subnet_id" hcl:"subnet_id"`
	SubnetName                *string                    `mapstructure:"subnet_name" required:"false" cty:"subnet_name" hcl:"subnet_name"`
	CidrBlock                 *string                    `mapstructure:"cidr_block" required:"false" cty:"cidr_block" hcl:"cidr_block"`
	SubnectCidrBlock          *string                    `mapstructure:"subnect_cidr_block" required:"false" cty:"subnect_cidr_block" hcl:"subnect_cidr_block"`
	InternetChargeType        *string                    `mapstructure:"internet_charge_type" cty:"internet_charge_type" hcl:"internet_charge_type"`
	InternetMaxBandwidthOut   *int64                     `mapstructure:"internet_max_bandwidth_out" required:"false" cty:"internet_max_bandwidth_out" hcl:"internet_max_bandwidth_out"`
	SecurityGroupId           *string                    `mapstructure:"security_group_id" required:"false" cty:"security_group_id" hcl:"security_group_id"`
	SecurityGroupName         *string                    `mapstructure:"security_group_name" required:"false" cty:"security_group_name" hcl:"security_group_name"`
	UserData                  *string                    `mapstructure:"user_data" required:"false" cty:"user_data" hcl:"user_data"`
	UserDataFile              *string                    `mapstructure:"user_data_file" required:"false" cty:"user_data_file" hcl:"user_data_file"`
	HostName                  *string                    `mapstructure:"host_name" required:"false" cty:"host_name" hcl:"host_name"`
	RunTags                   map[string]string          `mapstructure:"run_tags" required:"false" cty:"run_tags" hcl:"run_tags"`
	RunTag                    []config.FlatKeyValue      `mapstructure:"run_tag" required:"false" cty:"run_tag" hcl:"run_tag"`
	Type                      *string                    `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string                    `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string                    `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                   *int                       `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername               *string                    `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword               *string                    `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName            *string                    `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string                    `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType   *string                    `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits   *int                       `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                []string                   `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool                      `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos               []string                   `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile         *string                    `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile        *string                    `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                    *bool                      `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string                    `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string                    `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool                      `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool                      `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int                       `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string                    `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int                       `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool                      `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string                    `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string                    `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool                      `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string                    `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile *string                    `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod     *string                    `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string                    `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int                       `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string                    `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string                    `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string                    `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string                    `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string                   `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string                   `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte                     `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte                     `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string                    `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string                    `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string                    `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool                      `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int                       `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string                    `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool                      `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool                      `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool                      `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	SSHPrivateIp              *bool                      `mapstructure:"ssh_private_ip" cty:"ssh_private_ip" hcl:"ssh_private_ip"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":          &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"secret_id":                    &hcldec.AttrSpec{Name: "secret_id", Type: cty.String, Required: false},
		"secret_key":                   &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"region":                       &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"zone":                         &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"skip_region_validation":       &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"image_name":                   &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":            &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"reboot":                       &hcldec.AttrSpec{Name: "reboot", Type: cty.Bool, Required: false},
		"force_poweroff":               &hcldec.AttrSpec{Name: "force_poweroff", Type: cty.Bool, Required: false},
		"sysprep":                      &hcldec.AttrSpec{Name: "sysprep", Type: cty.Bool, Required: false},
		"image_force_delete":           &hcldec.AttrSpec{Name: "image_force_delete", Type: cty.Bool, Required: false},
		"image_copy_regions":           &hcldec.AttrSpec{Name: "image_copy_regions", Type: cty.List(cty.String), Required: false},
		"image_share_accounts":         &hcldec.AttrSpec{Name: "image_share_accounts", Type: cty.List(cty.String), Required: false},
		"associate_public_ip_address":  &hcldec.AttrSpec{Name: "associate_public_ip_address", Type: cty.Bool, Required: false},
		"source_image_id":              &hcldec.AttrSpec{Name: "source_image_id", Type: cty.String, Required: false},
		"source_image_name":            &hcldec.AttrSpec{Name: "source_image_name", Type: cty.String, Required: false},
		"instance_type":                &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"instance_name":                &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"disk_type":                    &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"data_disks":                   &hcldec.BlockListSpec{TypeName: "data_disks", Nested: hcldec.ObjectSpec((*FlattencentCloudDataDisk)(nil).HCL2Spec())},
		"vpc_id":                       &hcldec.AttrSpec{Name: "vpc_id", Type: cty.String, Required: false},
		"vpc_name":                     &hcldec.AttrSpec{Name: "vpc_name", Type: cty.String, Required: false},
		"vpc_ip":                       &hcldec.AttrSpec{Name: "vpc_ip", Type: cty.String, Required: false},
		"subnet_id":                    &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"subnet_name":                  &hcldec.AttrSpec{Name: "subnet_name", Type: cty.String, Required: false},
		"cidr_block":                   &hcldec.AttrSpec{Name: "cidr_block", Type: cty.String, Required: false},
		"subnect_cidr_block":           &hcldec.AttrSpec{Name: "subnect_cidr_block", Type: cty.String, Required: false},
		"internet_charge_type":         &hcldec.AttrSpec{Name: "internet_charge_type", Type: cty.String, Required: false},
		"internet_max_bandwidth_out":   &hcldec.AttrSpec{Name: "internet_max_bandwidth_out", Type: cty.Number, Required: false},
		"security_group_id":            &hcldec.AttrSpec{Name: "security_group_id", Type: cty.String, Required: false},
		"security_group_name":          &hcldec.AttrSpec{Name: "security_group_name", Type: cty.String, Required: false},
		"user_data":                    &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":               &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"host_name":                    &hcldec.AttrSpec{Name: "host_name", Type: cty.String, Required: false},
		"run_tags":                     &hcldec.AttrSpec{Name: "run_tags", Type: cty.Map(cty.String), Required: false},
		"run_tag":                      &hcldec.BlockListSpec{TypeName: "run_tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":      &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":      &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":  &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":         &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file": &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"ssh_private_ip":               &hcldec.AttrSpec{Name: "ssh_private_ip", Type: cty.Bool, Required: false},
	}
	return s
}
