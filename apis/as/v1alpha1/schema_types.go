/*
Copyright 2021 NDD.

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

package v1alpha1

// NddrAsPool struct
type NddrRegistry struct {
	Registry []*NddrRegistryRegistry `json:"registry,omitempty"`
}

// NddrAsPoolAsPool struct
type NddrRegistryRegistry struct {
	AdminState         *string `json:"admin-state,omitempty"`
	AllocationStrategy *string `json:"allocation-strategy,omitempty"`
	End                *uint32 `json:"as-end,omitempty"`
	Start              *uint32 `json:"as-start,omitempty"`
	Description        *string `json:"description,omitempty"`
	//Name               *string                `json:"name,omitempty"`
	State *NddrRegistryRegistryState `json:"state,omitempty"`
}

// NddrRegistryRegistryState struct
type NddrRegistryRegistryState struct {
	Total     *uint32   `json:"total,omitempty"`
	Allocated *uint32   `json:"allocated,omitempty"`
	Available *uint32   `json:"available,omitempty"`
	Used      []*uint32 `json:"used,omitempty"`
}

// Root is the root of the schema
type Root struct {
	RegistryNddrRegistry *NddrRegistry `json:"nddr-ni-registry,omitempty"`
}
