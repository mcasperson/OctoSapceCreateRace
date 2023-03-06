package main

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/mcasperson/OctopusTerraformTestFramework/test"
	"testing"
)

func TestCreateSpaceAndUseIt(t *testing.T) {
	testFramework := test.OctopusContainerTest{}

	for i := 0; i < 10; i++ {
		t.Log("Attempt " + fmt.Sprint(i+1))

		testFramework.ArrangeTest(t, func(t *testing.T, container *test.OctopusContainer, client *client.Client) error {
			_, err := testFramework.Act(t, container, "terraform", "2-usenewspace", []string{})
			return err
		})
	}
}
