// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"regexp"
	"strings"
	"sync/atomic"
)

// Inventory is the contents of Ansible inventory file.
type Inventory struct {
	Raw       []byte
	HostsRef  map[string]string `json:"host_refs,omitempty" yaml:"host_refs,omitempty"`
	GroupsRef map[string]bool   `json:"group_refs,omitempty" yaml:"group_refs,omitempty"`
	Hosts     []*InventoryHost  `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Groups    []*InventoryGroup `json:"groups,omitempty" yaml:"groups,omitempty"`
}

// InventoryHost is a host in Ansible inventory
type InventoryHost struct {
	Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
	Parent      string            `json:"parent_group,omitempty" yaml:"parent_group,omitempty"`
	Variables   map[string]string `json:"variables,omitempty" yaml:"variables,omitempty"`
	Groups      []string          `json:"groups,omitempty" yaml:"groups,omitempty"`
	GroupChains []string          `json:"group_chains,omitempty" yaml:"group_chains,omitempty"`
}

// InventoryGroup is an group of InventoryHost instances.
type InventoryGroup struct {
	Name      string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Ancestors []string               `json:"parent_groups,omitempty" yaml:"parent_groups,omitempty"`
	Variables map[string]string      `json:"variables,omitempty" yaml:"variables,omitempty"`
	Counters  InventoryGroupCounters `json:"counters,omitempty" yaml:"counters,omitempty"`
}

// InventoryGroupCounters are counters associated with InventoryGroup

