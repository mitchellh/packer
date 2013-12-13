package googlecompute

import (
	"errors"
	"fmt"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common/uuid"
	"github.com/mitchellh/packer/packer"
)

// StepCreateInstance represents a Packer build step that creates GCE instances.
type StepCreateInstance struct {
	instanceName string
}

// Run executes the Packer build step that creates a GCE instance.
func (s *StepCreateInstance) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	driver := state.Get("driver").(Driver)
	sshPublicKey := state.Get("ssh_public_key").(string)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Creating instance...")
	name := fmt.Sprintf("packer-%s", uuid.TimeOrderedUUID())

	errCh, err := driver.RunInstance(&InstanceConfig{
		Description: "New instance created by Packer",
		Image:       config.SourceImage,
		MachineType: config.MachineType,
		Metadata: map[string]string{
			"sshKeys": fmt.Sprintf("%s:%s", config.SSHUsername, sshPublicKey),
		},
		Name:    name,
		Network: config.Network,
		Tags:    config.Tags,
		Zone:    config.Zone,
	})

	if err == nil {
		ui.Message("Waiting for creation operation to complete...")
		select {
		case err = <-errCh:
		case <-time.After(config.stateTimeout):
			err = errors.New("time out while waiting for instance to create")
		}
	}

	if err != nil {
		err := fmt.Errorf("Error creating instance: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Message("Instance has been created!")

	// Things succeeded, store the name so we can remove it later
	state.Put("instance_name", name)
	s.instanceName = name

	return multistep.ActionContinue
}

// Cleanup destroys the GCE instance created during the image creation process.
func (s *StepCreateInstance) Cleanup(state multistep.StateBag) {
	if s.instanceName == "" {
		return
	}
	/*
		var (
			client = state.Get("client").(*GoogleComputeClient)
			config = state.Get("config").(*Config)
			ui     = state.Get("ui").(packer.Ui)
		)
		ui.Say("Destroying instance...")
		operation, err := client.DeleteInstance(config.Zone, s.instanceName)
		if err != nil {
			ui.Error(fmt.Sprintf("Error destroying instance. Please destroy it manually: %v", s.instanceName))
		}
		ui.Say("Waiting for the instance to be deleted...")
		for {
			status, err := client.ZoneOperationStatus(config.Zone, operation.Name)
			if err != nil {
				ui.Error(fmt.Sprintf("Error destroying instance. Please destroy it manually: %v", s.instanceName))
			}
			if status == "DONE" {
				break
			}
		}
	*/
	return
}
