package qemu

import (
	commonssh "github.com/cstuntz/packer/common/ssh"
	"github.com/cstuntz/packer/communicator/ssh"
	"github.com/mitchellh/multistep"
	gossh "golang.org/x/crypto/ssh"
)

func commHost(state multistep.StateBag) (string, error) {
	return "127.0.0.1", nil
}

func commPort(state multistep.StateBag) (int, error) {
	sshHostPort := state.Get("sshHostPort").(uint)
	return int(sshHostPort), nil
}

func sshConfig(state multistep.StateBag) (*gossh.ClientConfig, error) {
	config := state.Get("config").(*Config)

	auth := []gossh.AuthMethod{
		gossh.Password(config.Comm.SSHPassword),
		gossh.KeyboardInteractive(
			ssh.PasswordKeyboardInteractive(config.Comm.SSHPassword)),
	}

	if config.Comm.SSHPrivateKey != "" {
		signer, err := commonssh.FileSigner(config.Comm.SSHPrivateKey)
		if err != nil {
			return nil, err
		}

		auth = append(auth, gossh.PublicKeys(signer))
	}

	return &gossh.ClientConfig{
		User:            config.Comm.SSHUsername,
		Auth:            auth,
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}, nil
}
