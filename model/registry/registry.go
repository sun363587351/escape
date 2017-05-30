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

package registry

import (
	. "github.com/ankyra/escape-client/model/interfaces"
	"github.com/ankyra/escape-client/model/registry/remote"
	core "github.com/ankyra/escape-core"
)

var GlobalRegistry Registry

func LoadRegistryFromConfig(cfg EscapeConfig) Registry {
	GlobalRegistry = remote.NewRemoteRegistry(
		cfg.GetCurrentTarget().GetApiServer(),
		cfg.GetCurrentTarget().GetAuthToken(),
	)
	return GlobalRegistry
}

func QueryReleaseMetadata(project, name, version string) (*core.ReleaseMetadata, error) {
	return GlobalRegistry.QueryReleaseMetadata(project, name, version)
}

func QueryNextVersion(project, name, versionPrefix string) (string, error) {
	return GlobalRegistry.QueryNextVersion(project, name, versionPrefix)
}

func DownloadRelease(project, name, version, targetFile string) error {
	return GlobalRegistry.DownloadRelease(project, name, version, targetFile)
}

func UploadRelease(project, releasePath string, metadata *core.ReleaseMetadata) error {
	return GlobalRegistry.UploadRelease(project, releasePath, metadata)
}
