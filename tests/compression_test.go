// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crd

import (
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Exercise list bounds and CEL runtime cost with a fully populated policy.
func TestCompressionListLimits(t *testing.T) {
	v := NewIstioValidator(t)
	for _, tt := range []struct {
		name               string
		ports, compressors int
		wantErr            string
	}{
		{"maximum-policy", 4096, 64, ""},
		{"too-many-ports", 4097, 1, "at most 4096 items"},
		{"too-many-compressors", 1, 65, "at most 64 items"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			compressors := make([]any, tt.compressors)
			for i := range compressors {
				compressors[i] = map[string]any{"gzip": map[string]any{}}
			}
			http := map[string]any{"compression": map[string]any{"compressors": compressors}}
			ports := make([]any, tt.ports)
			for i := range ports {
				ports[i] = map[string]any{"port": map[string]any{"number": int64(i + 1)}, "http": http}
			}
			obj := &unstructured.Unstructured{Object: map[string]any{
				"apiVersion": "networking.istio.io/v1alpha1", "kind": "InboundTrafficPolicy",
				"metadata": map[string]any{"name": tt.name},
				"spec":     map[string]any{"http": http, "portLevelSettings": ports},
			}}
			err := v.ValidateCustomResource(obj)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("wanted %q, got %v", tt.wantErr, err)
			}
		})
	}
}
