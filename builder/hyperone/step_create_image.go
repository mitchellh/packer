package hyperone

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
	"github.com/hyperonecom/h1-client-go"
)

type stepCreateImage struct{}

func (s *stepCreateImage) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*openapi.APIClient)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("config").(*Config)
	vmID := state.Get("vm_id").(string)

	ui.Say("Creating image...")

	image, _, err := client.ImageApi.ImageCreate(ctx, openapi.ImageCreate{
		Name:        config.ImageName,
		Vm:          vmID,
		Service:     config.ImageService,
		Description: config.ImageDescription,
		Tag:         config.ImageTags,
	})
	if err != nil {
		err := fmt.Errorf("error creating image: %s", formatOpenAPIError(err))
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("image_id", image.Id)
	state.Put("image_name", image.Name)

	return multistep.ActionContinue
}

func (s *stepCreateImage) Cleanup(state multistep.StateBag) {}
