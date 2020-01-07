// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package ebs

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/builder/amazon/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName                           *string                                `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType                         *string                                `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug                               *bool                                  `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce                               *bool                                  `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError                             *string                                `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars                            map[string]string                      `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars                       []string                               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	AccessKey                                 *string                                `mapstructure:"access_key" required:"true" cty:"access_key"`
	CustomEndpointEc2                         *string                                `mapstructure:"custom_endpoint_ec2" required:"false" cty:"custom_endpoint_ec2"`
	DecodeAuthZMessages                       *bool                                  `mapstructure:"decode_authorization_messages" required:"false" cty:"decode_authorization_messages"`
	InsecureSkipTLSVerify                     *bool                                  `mapstructure:"insecure_skip_tls_verify" required:"false" cty:"insecure_skip_tls_verify"`
	MFACode                                   *string                                `mapstructure:"mfa_code" required:"false" cty:"mfa_code"`
	ProfileName                               *string                                `mapstructure:"profile" required:"false" cty:"profile"`
	RawRegion                                 *string                                `mapstructure:"region" required:"true" cty:"region"`
	SecretKey                                 *string                                `mapstructure:"secret_key" required:"true" cty:"secret_key"`
	SkipValidation                            *bool                                  `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation"`
	SkipMetadataApiCheck                      *bool                                  `mapstructure:"skip_metadata_api_check" cty:"skip_metadata_api_check"`
	Token                                     *string                                `mapstructure:"token" required:"false" cty:"token"`
	VaultAWSEngine                            *common.FlatVaultAWSEngineOptions      `mapstructure:"vault_aws_engine" required:"false" cty:"vault_aws_engine"`
	AMIName                                   *string                                `mapstructure:"ami_name" required:"true" cty:"ami_name"`
	AMIDescription                            *string                                `mapstructure:"ami_description" required:"false" cty:"ami_description"`
	AMIVirtType                               *string                                `mapstructure:"ami_virtualization_type" required:"false" cty:"ami_virtualization_type"`
	AMIUsers                                  []string                               `mapstructure:"ami_users" required:"false" cty:"ami_users"`
	AMIGroups                                 []string                               `mapstructure:"ami_groups" required:"false" cty:"ami_groups"`
	AMIProductCodes                           []string                               `mapstructure:"ami_product_codes" required:"false" cty:"ami_product_codes"`
	AMIRegions                                []string                               `mapstructure:"ami_regions" required:"false" cty:"ami_regions"`
	AMITags                                   common.TagMap                          `mapstructure:"tags" required:"false" cty:"tags"`
	AMIENASupport                             *bool                                  `mapstructure:"ena_support" required:"false" cty:"ena_support"`
	AMISriovNetSupport                        *bool                                  `mapstructure:"sriov_support" required:"false" cty:"sriov_support"`
	AMIForceDeregister                        *bool                                  `mapstructure:"force_deregister" required:"false" cty:"force_deregister"`
	AMIForceDeleteSnapshot                    *bool                                  `mapstructure:"force_delete_snapshot" required:"false" cty:"force_delete_snapshot"`
	AMIEncryptBootVolume                      *bool                                  `mapstructure:"encrypt_boot" required:"false" cty:"encrypt_boot"`
	AMIKmsKeyId                               *string                                `mapstructure:"kms_key_id" required:"false" cty:"kms_key_id"`
	AMIRegionKMSKeyIDs                        map[string]string                      `mapstructure:"region_kms_key_ids" required:"false" cty:"region_kms_key_ids"`
	AMISkipBuildRegion                        *bool                                  `mapstructure:"skip_save_build_region" cty:"skip_save_build_region"`
	SnapshotTags                              common.TagMap                          `mapstructure:"snapshot_tags" required:"false" cty:"snapshot_tags"`
	SnapshotUsers                             []string                               `mapstructure:"snapshot_users" required:"false" cty:"snapshot_users"`
	SnapshotGroups                            []string                               `mapstructure:"snapshot_groups" required:"false" cty:"snapshot_groups"`
	AssociatePublicIpAddress                  *bool                                  `mapstructure:"associate_public_ip_address" required:"false" cty:"associate_public_ip_address"`
	AvailabilityZone                          *string                                `mapstructure:"availability_zone" required:"false" cty:"availability_zone"`
	BlockDurationMinutes                      *int64                                 `mapstructure:"block_duration_minutes" required:"false" cty:"block_duration_minutes"`
	DisableStopInstance                       *bool                                  `mapstructure:"disable_stop_instance" required:"false" cty:"disable_stop_instance"`
	EbsOptimized                              *bool                                  `mapstructure:"ebs_optimized" required:"false" cty:"ebs_optimized"`
	EnableT2Unlimited                         *bool                                  `mapstructure:"enable_t2_unlimited" required:"false" cty:"enable_t2_unlimited"`
	IamInstanceProfile                        *string                                `mapstructure:"iam_instance_profile" required:"false" cty:"iam_instance_profile"`
	SkipProfileValidation                     *bool                                  `mapstructure:"skip_profile_validation" required:"false" cty:"skip_profile_validation"`
	TemporaryIamInstanceProfilePolicyDocument *common.FlatPolicyDocument             `mapstructure:"temporary_iam_instance_profile_policy_document" required:"false" cty:"temporary_iam_instance_profile_policy_document"`
	InstanceInitiatedShutdownBehavior         *string                                `mapstructure:"shutdown_behavior" required:"false" cty:"shutdown_behavior"`
	InstanceType                              *string                                `mapstructure:"instance_type" required:"true" cty:"instance_type"`
	SecurityGroupFilter                       *common.FlatSecurityGroupFilterOptions `mapstructure:"security_group_filter" required:"false" cty:"security_group_filter"`
	RunTags                                   map[string]string                      `mapstructure:"run_tags" required:"false" cty:"run_tags"`
	SecurityGroupId                           *string                                `mapstructure:"security_group_id" required:"false" cty:"security_group_id"`
	SecurityGroupIds                          []string                               `mapstructure:"security_group_ids" required:"false" cty:"security_group_ids"`
	SourceAmi                                 *string                                `mapstructure:"source_ami" required:"true" cty:"source_ami"`
	SourceAmiFilter                           *common.FlatAmiFilterOptions           `mapstructure:"source_ami_filter" required:"false" cty:"source_ami_filter"`
	SpotInstanceTypes                         []string                               `mapstructure:"spot_instance_types" required:"false" cty:"spot_instance_types"`
	SpotPrice                                 *string                                `mapstructure:"spot_price" required:"false" cty:"spot_price"`
	SpotPriceAutoProduct                      *string                                `mapstructure:"spot_price_auto_product" required:"false" cty:"spot_price_auto_product"`
	SpotTags                                  map[string]string                      `mapstructure:"spot_tags" required:"false" cty:"spot_tags"`
	SubnetFilter                              *common.FlatSubnetFilterOptions        `mapstructure:"subnet_filter" required:"false" cty:"subnet_filter"`
	SubnetId                                  *string                                `mapstructure:"subnet_id" required:"false" cty:"subnet_id"`
	TemporaryKeyPairName                      *string                                `mapstructure:"temporary_key_pair_name" required:"false" cty:"temporary_key_pair_name"`
	TemporarySGSourceCidrs                    []string                               `mapstructure:"temporary_security_group_source_cidrs" required:"false" cty:"temporary_security_group_source_cidrs"`
	UserData                                  *string                                `mapstructure:"user_data" required:"false" cty:"user_data"`
	UserDataFile                              *string                                `mapstructure:"user_data_file" required:"false" cty:"user_data_file"`
	VpcFilter                                 *common.FlatVpcFilterOptions           `mapstructure:"vpc_filter" required:"false" cty:"vpc_filter"`
	VpcId                                     *string                                `mapstructure:"vpc_id" required:"false" cty:"vpc_id"`
	WindowsPasswordTimeout                    *string                                `mapstructure:"windows_password_timeout" required:"false" cty:"windows_password_timeout"`
	Type                                      *string                                `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect                        *string                                `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                                   *string                                `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                                   *int                                   `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername                               *string                                `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword                               *string                                `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName                            *string                                `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHClearAuthorizedKeys                    *bool                                  `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile                         *string                                `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                                    *bool                                  `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                                *string                                `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth                              *bool                                  `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding                 *bool                                  `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts                      *int                                   `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost                            *string                                `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort                            *int                                   `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth                       *bool                                  `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername                        *string                                `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword                        *string                                `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile                  *string                                `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod                     *string                                `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost                              *string                                `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort                              *int                                   `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername                          *string                                `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword                          *string                                `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval                      *string                                `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout                       *string                                `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels                          []string                               `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels                           []string                               `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey                              []byte                                 `mapstructure:"ssh_public_key" cty:"ssh_public_key"`
	SSHPrivateKey                             []byte                                 `mapstructure:"ssh_private_key" cty:"ssh_private_key"`
	WinRMUser                                 *string                                `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword                             *string                                `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                                 *string                                `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                                 *int                                   `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout                              *string                                `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL                               *bool                                  `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure                             *bool                                  `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM                              *bool                                  `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	SSHInterface                              *string                                `mapstructure:"ssh_interface" cty:"ssh_interface"`
	AMIMappings                               []common.FlatBlockDevice               `mapstructure:"ami_block_device_mappings" required:"false" cty:"ami_block_device_mappings"`
	LaunchMappings                            []common.FlatBlockDevice               `mapstructure:"launch_block_device_mappings" required:"false" cty:"launch_block_device_mappings"`
	VolumeRunTags                             common.TagMap                          `mapstructure:"run_volume_tags" cty:"run_volume_tags"`
	NoEphemeral                               *bool                                  `mapstructure:"no_ephemeral" required:"false" cty:"no_ephemeral"`
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
		"packer_build_name":             &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":           &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                  &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                  &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":               &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":         &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":    &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"access_key":                    &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"custom_endpoint_ec2":           &hcldec.AttrSpec{Name: "custom_endpoint_ec2", Type: cty.String, Required: false},
		"decode_authorization_messages": &hcldec.AttrSpec{Name: "decode_authorization_messages", Type: cty.Bool, Required: false},
		"insecure_skip_tls_verify":      &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"mfa_code":                      &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"profile":                       &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"region":                        &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"secret_key":                    &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"skip_region_validation":        &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"skip_metadata_api_check":       &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"token":                         &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"vault_aws_engine":              &hcldec.BlockSpec{TypeName: "vault_aws_engine", Nested: hcldec.ObjectSpec((*common.FlatVaultAWSEngineOptions)(nil).HCL2Spec())},
		"ami_name":                      &hcldec.AttrSpec{Name: "ami_name", Type: cty.String, Required: false},
		"ami_description":               &hcldec.AttrSpec{Name: "ami_description", Type: cty.String, Required: false},
		"ami_virtualization_type":       &hcldec.AttrSpec{Name: "ami_virtualization_type", Type: cty.String, Required: false},
		"ami_users":                     &hcldec.AttrSpec{Name: "ami_users", Type: cty.List(cty.String), Required: false},
		"ami_groups":                    &hcldec.AttrSpec{Name: "ami_groups", Type: cty.List(cty.String), Required: false},
		"ami_product_codes":             &hcldec.AttrSpec{Name: "ami_product_codes", Type: cty.List(cty.String), Required: false},
		"ami_regions":                   &hcldec.AttrSpec{Name: "ami_regions", Type: cty.List(cty.String), Required: false},
		"tags":                          &hcldec.BlockAttrsSpec{TypeName: "tags", ElementType: cty.String, Required: false},
		"ena_support":                   &hcldec.AttrSpec{Name: "ena_support", Type: cty.Bool, Required: false},
		"sriov_support":                 &hcldec.AttrSpec{Name: "sriov_support", Type: cty.Bool, Required: false},
		"force_deregister":              &hcldec.AttrSpec{Name: "force_deregister", Type: cty.Bool, Required: false},
		"force_delete_snapshot":         &hcldec.AttrSpec{Name: "force_delete_snapshot", Type: cty.Bool, Required: false},
		"encrypt_boot":                  &hcldec.AttrSpec{Name: "encrypt_boot", Type: cty.Bool, Required: false},
		"kms_key_id":                    &hcldec.AttrSpec{Name: "kms_key_id", Type: cty.String, Required: false},
		"region_kms_key_ids":            &hcldec.BlockAttrsSpec{TypeName: "region_kms_key_ids", ElementType: cty.String, Required: false},
		"skip_save_build_region":        &hcldec.AttrSpec{Name: "skip_save_build_region", Type: cty.Bool, Required: false},
		"snapshot_tags":                 &hcldec.BlockAttrsSpec{TypeName: "snapshot_tags", ElementType: cty.String, Required: false},
		"snapshot_users":                &hcldec.AttrSpec{Name: "snapshot_users", Type: cty.List(cty.String), Required: false},
		"snapshot_groups":               &hcldec.AttrSpec{Name: "snapshot_groups", Type: cty.List(cty.String), Required: false},
		"associate_public_ip_address":   &hcldec.AttrSpec{Name: "associate_public_ip_address", Type: cty.Bool, Required: false},
		"availability_zone":             &hcldec.AttrSpec{Name: "availability_zone", Type: cty.String, Required: false},
		"block_duration_minutes":        &hcldec.AttrSpec{Name: "block_duration_minutes", Type: cty.Number, Required: false},
		"disable_stop_instance":         &hcldec.AttrSpec{Name: "disable_stop_instance", Type: cty.Bool, Required: false},
		"ebs_optimized":                 &hcldec.AttrSpec{Name: "ebs_optimized", Type: cty.Bool, Required: false},
		"enable_t2_unlimited":           &hcldec.AttrSpec{Name: "enable_t2_unlimited", Type: cty.Bool, Required: false},
		"iam_instance_profile":          &hcldec.AttrSpec{Name: "iam_instance_profile", Type: cty.String, Required: false},
		"skip_profile_validation":       &hcldec.AttrSpec{Name: "skip_profile_validation", Type: cty.Bool, Required: false},
		"temporary_iam_instance_profile_policy_document": &hcldec.BlockSpec{TypeName: "temporary_iam_instance_profile_policy_document", Nested: hcldec.ObjectSpec((*common.FlatPolicyDocument)(nil).HCL2Spec())},
		"shutdown_behavior":                     &hcldec.AttrSpec{Name: "shutdown_behavior", Type: cty.String, Required: false},
		"instance_type":                         &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"security_group_filter":                 &hcldec.BlockSpec{TypeName: "security_group_filter", Nested: hcldec.ObjectSpec((*common.FlatSecurityGroupFilterOptions)(nil).HCL2Spec())},
		"run_tags":                              &hcldec.BlockAttrsSpec{TypeName: "run_tags", ElementType: cty.String, Required: false},
		"security_group_id":                     &hcldec.AttrSpec{Name: "security_group_id", Type: cty.String, Required: false},
		"security_group_ids":                    &hcldec.AttrSpec{Name: "security_group_ids", Type: cty.List(cty.String), Required: false},
		"source_ami":                            &hcldec.AttrSpec{Name: "source_ami", Type: cty.String, Required: false},
		"source_ami_filter":                     &hcldec.BlockSpec{TypeName: "source_ami_filter", Nested: hcldec.ObjectSpec((*common.FlatAmiFilterOptions)(nil).HCL2Spec())},
		"spot_instance_types":                   &hcldec.AttrSpec{Name: "spot_instance_types", Type: cty.List(cty.String), Required: false},
		"spot_price":                            &hcldec.AttrSpec{Name: "spot_price", Type: cty.String, Required: false},
		"spot_price_auto_product":               &hcldec.AttrSpec{Name: "spot_price_auto_product", Type: cty.String, Required: false},
		"spot_tags":                             &hcldec.BlockAttrsSpec{TypeName: "spot_tags", ElementType: cty.String, Required: false},
		"subnet_filter":                         &hcldec.BlockSpec{TypeName: "subnet_filter", Nested: hcldec.ObjectSpec((*common.FlatSubnetFilterOptions)(nil).HCL2Spec())},
		"subnet_id":                             &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"temporary_key_pair_name":               &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_security_group_source_cidrs": &hcldec.AttrSpec{Name: "temporary_security_group_source_cidrs", Type: cty.List(cty.String), Required: false},
		"user_data":                             &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":                        &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"vpc_filter":                            &hcldec.BlockSpec{TypeName: "vpc_filter", Nested: hcldec.ObjectSpec((*common.FlatVpcFilterOptions)(nil).HCL2Spec())},
		"vpc_id":                                &hcldec.AttrSpec{Name: "vpc_id", Type: cty.String, Required: false},
		"windows_password_timeout":              &hcldec.AttrSpec{Name: "windows_password_timeout", Type: cty.String, Required: false},
		"communicator":                          &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":               &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                              &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                              &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                          &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                          &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":                      &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":             &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":                  &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                               &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                           &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                        &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":          &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":                &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":                      &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":                      &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":                &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":                  &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":                  &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file":          &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":              &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                        &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                        &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":                    &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":                    &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":               &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":                &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":                    &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":                     &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                        &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                       &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                        &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                        &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                            &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                            &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                         &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                         &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                        &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                        &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"ssh_interface":                         &hcldec.AttrSpec{Name: "ssh_interface", Type: cty.String, Required: false},
		"ami_block_device_mappings":             &hcldec.BlockListSpec{TypeName: "ami_block_device_mappings", Nested: hcldec.ObjectSpec((*common.FlatBlockDevice)(nil).HCL2Spec())},
		"launch_block_device_mappings":          &hcldec.BlockListSpec{TypeName: "launch_block_device_mappings", Nested: hcldec.ObjectSpec((*common.FlatBlockDevice)(nil).HCL2Spec())},
		"run_volume_tags":                       &hcldec.BlockAttrsSpec{TypeName: "run_volume_tags", ElementType: cty.String, Required: false},
		"no_ephemeral":                          &hcldec.AttrSpec{Name: "no_ephemeral", Type: cty.Bool, Required: false},
	}
	return s
}
