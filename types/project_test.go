/*
   Copyright 2020 The Compose Specification Authors.

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

package types

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_ApplyProfiles(t *testing.T) {
	p := makeProject()
	p.ApplyProfiles([]string{"foo"})
	assert.Equal(t, len(p.Services), 2)
	assert.Equal(t, p.Services[0].Name, "service_1")
	assert.Equal(t, p.Services[1].Name, "service_2")
	assert.Equal(t, len(p.DisabledServices), 1)
	assert.Equal(t, p.DisabledServices[0].Name, "service_3")
}

func Test_NoProfiles(t *testing.T) {
	p := makeProject()
	p.ApplyProfiles(nil)
	assert.Equal(t, len(p.Services), 1)
	assert.Equal(t, len(p.DisabledServices), 2)
	assert.Equal(t, p.Services[0].Name, "service_1")
}

func Test_ServiceProfiles(t *testing.T) {
	p := makeProject()
	services, err := p.GetServices([]string{"service_1", "service_2"})
	assert.NilError(t, err)

	profiles := services.GetProfiles()
	assert.Equal(t, len(profiles), 1)
	assert.Equal(t, profiles[0], "foo")
}

func Test_ForServices(t *testing.T) {
	p := makeProject()
	err := p.ForServices([]string{"service_2"})
	assert.NilError(t, err)

	assert.Equal(t, len(p.DisabledServices), 1)
	assert.Equal(t, p.DisabledServices[0].Name, "service_3")
}

func makeProject() Project {
	return Project{
		Services: append(Services{},
			ServiceConfig{
				Name: "service_1",
			}, ServiceConfig{
				Name:      "service_2",
				Profiles:  []string{"foo"},
				DependsOn: map[string]ServiceDependency{"service_1": {}},
			}, ServiceConfig{
				Name:     "service_3",
				Profiles: []string{"bar"},
			}),
	}
}
