/*
Copyright paskal.maksim@gmail.com
Licensed under the Apache License, Version 2.0 (the "License")
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config_test

import (
	"strings"
	"testing"

	"github.com/maksim-paskal/hcloud-k8s-ctl/pkg/config"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	if err := config.Load(); err != nil {
		t.Fatal(err)
	}

	if config.Get().IPRange != "11.0.0.0/16" {
		t.Fatal("IPRange is not valid")
	}

	if config.Get().IPRangeSubnet != "11.0.0.0/17" {
		t.Fatal("IPRangeSubnet is not valid")
	}

	if config.Get().MasterCount != 33 {
		t.Fatal("MasterCount != 33")
	}

	if config.Get().HetznerToken.Main != "sometoken" {
		t.Fatal("HetznerToken.Main != sometoken")
	}

	if strings.Contains(config.String(), "sometoken") {
		t.Fatal("config has secret tokens")
	}
}
