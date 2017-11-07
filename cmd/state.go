/*
Copyright 2017 Ankyra

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

package cmd

import (
	"fmt"

	"github.com/ankyra/escape/controllers"
	"github.com/spf13/cobra"
)

var deployStage bool
var extraVars, extraProviders []string

var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Manage the Escape state file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("Unknown command " + args[0])
		}
		cmd.UsageFunc()(cmd)
		return nil
	},
}

var listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "Show the deployments",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := context.LoadLocalState(state, environment); err != nil {
			return err
		}
		return controllers.StateController{}.ListDeployments(context)
	},
}

var showDeploymentCmd = &cobra.Command{
	Use:   "show-deployment",
	Short: "Show a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		if deployment == "" {
			return fmt.Errorf("Missing deployment name")
		}
		if err := ProcessFlagsForContext(false); err != nil {
			return err
		}
		return controllers.StateController{}.ShowDeployment(context, deployment)
	},
}

var showProvidersCmd = &cobra.Command{
	Use:   "show-providers",
	Short: "Show the providers available in the environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := ProcessFlagsForContext(false); err != nil {
			return err
		}
		return controllers.StateController{}.ShowProviders(context)
	},
}

var createStateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create state for the given escape plan",
	RunE: func(cmd *cobra.Command, args []string) error {
		useEscapePlan := len(args) == 0
		if err := ProcessFlagsForContext(useEscapePlan); err != nil {
			return err
		}
		stage := "build"
		if deployStage {
			stage = "deploy"
		}
		parsedExtraVars, err := ParseExtraVars(extraVars)
		if err != nil {
			return err
		}
		parsedExtraProviders, err := ParseExtraVars(extraProviders)
		if err != nil {
			return err
		}
		if !useEscapePlan {
			if err := context.InitReleaseMetadataByReleaseId(args[0]); err != nil {
				return err
			}
		}
		return controllers.StateController{}.CreateState(context, stage, parsedExtraVars, parsedExtraProviders)
	},
}

var showStateCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := ProcessFlagsForContext(true); err != nil {
			return err
		}
		return controllers.StateController{}.ShowDeployment(context, context.GetRootDeploymentName())
	},
}

func init() {
	RootCmd.AddCommand(stateCmd)
	stateCmd.AddCommand(listDeploymentsCmd)
	stateCmd.AddCommand(showDeploymentCmd)
	stateCmd.AddCommand(showProvidersCmd)
	stateCmd.AddCommand(createStateCmd)
	stateCmd.AddCommand(showStateCmd)

	setEscapeStateLocationFlag(listDeploymentsCmd)
	setEscapeStateEnvironmentFlag(listDeploymentsCmd)

	setEscapeStateLocationFlag(showDeploymentCmd)
	setEscapeStateEnvironmentFlag(showDeploymentCmd)
	setEscapeDeploymentFlag(showDeploymentCmd)

	setEscapeStateLocationFlag(showProvidersCmd)
	setEscapeStateEnvironmentFlag(showProvidersCmd)

	setPlanAndStateFlags(showStateCmd)
	setPlanAndStateFlags(createStateCmd)
	createStateCmd.Flags().BoolVarP(&deployStage, "deploy", "", false, "Use deployment instead of build stage")
	createStateCmd.Flags().StringArrayVarP(&extraVars, "extra-vars", "v", []string{}, "Extra variables (format: key=value, key=@value.txt, @values.json)")
	createStateCmd.Flags().StringArrayVarP(&extraProviders, "extra-providers", "p", []string{}, "Extra providers (format: provider=deployment, provider=@deployment.txt, @values.json)")
}
