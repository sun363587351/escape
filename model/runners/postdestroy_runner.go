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

package runners

import (
	. "github.com/ankyra/escape-client/model/interfaces"
)

func NewPostDestroyRunner(stage string) Runner {
	return NewRunner(func(ctx RunnerContext) error {
		step := NewScriptStep(ctx, stage, "post_destroy", true)
		step.Commit = deleteCommit
		return step.Run(ctx)
	})
}

func deleteCommit(ctx RunnerContext, depl DeploymentState, stage string) error {
	if err := depl.SetVersion(stage, ""); err != nil {
		return err
	}
	if err := depl.UpdateInputs(stage, nil); err != nil {
		return err
	}
	if err := depl.UpdateOutputs(stage, nil); err != nil {
		return err
	}
	return nil
}
