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

package errand

import (
	. "github.com/ankyra/escape/model/runners"
	core "github.com/ankyra/escape-core"
)

func NewErrandRunner(errand *core.Errand, extraVars map[string]string) Runner {
	return NewRunner(func(ctx RunnerContext) error {
		step := NewScriptStep(ctx, "deploy", errand.Name, true)
		step.Inputs = func(ctx RunnerContext, stage string) (map[string]interface{}, error) {
			inputs, err := NewEnvironmentBuilder().GetInputsForErrand(ctx, errand, extraVars)
			if err != nil {
				return nil, err
			}
			return inputs, nil
		}
		step.ScriptPath = errand.Script
		return step.Run(ctx)
	})
}
