//
// This file is automatically generated by scripts/generate-plugins.go -- Do not edit!
//

package command

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/plugin"

	alicloudecsbuilder "github.com/hashicorp/packer/builder/alicloud/ecs"
	azurearmbuilder "github.com/hashicorp/packer/builder/azure/arm"
	azurechrootbuilder "github.com/hashicorp/packer/builder/azure/chroot"
	azuredtlbuilder "github.com/hashicorp/packer/builder/azure/dtl"
	cloudstackbuilder "github.com/hashicorp/packer/builder/cloudstack"
	digitaloceanbuilder "github.com/hashicorp/packer/builder/digitalocean"
	filebuilder "github.com/hashicorp/packer/builder/file"
	hcloudbuilder "github.com/hashicorp/packer/builder/hcloud"
	hyperonebuilder "github.com/hashicorp/packer/builder/hyperone"
	hypervisobuilder "github.com/hashicorp/packer/builder/hyperv/iso"
	hypervvmcxbuilder "github.com/hashicorp/packer/builder/hyperv/vmcx"
	jdcloudbuilder "github.com/hashicorp/packer/builder/jdcloud"
	linodebuilder "github.com/hashicorp/packer/builder/linode"
	lxcbuilder "github.com/hashicorp/packer/builder/lxc"
	lxdbuilder "github.com/hashicorp/packer/builder/lxd"
	nullbuilder "github.com/hashicorp/packer/builder/null"
	oneandonebuilder "github.com/hashicorp/packer/builder/oneandone"
	openstackbuilder "github.com/hashicorp/packer/builder/openstack"
	oracleclassicbuilder "github.com/hashicorp/packer/builder/oracle/classic"
	oracleocibuilder "github.com/hashicorp/packer/builder/oracle/oci"
	parallelsisobuilder "github.com/hashicorp/packer/builder/parallels/iso"
	parallelspvmbuilder "github.com/hashicorp/packer/builder/parallels/pvm"
	profitbricksbuilder "github.com/hashicorp/packer/builder/profitbricks"
	scalewaybuilder "github.com/hashicorp/packer/builder/scaleway"
	tencentcloudcvmbuilder "github.com/hashicorp/packer/builder/tencentcloud/cvm"
	tritonbuilder "github.com/hashicorp/packer/builder/triton"
	uclouduhostbuilder "github.com/hashicorp/packer/builder/ucloud/uhost"
	vagrantbuilder "github.com/hashicorp/packer/builder/vagrant"
	yandexbuilder "github.com/hashicorp/packer/builder/yandex"
	alicloudimportpostprocessor "github.com/hashicorp/packer/post-processor/alicloud-import"
	artificepostprocessor "github.com/hashicorp/packer/post-processor/artifice"
	checksumpostprocessor "github.com/hashicorp/packer/post-processor/checksum"
	compresspostprocessor "github.com/hashicorp/packer/post-processor/compress"
	digitaloceanimportpostprocessor "github.com/hashicorp/packer/post-processor/digitalocean-import"
	manifestpostprocessor "github.com/hashicorp/packer/post-processor/manifest"
	shelllocalpostprocessor "github.com/hashicorp/packer/post-processor/shell-local"
	ucloudimportpostprocessor "github.com/hashicorp/packer/post-processor/ucloud-import"
	vagrantpostprocessor "github.com/hashicorp/packer/post-processor/vagrant"
	vagrantcloudpostprocessor "github.com/hashicorp/packer/post-processor/vagrant-cloud"
	yandexexportpostprocessor "github.com/hashicorp/packer/post-processor/yandex-export"
	yandeximportpostprocessor "github.com/hashicorp/packer/post-processor/yandex-import"
	azuredtlartifactprovisioner "github.com/hashicorp/packer/provisioner/azure-dtlartifact"
	breakpointprovisioner "github.com/hashicorp/packer/provisioner/breakpoint"
	chefclientprovisioner "github.com/hashicorp/packer/provisioner/chef-client"
	chefsoloprovisioner "github.com/hashicorp/packer/provisioner/chef-solo"
	convergeprovisioner "github.com/hashicorp/packer/provisioner/converge"
	fileprovisioner "github.com/hashicorp/packer/provisioner/file"
	inspecprovisioner "github.com/hashicorp/packer/provisioner/inspec"
	powershellprovisioner "github.com/hashicorp/packer/provisioner/powershell"
	puppetmasterlessprovisioner "github.com/hashicorp/packer/provisioner/puppet-masterless"
	puppetserverprovisioner "github.com/hashicorp/packer/provisioner/puppet-server"
	saltmasterlessprovisioner "github.com/hashicorp/packer/provisioner/salt-masterless"
	shellprovisioner "github.com/hashicorp/packer/provisioner/shell"
	shelllocalprovisioner "github.com/hashicorp/packer/provisioner/shell-local"
	sleepprovisioner "github.com/hashicorp/packer/provisioner/sleep"
	windowsrestartprovisioner "github.com/hashicorp/packer/provisioner/windows-restart"
	windowsshellprovisioner "github.com/hashicorp/packer/provisioner/windows-shell"
)

type PluginCommand struct {
	Meta
}

var Builders = map[string]packersdk.Builder{
	"alicloud-ecs":     new(alicloudecsbuilder.Builder),
	"azure-arm":        new(azurearmbuilder.Builder),
	"azure-chroot":     new(azurechrootbuilder.Builder),
	"azure-dtl":        new(azuredtlbuilder.Builder),
	"cloudstack":       new(cloudstackbuilder.Builder),
	"digitalocean":     new(digitaloceanbuilder.Builder),
	"file":             new(filebuilder.Builder),
	"hcloud":           new(hcloudbuilder.Builder),
	"hyperone":         new(hyperonebuilder.Builder),
	"hyperv-iso":       new(hypervisobuilder.Builder),
	"hyperv-vmcx":      new(hypervvmcxbuilder.Builder),
	"jdcloud":          new(jdcloudbuilder.Builder),
	"linode":           new(linodebuilder.Builder),
	"lxc":              new(lxcbuilder.Builder),
	"lxd":              new(lxdbuilder.Builder),
	"null":             new(nullbuilder.Builder),
	"oneandone":        new(oneandonebuilder.Builder),
	"openstack":        new(openstackbuilder.Builder),
	"oracle-classic":   new(oracleclassicbuilder.Builder),
	"oracle-oci":       new(oracleocibuilder.Builder),
	"parallels-iso":    new(parallelsisobuilder.Builder),
	"parallels-pvm":    new(parallelspvmbuilder.Builder),
	"profitbricks":     new(profitbricksbuilder.Builder),
	"tencentcloud-cvm": new(tencentcloudcvmbuilder.Builder),
	"triton":           new(tritonbuilder.Builder),
	"ucloud-uhost":     new(uclouduhostbuilder.Builder),
	"vagrant":          new(vagrantbuilder.Builder),
	"yandex":           new(yandexbuilder.Builder),
}

var Provisioners = map[string]packersdk.Provisioner{
	"azure-dtlartifact": new(azuredtlartifactprovisioner.Provisioner),
	"breakpoint":        new(breakpointprovisioner.Provisioner),
	"chef-client":       new(chefclientprovisioner.Provisioner),
	"chef-solo":         new(chefsoloprovisioner.Provisioner),
	"converge":          new(convergeprovisioner.Provisioner),
	"file":              new(fileprovisioner.Provisioner),
	"inspec":            new(inspecprovisioner.Provisioner),
	"powershell":        new(powershellprovisioner.Provisioner),
	"puppet-masterless": new(puppetmasterlessprovisioner.Provisioner),
	"puppet-server":     new(puppetserverprovisioner.Provisioner),
	"salt-masterless":   new(saltmasterlessprovisioner.Provisioner),
	"shell":             new(shellprovisioner.Provisioner),
	"shell-local":       new(shelllocalprovisioner.Provisioner),
	"sleep":             new(sleepprovisioner.Provisioner),
	"windows-restart":   new(windowsrestartprovisioner.Provisioner),
	"windows-shell":     new(windowsshellprovisioner.Provisioner),
}

var PostProcessors = map[string]packersdk.PostProcessor{
	"alicloud-import":     new(alicloudimportpostprocessor.PostProcessor),
	"artifice":            new(artificepostprocessor.PostProcessor),
	"checksum":            new(checksumpostprocessor.PostProcessor),
	"compress":            new(compresspostprocessor.PostProcessor),
	"digitalocean-import": new(digitaloceanimportpostprocessor.PostProcessor),
	"manifest":            new(manifestpostprocessor.PostProcessor),
	"shell-local":         new(shelllocalpostprocessor.PostProcessor),
	"ucloud-import":       new(ucloudimportpostprocessor.PostProcessor),
	"vagrant":             new(vagrantpostprocessor.PostProcessor),
	"vagrant-cloud":       new(vagrantcloudpostprocessor.PostProcessor),
	"yandex-export":       new(yandexexportpostprocessor.PostProcessor),
	"yandex-import":       new(yandeximportpostprocessor.PostProcessor),
}

var Datasources = map[string]packersdk.Datasource{}

var pluginRegexp = regexp.MustCompile("packer-(builder|post-processor|provisioner|datasource)-(.+)")

func (c *PluginCommand) Run(args []string) int {
	// This is an internal call (users should not call this directly) so we're
	// not going to do much input validation. If there's a problem we'll often
	// just crash. Error handling should be added to facilitate debugging.
	log.Printf("args: %#v", args)
	if len(args) != 1 {
		c.Ui.Error("Wrong number of args")
		return 1
	}

	// Plugin will match something like "packer-builder-amazon-ebs"
	parts := pluginRegexp.FindStringSubmatch(args[0])
	if len(parts) != 3 {
		c.Ui.Error(fmt.Sprintf("Error parsing plugin argument [DEBUG]: %#v", parts))
		return 1
	}
	pluginType := parts[1] // capture group 1 (builder|post-processor|provisioner)
	pluginName := parts[2] // capture group 2 (.+)

	server, err := plugin.Server()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error starting plugin server: %s", err))
		return 1
	}

	switch pluginType {
	case "builder":
		builder, found := Builders[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load builder: %s", pluginName))
			return 1
		}
		server.RegisterBuilder(builder)
	case "provisioner":
		provisioner, found := Provisioners[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load provisioner: %s", pluginName))
			return 1
		}
		server.RegisterProvisioner(provisioner)
	case "post-processor":
		postProcessor, found := PostProcessors[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load post-processor: %s", pluginName))
			return 1
		}
		server.RegisterPostProcessor(postProcessor)
	case "datasource":
		datasource, found := Datasources[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load datasource: %s", pluginName))
			return 1
		}
		server.RegisterDatasource(datasource)
	}

	server.Serve()

	return 0
}

func (*PluginCommand) Help() string {
	helpText := `
Usage: packer plugin PLUGIN

  Runs an internally-compiled version of a plugin from the packer binary.

  NOTE: this is an internal command and you should not call it yourself.
`

	return strings.TrimSpace(helpText)
}

func (c *PluginCommand) Synopsis() string {
	return "internal plugin command"
}
