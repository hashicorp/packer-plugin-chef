package main

import (
	"fmt"
	"os"

	chefclient "github.com/hashicorp/packer-plugin-chef/provisioner/chef-client"
	chefsolo "github.com/hashicorp/packer-plugin-chef/provisioner/chef-solo"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner("client", new(chefclient.Provisioner))
	pps.RegisterProvisioner("solo", new(chefsolo.Provisioner))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
