/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"github.com/spf13/cobra"

	// common imports for subcommands
	cmdgenerate "github.com/tbd-paas/platform-config-operator/cmd/platformctl/commands/generate"
	cmdinit "github.com/tbd-paas/platform-config-operator/cmd/platformctl/commands/init"
	cmdversion "github.com/tbd-paas/platform-config-operator/cmd/platformctl/commands/version"

	// specific imports for workloads
	generatedeploy "github.com/tbd-paas/platform-config-operator/cmd/platformctl/commands/generate/deploy"
	initdeploy "github.com/tbd-paas/platform-config-operator/cmd/platformctl/commands/init/deploy"
	versiondeploy "github.com/tbd-paas/platform-config-operator/cmd/platformctl/commands/version/deploy"
	// +kubebuilder:scaffold:operator-builder:subcommands:imports
)

// PlatformctlCommand represents the base command when called without any subcommands.
type PlatformctlCommand struct {
	*cobra.Command
}

// NewPlatformctlCommand returns an instance of the PlatformctlCommand.
func NewPlatformctlCommand() *PlatformctlCommand {
	c := &PlatformctlCommand{
		Command: &cobra.Command{
			Use:   "platformctl",
			Short: "Manage the platform configuration",
			Long:  "Manage the platform configuration",
		},
	}

	c.addSubCommands()

	return c
}

// Run represents the main entry point into the command
// This is called by main.main() to execute the root command.
func (c *PlatformctlCommand) Run() {
	cobra.CheckErr(c.Execute())
}

func (c *PlatformctlCommand) newInitSubCommand() {
	parentCommand := cmdinit.GetParent(c.Command)
	_ = parentCommand

	// add the init subcommands
	initdeploy.NewPlatformOperatorsSubCommand(parentCommand)
	initdeploy.NewPlatformConfigSubCommand(parentCommand)
	// +kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *PlatformctlCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(c.Command)
	_ = parentCommand

	// add the generate subcommands
	generatedeploy.NewPlatformOperatorsSubCommand(parentCommand)
	generatedeploy.NewPlatformConfigSubCommand(parentCommand)
	// +kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *PlatformctlCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(c.Command)
	_ = parentCommand

	// add the version subcommands
	versiondeploy.NewPlatformOperatorsSubCommand(parentCommand)
	versiondeploy.NewPlatformConfigSubCommand(parentCommand)
	// +kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *PlatformctlCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
