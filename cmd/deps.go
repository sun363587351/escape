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

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Install dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("Unknown command '%s'", args[0])
		}
		cmd.UsageFunc()(cmd)
		return nil
	},
}

var depsFetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Install dependencies",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := context.LoadEscapePlan(escapePlanLocation)
		if err != nil {
			return err
		}
		return controllers.DepsController{}.Fetch(context)
	},
}

func init() {
	RootCmd.AddCommand(depsCmd)
	depsCmd.AddCommand(depsFetchCmd)
	setEscapePlanLocationFlag(depsCmd)
}