package packer_registry

import (
	"crypto/sha1"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2021-04-30/models"
)

type Builds struct {
	sync.RWMutex
	m map[string]*Build
}

type Build struct {
	ID            string
	CloudProvider string
	ComponentType string
	RunUUID       string
	Metadata      map[string]string
	PARtifacts    []PARtifact
	Status        models.HashicorpCloudPackerBuildStatus
}

func NewBuilds() Builds {
	return Builds{
		m: make(map[string]*Build),
	}
}

type PARtifact struct {
	ID                           string
	ProviderName, ProviderRegion string
}

type Iteration struct {
	ID           string
	Author       string
	AncestorSlug string
	Fingerprint  string
	RunUUID      string
	Labels       map[string]string
	Builds       Builds
}

type IterationOptions struct {
	RunUUID       string
	UseGitBackend bool
}

func GetGitFingerprint() (string, error) {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return "", fmt.Errorf("Error loading git sha", err)
	}
	// The config can be used to retrieve user identity. for example,
	// c.User.Email. Leaving in but commented because I'm not sure we care
	// about this identity right now. - Megan
	//
	// c, err := r.ConfigScoped(config.GlobalScope)
	// if err != nil {
	// 	return "", fmt.Errorf("Error setting git scope", err)
	// }
	ref, _ := r.Head()
	// log.Printf("Author: %v, Commit: %v\n", c.User.Email, ref.Hash())
	return ref.Hash().String(), nil
}

func NewIteration(opts IterationOptions) (*Iteration, error) {
	i := Iteration{
		Builds:  NewBuilds(),
		RunUUID: opts.RunUUID,
	}

	i.RunUUID = os.Getenv("PACKER_RUN_UUID")

	if !opts.UseGitBackend {
		i.Author = os.Getenv("USER")
		s := []byte(time.Now().String())
		// TODO allow user to set fingerprint through Packer block or
		// environment variable?
		i.Fingerprint = fmt.Sprintf("%x", sha1.Sum(s))
		//i.Fingerprint = "00ee249320213a1e20578a551c11f47bbdd94ea4"
	} else {
		fp, err := GetGitFingerprint()
		if err != nil {
			return nil, err
		}
		i.Fingerprint = fp
	}

	return &i, nil
}
