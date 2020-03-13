// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package iso

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	HTTPDir                   *string           `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin               *int              `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax               *int              `mapstructure:"http_port_max" cty:"http_port_max"`
	VCenterServer             *string           `mapstructure:"vcenter_server" cty:"vcenter_server"`
	Username                  *string           `mapstructure:"username" cty:"username"`
	Password                  *string           `mapstructure:"password" cty:"password"`
	InsecureConnection        *bool             `mapstructure:"insecure_connection" cty:"insecure_connection"`
	Datacenter                *string           `mapstructure:"datacenter" cty:"datacenter"`
	Version                   *uint             `mapstructure:"vm_version" cty:"vm_version"`
	GuestOSType               *string           `mapstructure:"guest_os_type" cty:"guest_os_type"`
	Firmware                  *string           `mapstructure:"firmware" cty:"firmware"`
	DiskControllerType        *string           `mapstructure:"disk_controller_type" cty:"disk_controller_type"`
	DiskSize                  *int64            `mapstructure:"disk_size" cty:"disk_size"`
	DiskThinProvisioned       *bool             `mapstructure:"disk_thin_provisioned" cty:"disk_thin_provisioned"`
	DiskEagerlyScrub          *bool             `mapstructure:"disk_eagerly_scrub" cty:"disk_eagerly_scrub"`
	Storage                   []FlatDiskConfig  `mapstructure:"storage" cty:"storage"`
	Network                   *string           `mapstructure:"network" cty:"network"`
	NetworkCard               *string           `mapstructure:"network_card" cty:"network_card"`
	NICs                      []FlatNIC         `mapstructure:"network_adapters" cty:"network_adapters"`
	USBController             *bool             `mapstructure:"usb_controller" cty:"usb_controller"`
	Notes                     *string           `mapstructure:"notes" cty:"notes"`
	VMName                    *string           `mapstructure:"vm_name" cty:"vm_name"`
	Folder                    *string           `mapstructure:"folder" cty:"folder"`
	Cluster                   *string           `mapstructure:"cluster" cty:"cluster"`
	Host                      *string           `mapstructure:"host" cty:"host"`
	ResourcePool              *string           `mapstructure:"resource_pool" cty:"resource_pool"`
	Datastore                 *string           `mapstructure:"datastore" cty:"datastore"`
	CPUs                      *int32            `mapstructure:"CPUs" cty:"CPUs"`
	CpuCores                  *int32            `mapstructure:"cpu_cores" cty:"cpu_cores"`
	CPUReservation            *int64            `mapstructure:"CPU_reservation" cty:"CPU_reservation"`
	CPULimit                  *int64            `mapstructure:"CPU_limit" cty:"CPU_limit"`
	CpuHotAddEnabled          *bool             `mapstructure:"CPU_hot_plug" cty:"CPU_hot_plug"`
	RAM                       *int64            `mapstructure:"RAM" cty:"RAM"`
	RAMReservation            *int64            `mapstructure:"RAM_reservation" cty:"RAM_reservation"`
	RAMReserveAll             *bool             `mapstructure:"RAM_reserve_all" cty:"RAM_reserve_all"`
	MemoryHotAddEnabled       *bool             `mapstructure:"RAM_hot_plug" cty:"RAM_hot_plug"`
	VideoRAM                  *int64            `mapstructure:"video_ram" cty:"video_ram"`
	NestedHV                  *bool             `mapstructure:"NestedHV" cty:"NestedHV"`
	ConfigParams              map[string]string `mapstructure:"configuration_parameters" cty:"configuration_parameters"`
	ISOChecksum               *string           `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum"`
	ISOChecksumURL            *string           `mapstructure:"iso_checksum_url" cty:"iso_checksum_url"`
	ISOChecksumType           *string           `mapstructure:"iso_checksum_type" cty:"iso_checksum_type"`
	RawSingleISOUrl           *string           `mapstructure:"iso_url" required:"true" cty:"iso_url"`
	ISOUrls                   []string          `mapstructure:"iso_urls" cty:"iso_urls"`
	TargetPath                *string           `mapstructure:"iso_target_path" cty:"iso_target_path"`
	TargetExtension           *string           `mapstructure:"iso_target_extension" cty:"iso_target_extension"`
	CdromType                 *string           `mapstructure:"cdrom_type" cty:"cdrom_type"`
	ISOPaths                  []string          `mapstructure:"iso_paths" cty:"iso_paths"`
	RemoveCdrom               *bool             `mapstructure:"remove_cdrom" cty:"remove_cdrom"`
	FloppyIMGPath             *string           `mapstructure:"floppy_img_path" cty:"floppy_img_path"`
	FloppyFiles               []string          `mapstructure:"floppy_files" cty:"floppy_files"`
	FloppyDirectories         []string          `mapstructure:"floppy_dirs" cty:"floppy_dirs"`
	BootOrder                 *string           `mapstructure:"boot_order" cty:"boot_order"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command"`
	BootWait                  *string           `mapstructure:"boot_wait" cty:"boot_wait"`
	HTTPIP                    *string           `mapstructure:"http_ip" cty:"http_ip"`
	WaitTimeout               *string           `mapstructure:"ip_wait_timeout" cty:"ip_wait_timeout"`
	SettleTimeout             *string           `mapstructure:"ip_settle_timeout" cty:"ip_settle_timeout"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionInteractive     *bool             `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `mapstructure:"ssh_public_key" cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `mapstructure:"ssh_private_key" cty:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	Command                   *string           `mapstructure:"shutdown_command" cty:"shutdown_command"`
	Timeout                   *string           `mapstructure:"shutdown_timeout" cty:"shutdown_timeout"`
	CreateSnapshot            *bool             `mapstructure:"create_snapshot" cty:"create_snapshot"`
	ConvertToTemplate         *bool             `mapstructure:"convert_to_template" cty:"convert_to_template"`
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
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"vcenter_server":               &hcldec.AttrSpec{Name: "vcenter_server", Type: cty.String, Required: false},
		"username":                     &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":                     &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"insecure_connection":          &hcldec.AttrSpec{Name: "insecure_connection", Type: cty.Bool, Required: false},
		"datacenter":                   &hcldec.AttrSpec{Name: "datacenter", Type: cty.String, Required: false},
		"vm_version":                   &hcldec.AttrSpec{Name: "vm_version", Type: cty.Number, Required: false},
		"guest_os_type":                &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"firmware":                     &hcldec.AttrSpec{Name: "firmware", Type: cty.String, Required: false},
		"disk_controller_type":         &hcldec.AttrSpec{Name: "disk_controller_type", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"disk_thin_provisioned":        &hcldec.AttrSpec{Name: "disk_thin_provisioned", Type: cty.Bool, Required: false},
		"disk_eagerly_scrub":           &hcldec.AttrSpec{Name: "disk_eagerly_scrub", Type: cty.Bool, Required: false},
		"storage":                      &hcldec.BlockListSpec{TypeName: "storage", Nested: hcldec.ObjectSpec((*FlatDiskConfig)(nil).HCL2Spec())},
		"network":                      &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"network_card":                 &hcldec.AttrSpec{Name: "network_card", Type: cty.String, Required: false},
		"network_adapters":             &hcldec.BlockListSpec{TypeName: "network_adapters", Nested: hcldec.ObjectSpec((*FlatNIC)(nil).HCL2Spec())},
		"usb_controller":               &hcldec.AttrSpec{Name: "usb_controller", Type: cty.Bool, Required: false},
		"notes":                        &hcldec.AttrSpec{Name: "notes", Type: cty.String, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"folder":                       &hcldec.AttrSpec{Name: "folder", Type: cty.String, Required: false},
		"cluster":                      &hcldec.AttrSpec{Name: "cluster", Type: cty.String, Required: false},
		"host":                         &hcldec.AttrSpec{Name: "host", Type: cty.String, Required: false},
		"resource_pool":                &hcldec.AttrSpec{Name: "resource_pool", Type: cty.String, Required: false},
		"datastore":                    &hcldec.AttrSpec{Name: "datastore", Type: cty.String, Required: false},
		"CPUs":                         &hcldec.AttrSpec{Name: "CPUs", Type: cty.Number, Required: false},
		"cpu_cores":                    &hcldec.AttrSpec{Name: "cpu_cores", Type: cty.Number, Required: false},
		"CPU_reservation":              &hcldec.AttrSpec{Name: "CPU_reservation", Type: cty.Number, Required: false},
		"CPU_limit":                    &hcldec.AttrSpec{Name: "CPU_limit", Type: cty.Number, Required: false},
		"CPU_hot_plug":                 &hcldec.AttrSpec{Name: "CPU_hot_plug", Type: cty.Bool, Required: false},
		"RAM":                          &hcldec.AttrSpec{Name: "RAM", Type: cty.Number, Required: false},
		"RAM_reservation":              &hcldec.AttrSpec{Name: "RAM_reservation", Type: cty.Number, Required: false},
		"RAM_reserve_all":              &hcldec.AttrSpec{Name: "RAM_reserve_all", Type: cty.Bool, Required: false},
		"RAM_hot_plug":                 &hcldec.AttrSpec{Name: "RAM_hot_plug", Type: cty.Bool, Required: false},
		"video_ram":                    &hcldec.AttrSpec{Name: "video_ram", Type: cty.Number, Required: false},
		"NestedHV":                     &hcldec.AttrSpec{Name: "NestedHV", Type: cty.Bool, Required: false},
		"configuration_parameters":     &hcldec.BlockAttrsSpec{TypeName: "configuration_parameters", ElementType: cty.String, Required: false},
		"iso_checksum":                 &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_checksum_url":             &hcldec.AttrSpec{Name: "iso_checksum_url", Type: cty.String, Required: false},
		"iso_checksum_type":            &hcldec.AttrSpec{Name: "iso_checksum_type", Type: cty.String, Required: false},
		"iso_url":                      &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                     &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":              &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":         &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"cdrom_type":                   &hcldec.AttrSpec{Name: "cdrom_type", Type: cty.String, Required: false},
		"iso_paths":                    &hcldec.AttrSpec{Name: "iso_paths", Type: cty.List(cty.String), Required: false},
		"remove_cdrom":                 &hcldec.AttrSpec{Name: "remove_cdrom", Type: cty.Bool, Required: false},
		"floppy_img_path":              &hcldec.AttrSpec{Name: "floppy_img_path", Type: cty.String, Required: false},
		"floppy_files":                 &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"floppy_dirs":                  &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"boot_order":                   &hcldec.AttrSpec{Name: "boot_order", Type: cty.String, Required: false},
		"boot_command":                 &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"boot_wait":                    &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"http_ip":                      &hcldec.AttrSpec{Name: "http_ip", Type: cty.String, Required: false},
		"ip_wait_timeout":              &hcldec.AttrSpec{Name: "ip_wait_timeout", Type: cty.String, Required: false},
		"ip_settle_timeout":            &hcldec.AttrSpec{Name: "ip_settle_timeout", Type: cty.String, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
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
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"shutdown_command":             &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"shutdown_timeout":             &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"create_snapshot":              &hcldec.AttrSpec{Name: "create_snapshot", Type: cty.Bool, Required: false},
		"convert_to_template":          &hcldec.AttrSpec{Name: "convert_to_template", Type: cty.Bool, Required: false},
	}
	return s
}
