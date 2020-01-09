package hcl2template

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/packer/packer"
	"github.com/zclconf/go-cty/cty"
)

// PackerConfig represents a loaded packer config
type PackerConfig struct {
	Sources map[SourceRef]*Source

	Variables PackerV1Variables

	Builds Builds
}

func (p *Parser) CoreBuildProvisioners(blocks []*ProvisionerBlock, ectx *hcl.EvalContext, generatedVars []string) ([]packer.CoreBuildProvisioner, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	res := []packer.CoreBuildProvisioner{}
	for _, pb := range blocks {
		provisioner, moreDiags := p.StartProvisioner(pb, ectx, generatedVars)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			continue
		}
		res = append(res, packer.CoreBuildProvisioner{
			PType:       pb.PType,
			Provisioner: provisioner,
		})
	}
	return res, diags
}

func (p *Parser) CoreBuildPostProcessors(blocks []*PostProcessorBlock, ectx *hcl.EvalContext) ([]packer.CoreBuildPostProcessor, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	res := []packer.CoreBuildPostProcessor{}
	for _, pp := range blocks {
		postProcessor, moreDiags := p.StartPostProcessor(pp, ectx)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			continue
		}
		res = append(res, packer.CoreBuildPostProcessor{
			PostProcessor: postProcessor,
			PType:         pp.PType,
		})
	}

	return res, diags
}

func (p *Parser) getBuilds(cfg *PackerConfig) ([]packer.Build, hcl.Diagnostics) {
	res := []packer.Build{}
	var diags hcl.Diagnostics

	ectx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"var": cty.ObjectVal(cfg.Variables),
		},
	}

	for _, build := range cfg.Builds {
		for _, from := range build.Froms {
			src, found := cfg.Sources[from]
			if !found {
				diags = append(diags, &hcl.Diagnostic{
					Summary:  "Unknown " + sourceLabel + " " + from.String(),
					Subject:  build.HCL2Ref.DeclRange.Ptr(),
					Severity: hcl.DiagError,
				})
				continue
			}
			builder, moreDiags, generatedVars := p.StartBuilder(src, ectx)
			diags = append(diags, moreDiags...)
			if moreDiags.HasErrors() {
				continue
			}
			provisioners, moreDiags := p.CoreBuildProvisioners(build.ProvisionerBlocks, ectx, generatedVars)
			diags = append(diags, moreDiags...)
			if moreDiags.HasErrors() {
				continue
			}
			postProcessors, moreDiags := p.CoreBuildPostProcessors(build.PostProcessors, ectx)
			pps := [][]packer.CoreBuildPostProcessor{}
			if len(postProcessors) > 0 {
				pps = [][]packer.CoreBuildPostProcessor{postProcessors}
			}
			diags = append(diags, moreDiags...)
			if moreDiags.HasErrors() {
				continue
			}

			pcb := &packer.CoreBuild{
				Type:           src.Type,
				Builder:        builder,
				Provisioners:   provisioners,
				PostProcessors: pps,
			}
			res = append(res, pcb)
		}
	}
	return res, diags
}

// Parse will parse HCL file(s) in path. Path can be a folder or a file.
//
// Parse will first parse variables and then the rest; so that interpolation
// can happen.
//
// Parse then return a slice of packer.Builds; which are what packer core uses
// to run builds.
func (p *Parser) Parse(path string) ([]packer.Build, hcl.Diagnostics) {
	cfg, diags := p.parse(path)
	if diags.HasErrors() {
		return nil, diags
	}

	builds, moreDiags := p.getBuilds(cfg)
	return builds, append(diags, moreDiags...)
}
