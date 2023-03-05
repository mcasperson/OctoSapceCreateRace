package main

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/mcasperson/OctoSapceCreateRace/internal/test"
	"path/filepath"
	"testing"
)

func TestCreateSpaceAndUseIt(t *testing.T) {
	testFramework := test.OctopusContainerTest{}
	testFramework.ArrangeTest(t, func(t *testing.T, container *test.OctopusContainer, client *client.Client) error {
		_, err := testFramework.Act(t, container, filepath.Join("test", "terraform", "2-usenewspace"), []string{})
		return err
	})
}
