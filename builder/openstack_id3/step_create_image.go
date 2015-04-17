package openstack_id3

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
	"time"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
)

type stepCreateImage struct{}

func (s *stepCreateImage) Run(state multistep.StateBag) multistep.StepAction {
	
	computeClient := state.Get("compute_client").(*gophercloud.ServiceClient)
	config := state.Get("config").(config)
	server := state.Get("server").(*servers.Server)
	ui := state.Get("ui").(packer.Ui)
	
	// Create the image
	ui.Say(fmt.Sprintf("Creating the image: %s", config.ImageName))

	imageId, err := servers.CreateImage(computeClient, server.ID, 
										servers.CreateImageOpts{Name: config.ImageName}).ExtractImageID()
	if err != err {
		err := fmt.Errorf("Error creating image: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt		
	}

	// Set the Image ID in the state
	ui.Say(fmt.Sprintf("Image: %s", imageId))
	state.Put("image", imageId)

	// Wait for the image to become ready
	ui.Say("Waiting for image to become ready...")
	if err := WaitForImage(computeClient, imageId); err != nil {
		err := fmt.Errorf("Error waiting for image: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *stepCreateImage) Cleanup(multistep.StateBag) {
	// No cleanup...
}

// WaitForImage waits for the given Image ID to become ready.
func WaitForImage(computeClient *gophercloud.ServiceClient, imageId string) error {
	for {
		image, err := images.Get(computeClient, imageId).Extract()
		if err != nil {
			return err
		}

		if image.Status == "ACTIVE" {
			return nil
		}

		log.Printf("Waiting for image creation status: %s (%d%%)", image.Status, image.Progress)
		time.Sleep(2 * time.Second)
	}
}
